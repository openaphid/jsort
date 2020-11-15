package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type Row struct {
	testName     string
	algo         string
	dataProvider string
	dataSize     int
	value        int
}

var pattern, _ = regexp.Compile(`(.*)/(.*)-(Random|Xor)-(\d+)`)

func parseRow(line string) (Row, error) {
	line = strings.TrimSpace(line)
	tokens := strings.Split(line, ",")

	if len(tokens) < 2 {
		return Row{}, fmt.Errorf("invalid length of splitted tokens: %v", tokens)
	}

	sub := pattern.FindStringSubmatch(tokens[0])

	if len(sub) < 4 {
		return Row{}, fmt.Errorf("can't parse test name: %v from %s", sub, tokens[0])
	}

	r := Row{}

	r.testName = sub[1]
	r.algo = sub[2]
	r.dataProvider = sub[3]

	dataSize, err := strconv.Atoi(sub[4])
	if err != nil {
		return Row{}, err
	}
	r.dataSize = dataSize

	value, err := strconv.ParseFloat(tokens[1], 64)
	if err != nil {
		return Row{}, err
	}
	r.value = int(value)

	return r, nil
}

type Group struct {
	name string
	rows []Row
}

func (g Group) filter(f func(r Row) bool) []Row {
	var ret []Row

	for _, r := range g.rows {
		if f(r) {
			ret = append(ret, r)
		}
	}

	return ret
}

func main() {
	benchFile := os.Args[1]

	dir := filepath.Dir(benchFile)

	content, err := ioutil.ReadFile(benchFile)

	if err != nil {
		log.Fatalf("read benchstats csv results: %v", err)
	}

	timeGroup := &Group{name: "time"}
	bytesGroup := &Group{name: "bytes"}
	allocsGroup := &Group{name: "allocs"}

	lines := strings.Split(string(content), "\n")

	var currentGroup *Group = nil

	testNames := make(map[string]int)

	for _, line := range lines {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		if strings.HasPrefix(line, "name,time/op") {
			currentGroup = timeGroup
			continue
		} else if strings.HasPrefix(line, "name,alloc/op (B/op)") {
			currentGroup = bytesGroup
			continue
		} else if strings.HasPrefix(line, "name,allocs/op (allocs/op)") {
			currentGroup = allocsGroup
			continue
		}

		row, err := parseRow(line)
		if err != nil {
			log.Println(err)
			continue
		}

		if currentGroup == nil {
			log.Panic("current group is nil")
		}

		testNames[row.testName]++

		currentGroup.rows = append(currentGroup.rows, row)
	}

	keyFunc := func(r Row) string {
		return fmt.Sprintf("%d|%s|%s", r.dataSize, r.testName, r.algo)
	}

	bytesMap := make(map[string]int)
	for _, r := range bytesGroup.rows {
		bytesMap[keyFunc(r)] = r.value
	}

	for testName, _ := range testNames {
		for _, dataProvider := range []string{"Random", "Xor"} {
			{
				rows := timeGroup.filter(func(r Row) bool {
					return r.testName == testName && r.dataProvider == dataProvider
				})

				outputFile := filepath.Join(dir, fmt.Sprintf("%s-%s.csv", testName, dataProvider))

				lines := make([]string, len(rows)+1)
				lines[0] = fmt.Sprintf("datasize,name,ns/op,B/op")
				for i, r := range rows {
					bytes, ok := bytesMap[keyFunc(r)]
					if !ok {
						log.Fatalf("can't find bytes data for row: %v", r)
					}
					lines[i+1] = fmt.Sprintf("%d,%s,%d,%d", r.dataSize, r.algo, r.value, bytes)
				}

				err := ioutil.WriteFile(outputFile, []byte(strings.Join(lines, "\n")), 0644)
				if err != nil {
					log.Fatal(outputFile, err)
				}

				log.Printf("saved %s\n", outputFile)

			}
		}
	}

}

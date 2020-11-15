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

type benchCtx struct {
	testName string
	algo     string
	dataSize int
}

type Row struct {
	benchCtx
	dataProvider string
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

	type DataProvider = string
	type LookupTable = map[DataProvider]map[benchCtx]int

	buildTableFunc := func(g *Group, testName string) LookupTable {
		table := make(LookupTable)
		for _, r := range g.rows {
			if r.testName != testName {
				continue
			}

			data, ok := table[r.dataProvider]
			if !ok {
				data = make(map[benchCtx]int)
				table[r.dataProvider] = data
			}
			key := r.benchCtx
			if _, ok := data[key]; ok {
				log.Fatal("duplicated data key: ", key)
			}
			data[key] = r.value
		}

		return table
	}

	for testName, _ := range testNames {
		providerTimeData := buildTableFunc(timeGroup, testName)
		providerBytesData := buildTableFunc(bytesGroup, testName)

		rows := timeGroup.filter(func(r Row) bool { // pick the random rows as a pivot
			return r.testName == testName && r.dataProvider == "Random"
		})

		providers := make([]string, len(providerTimeData))
		{
			providers[0] = "Random" // Random data set is the 1st
			i := 1
			for k, _ := range providerTimeData {
				if k == "Random" {
					continue
				}
				providers[i] = k
				i++
			}
		}

		lines := make([]string, len(rows)+1)
		lines[0] = fmt.Sprintf("datasize,name")
		for _, k := range providers {
			lines[0] = lines[0] + fmt.Sprintf(",ns/op (%s),B/op (%s)", k, k)
		}

		for i, r := range rows {
			var b strings.Builder
			fmt.Fprintf(&b, "%d,%s", r.dataSize, r.algo)

			for _, provider := range providers {
				bytes, ok := providerBytesData[provider][r.benchCtx]
				if !ok {
					log.Fatalf("can't find bytes data for: %v, %s", r.benchCtx, provider)
				}
				fmt.Fprintf(&b, ",%d,%d", providerTimeData[provider][r.benchCtx], bytes)
			}
			lines[i+1] = b.String()
		}

		outputFile := filepath.Join(dir, fmt.Sprintf("%s.csv", testName))

		err := ioutil.WriteFile(outputFile, []byte(strings.Join(lines, "\n")), 0644)
		if err != nil {
			log.Fatal(outputFile, err)
		}

		log.Printf("saved %s\n", outputFile)
	}

}

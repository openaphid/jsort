package jsort

import (
	"fmt"
	"github.com/openaphid/jsort/internal/sort_slice_dps_ts"
	"github.com/openaphid/jsort/internal/testdata"
	"io/ioutil"
	"sort"
	"strings"
	"testing"
)

type opStat struct {
	data  []Person
	stats map[string]int
}

func newOpStat(size int) *opStat {
	return &opStat{
		data:  make([]Person, size),
		stats: make(map[string]int),
	}
}

func (a opStat) resetStat() {
	a.stats = make(map[string]int)
}

type statInterface opStat

func (s statInterface) Len() int {
	s.stats["Len"] += 1
	return len(s.data)
}

func (s statInterface) Swap(i, j int) {
	s.stats["Swap"] += 1
	s.data[i], s.data[j] = s.data[j], s.data[i]
}

func (s statInterface) Less(i, j int) bool {
	s.stats["Less"] += 1
	return s.data[i].Age < s.data[j].Age
}

var (
	_ sort.Interface = (*statInterface)(nil)
)

type operationDataCase struct {
	name        string
	size        int
	prepareFunc func([]Person)
}

type compareRecordKey struct {
	algo string
	size int
}

func TestOperationStats(t *testing.T) {
	var dataCases []operationDataCase

	for _, s := range benchmarkSizes {
		dataCases = append(dataCases, operationDataCase{"Random", s, testdata.PrepareRandomAges})
	}

	for _, s := range benchmarkSizes {
		dataCases = append(dataCases, operationDataCase{"Xor", s, testdata.PrepareXorAges})
	}

	perNameData := make(map[string]map[compareRecordKey]int)
	perNameData["Random"] = make(map[compareRecordKey]int)
	perNameData["Xor"] = make(map[compareRecordKey]int)

	for _, c := range dataCases {
		name := c.name
		data := make([]Person, c.size)

		compareRecords, _ := perNameData[name]

		compareCounts := make(map[string]int)

		const loop = 20

		for i := 0; i < loop; i++ {
			c.prepareFunc(data)

			{
				dup := copyPersonSlice(data)

				cnt := 0
				sort_slice_dps_ts.Sort(dup, func(o1, o2 interface{}) int {
					cnt++
					return o1.(Person).Age - o2.(Person).Age
				})

				compareCounts["Unstable-DPS"] += cnt
			}

			{
				stat := newOpStat(c.size)
				copy(stat.data, data)

				Sort(statInterface(*stat))

				compareCounts["Stable-TimSort"] += stat.stats["Less"]
			}

			{
				stat := newOpStat(c.size)
				copy(stat.data, data)

				sort.Sort(statInterface(*stat))

				compareCounts["Unstable-Builtin"] += stat.stats["Less"]
			}

			{
				stat := newOpStat(c.size)
				copy(stat.data, data)

				sort.Stable(statInterface(*stat))

				compareCounts["Stable-Builtin"] += stat.stats["Less"]
			}
		}

		for algo, v := range compareCounts {
			k := compareRecordKey{algo: algo, size: c.size}
			compareRecords[k] = v / loop
		}
	}

	randomData := perNameData["Random"]
	xorData := perNameData["Xor"]

	var lines []string
	lines = append(lines, "size,name,compares/op (random),compares/op (xor)")

	var randomDataKeys []compareRecordKey
	for k, _ := range randomData {
		randomDataKeys = append(randomDataKeys, k)
	}
	Slice(randomDataKeys, func(i, j int) bool {
		return randomDataKeys[i].size < randomDataKeys[j].size
	})

	for _, k := range randomDataKeys {
		v := randomData[k]
		lines = append(lines, fmt.Sprintf("%d,%s,%d,%d", k.size, k.algo, v, xorData[k]))
	}

	outputFile := "BenchmarkResult/compares.csv"
	content := strings.Join(lines, "\n")
	fmt.Println("Saving to ", outputFile)
	fmt.Println(content)
	err := ioutil.WriteFile(outputFile, []byte(content), 0644)
	if err != nil {
		t.Fatal(outputFile, err)
	}
}

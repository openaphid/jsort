package jsort

import (
	"fmt"
	"github.com/openaphid/jsort/internal/sort_slice_dps_ts"
	"github.com/openaphid/jsort/internal/testdata"
	"sort"
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

func TestOperationStats(t *testing.T) {
	var dataCases []operationDataCase

	for _, s := range benchmarkSizes {
		dataCases = append(dataCases, operationDataCase{"Random", s, testdata.PrepareRandomAges})
	}

	for _, s := range benchmarkSizes {
		dataCases = append(dataCases, operationDataCase{"Xor", s, testdata.PrepareXorAges})
	}

	for _, c := range dataCases {
		name := c.name
		data := make([]Person, c.size)
		c.prepareFunc(data)

		{
			dup := copyPersonSlice(data)

			compareCount := 0

			sort_slice_dps_ts.Sort(dup, func(o1, o2 interface{}) int {
				compareCount++
				return o1.(Person).Age - o2.(Person).Age
			})

			fmt.Printf("DPS(%s-%d):%10d %s\n",
				name, c.size,
				compareCount, "Compare")
		}

		{
			stat := newOpStat(c.size)
			copy(stat.data, data)

			Sort(statInterface(*stat))

			fmt.Printf("TimSort(%s-%d):%10d %s\t%10d %s\t%10d %s\n",
				name, c.size,
				stat.stats["Less"], "Less",
				// The number of Swap is misleading here as it only counts the swap operation in the final sort after indices are fully sorted
				stat.stats["Swap"], "Swap(?)",
				stat.stats["Len"], "Len")
		}

		{
			stat := newOpStat(c.size)
			copy(stat.data, data)

			sort.Stable(statInterface(*stat))

			fmt.Printf("Builtin(%s-%d): %10d %s\t%10d %s\t%10d %s\n",
				name, c.size,
				stat.stats["Less"], "Less",
				stat.stats["Swap"], "Swap",
				stat.stats["Len"], "Len")
		}
	}
}

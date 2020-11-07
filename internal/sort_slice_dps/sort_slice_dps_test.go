package sort_slice_dps

import (
	"fmt"
	"github.com/openaphid/jsort/internal/testdata"
	"log"
	builtinsort "sort"
	"testing"
)

type Person = testdata.Person

var prepare = testdata.Prepare

var byAge = func(o1, o2 interface{}) int {
	return o1.(Person).Age - o2.(Person).Age
}

func TestByAge(t *testing.T) {
	for i := 1; i <= 1024*10; i++ {
		persons := make([]Person, i)
		prepare(persons)

		Sort(persons, byAge)

		sorted := IsSorted(persons, byAge)

		if !sorted {
			log.Panicf("should be sorted: %d", i)
		}
	}
}

var benchmarkSizes = testdata.GenBenchmarkSizes(256, 4, 5)

func BenchmarkVSSort(t *testing.B) {
	for _, size := range benchmarkSizes {
		var data = make([]Person, size)
		prepare(data)

		t.Run(fmt.Sprintf("DPS-%d", size), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
				t.StopTimer()
				dup := make([]Person, size)
				copy(dup, data)
				t.StartTimer()

				Sort(dup, func(o1, o2 interface{}) int {
					return o1.(Person).Age - o2.(Person).Age
				})
			}
		})

		t.Run(fmt.Sprintf("DPS-ManualConvert-%d", size), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
				t.StopTimer()
				dup := make([]Person, size)
				copy(dup, data)
				t.StartTimer()

				convert := make([]interface{}, size)
				for i, _ := range convert {
					convert[i] = dup[i]
				}

				SortInterfaces(convert, func(o1, o2 interface{}) int {
					return o1.(Person).Age - o2.(Person).Age
				})

				for i, _ := range dup {
					dup[i] = convert[i].(Person)
				}
			}
		})

		t.Run(fmt.Sprintf("Builtin-%d", size), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
				t.StopTimer()
				dup := make([]Person, size)
				copy(dup, data)
				t.StartTimer()

				builtinsort.Slice(dup, func(i, j int) bool {
					return dup[i].Age < dup[j].Age
				})
			}
		})
	}
}

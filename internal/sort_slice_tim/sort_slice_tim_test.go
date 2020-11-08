package sort_slice_tim

import (
	"fmt"
	"github.com/openaphid/jsort/internal/testdata"
	"log"
	builtinsort "sort"
	"testing"
)

type Person = testdata.Person

var prepare = testdata.PrepareRandomAges

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

func TestInts(t *testing.T) {
	a := []int{0, 2, 3, 1}
	Sort(a, func(i, j interface{}) int {
		return i.(int) - j.(int)
	})

	if !builtinsort.IntsAreSorted(a) {
		log.Panic("should be sorted: ", a)
	}
}

func TestShuffledSeq(t *testing.T) {
	for i := 1; i <= 1024*5; i++ {
		persons := make([]Person, i)
		testdata.PrepareShuffledSeq(persons)

		Sort(persons, byAge)

		for j := range persons {
			if persons[j].Age != j {
				t.Fatalf("Age(%d) should be %d for test #%d", persons[j].Age, j, i)
			}
		}
	}
}

var benchmarkSizes = []int{256, 1024, 4192, 16768, 67072}

func BenchmarkVSSort(t *testing.B) {
	for _, size := range benchmarkSizes {
		var data = make([]Person, size)
		prepare(data)

		t.Run(fmt.Sprintf("TimSort-%d", size), func(t *testing.B) {
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

		t.Run(fmt.Sprintf("TimSort-ManualConvert-%d", size), func(t *testing.B) {
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

		t.Run(fmt.Sprintf("BuiltinSort-%d", size), func(t *testing.B) {
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

		t.Run(fmt.Sprintf("BuiltinSortStable-%d", size), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
				t.StopTimer()
				dup := make([]Person, size)
				copy(dup, data)
				t.StartTimer()
				builtinsort.SliceStable(dup, func(i, j int) bool {
					return dup[i].Age < dup[j].Age
				})
			}
		})
	}
}

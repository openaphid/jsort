package sort_slice_tim_interface

import (
	"fmt"
	"github.com/openaphid/jsort/internal/sort_slice_tim_ts"
	"github.com/openaphid/jsort/internal/testdata"
	"log"
	builtinsort "sort"
	"testing"
)

type Person = testdata.Person

var prepare = testdata.PrepareRandomAges

type PersonInterface []Person

func (p PersonInterface) Len() int {
	return len(p)
}

func (p PersonInterface) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}

func (p PersonInterface) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

var _ builtinsort.Interface = (*PersonInterface)(nil)

func TestByAge(t *testing.T) {
	for i := 1; i <= 1024*10; i++ {
		persons := make(PersonInterface, i)
		prepare(persons)

		Sort(persons)

		sorted := IsSorted(persons)

		if !sorted {
			log.Panicf("should be sorted: %d", i)
		}
	}
}

func TestShuffledSeq(t *testing.T) {
	for i := 3; i <= 1024*5; i++ {
		persons := make(PersonInterface, i)
		testdata.PrepareShuffledSeq(persons)

		Sort(persons)

		for j := range persons {
			if persons[j].Age != j {
				t.Fatalf("Age(%d) should be %d for test #%d", persons[j].Age, j, i)
			}
		}
	}
}

var benchmarkSizes = testdata.GenBenchmarkSizes(256, 4, 5)

func BenchmarkVSSort(t *testing.B) {
	for _, size := range benchmarkSizes {
		var data = make([]Person, size)
		prepare(data)

		t.Run(fmt.Sprintf("TimSortInterface-%d", size), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
				t.StopTimer()
				dup := make([]Person, size)
				copy(dup, data)

				t.StartTimer()
				Sort(PersonInterface(dup))
			}
		})

		t.Run(fmt.Sprintf("TimSortInterfaceSlice-%d", size), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
				t.StopTimer()
				dup := make(PersonInterface, size)
				copy(dup, data)
				t.StartTimer()
				Slice(dup, func(i, j int) bool {
					return dup[i].Age < dup[j].Age
				})
			}
		})

		t.Run(fmt.Sprintf("TimSort-TypeAssertion-%d", size), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
				t.StopTimer()
				dup := make([]Person, size)
				copy(dup, data)
				t.StartTimer()
				sort_slice_tim_ts.Sort(dup, func(i, j interface{}) int {
					return i.(Person).Age - j.(Person).Age
				})
			}
		})

		t.Run(fmt.Sprintf("BuiltinSortInterface-%d", size), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
				t.StopTimer()
				dup := make([]Person, size)
				copy(dup, data)
				t.StartTimer()
				builtinsort.Sort(PersonInterface(dup))
			}
		})

		t.Run(fmt.Sprintf("BuiltinSortInterfaceStable-%d", size), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
				t.StopTimer()
				dup := make([]Person, size)
				copy(dup, data)
				t.StartTimer()
				builtinsort.Stable(PersonInterface(dup))
			}
		})

		t.Run(fmt.Sprintf("BuiltinSortSlice-%d", size), func(t *testing.B) {
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

		t.Run(fmt.Sprintf("BuiltinSortSliceStable-%d", size), func(t *testing.B) {
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

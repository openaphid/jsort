package dualpivotsort

import (
	"fmt"
	"github.com/openaphid/dualpivotsort/internal/testdata"
	"math/rand"
	"sort"
	"testing"
	"time"
)

type Person = testdata.Person

var prepare = testdata.Prepare

var benchmarkSizes = testdata.GenBenchmarkSizes(256, 4, 5)

func prepareInts(src []int) {
	rand.Seed(time.Now().Unix())
	for i := range src {
		src[i] = rand.Int()
	}
}

func copyInts(src []int) []int {
	dup := make([]int, len(src))
	copy(dup, src)

	return dup
}

type IntCompareInterface []int

func (a IntCompareInterface) Len() int {
	return len(a)
}

func (a IntCompareInterface) Compare(i, j int) int {
	return a[i] - a[j]
}

func (a IntCompareInterface) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

var _ CompareInterface = (*IntCompareInterface)(nil)

type PersonCompareInterface []Person

func (p PersonCompareInterface) Len() int {
	return len(p)
}

func (p PersonCompareInterface) Compare(i, j int) int {
	return p[i].Age - p[j].Age
}

func (p PersonCompareInterface) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

var _ CompareInterface = (*PersonCompareInterface)(nil)

func BenchmarkInts(t *testing.B) {
	for _, size := range benchmarkSizes {
		t.StopTimer()
		var data = make([]int, size)
		prepareInts(data)
		t.StartTimer()

		t.Run(fmt.Sprintf("DpsSpecializedInts-%d", size), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
				t.StopTimer()
				dup := copyInts(data)

				t.StartTimer()
				Ints(dup)
			}
		})

		t.Run(fmt.Sprintf("DpsSlice-%d", size), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
				t.StopTimer()
				dup := copyInts(data)

				t.StartTimer()
				Slice(dup, func(o1, o2 interface{}) int {
					return o1.(int) - o2.(int)
				})
			}
		})

		t.Run(fmt.Sprintf("TimSortSlice-%d", size), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
				t.StopTimer()
				dup := copyInts(data)

				t.StartTimer()
				SliceStable(dup, func(o1, o2 interface{}) int {
					return o1.(int) - o2.(int)
				})
			}
		})

		t.Run(fmt.Sprintf("TimSortSliceInterface-%d", size), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
				t.StopTimer()
				dup := copyInts(data)

				t.StartTimer()
				SliceInterface(IntCompareInterface(dup))
			}
		})

		t.Run(fmt.Sprintf("BuiltinSortInts-%d", size), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
				t.StopTimer()
				dup := copyInts(data)

				t.StartTimer()
				sort.Ints(dup)
			}
		})

		t.Run(fmt.Sprintf("BuiltinSortSpecializedInts-%d", size), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
				t.StopTimer()
				dup := copyInts(data)

				t.StartTimer()
				BuiltinSortSpecializedInts(dup)
			}
		})

		t.Run(fmt.Sprintf("BuiltinSortSlice-%d", size), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
				t.StopTimer()
				dup := copyInts(data)

				t.StartTimer()
				sort.Slice(dup, func(i, j int) bool {
					return dup[i] < dup[j]
				})
			}
		})

		t.Run(fmt.Sprintf("BuiltinSortSliceStable-%d", size), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
				t.StopTimer()
				dup := copyInts(data)

				t.StartTimer()
				sort.SliceStable(dup, func(i, j int) bool {
					return dup[i] < dup[j]
				})
			}
		})
	}
}

func copyPersonSlice(src []Person) []Person {
	dup := make([]Person, len(src))
	copy(dup, src)

	return dup
}

func BenchmarkStructSlice(t *testing.B) {
	for _, size := range benchmarkSizes {
		t.StopTimer()
		var data = make([]Person, size)
		prepare(data)
		t.StartTimer()

		t.Run(fmt.Sprintf("DpsSlice-%d", size), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
				t.StopTimer()
				dup := copyPersonSlice(data)

				t.StartTimer()
				Slice(dup, func(o1, o2 interface{}) int {
					return o1.(Person).Age - o2.(Person).Age
				})
			}
		})

		t.Run(fmt.Sprintf("TimSortSlice-%d", size), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
				t.StopTimer()
				dup := copyPersonSlice(data)

				t.StartTimer()
				SliceStable(dup, func(o1, o2 interface{}) int {
					return o1.(Person).Age - o2.(Person).Age
				})
			}
		})

		t.Run(fmt.Sprintf("TimSortSliceInterface-%d", size), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
				t.StopTimer()
				dup := copyPersonSlice(data)

				t.StartTimer()
				SliceInterface(PersonCompareInterface(dup))
			}
		})

		t.Run(fmt.Sprintf("BuiltinSortSlice-%d", size), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
				t.StopTimer()
				dup := copyPersonSlice(data)

				t.StartTimer()
				sort.Slice(dup, func(i, j int) bool {
					return dup[i].Age < dup[j].Age
				})
			}
		})

		t.Run(fmt.Sprintf("BuiltinSortSliceStable-%d", size), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
				t.StopTimer()
				dup := copyPersonSlice(data)

				t.StartTimer()
				sort.SliceStable(dup, func(i, j int) bool {
					return dup[i].Age < dup[j].Age
				})
			}
		})
	}
}

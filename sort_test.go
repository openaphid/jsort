package jsort

import (
	"fmt"
	"github.com/openaphid/jsort/internal/testdata"
	"math/rand"
	"sort"
	"strings"
	"testing"
	"time"
)

type Person = testdata.Person

var prepare = testdata.PrepareRandomAges

var benchmarkSizes = testdata.GenBenchmarkSizes(256, 4, 5)

func prepareRandomInts(src []int) {
	rand.Seed(time.Now().Unix())
	for i := range src {
		src[i] = rand.Int()
	}
}

func prepareXorInts(src []int) {
	for i := range src {
		src[i] = i ^ 0x2cc
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

type AgeCompareInterface []Person

func (p AgeCompareInterface) Len() int {
	return len(p)
}

func (p AgeCompareInterface) Compare(i, j int) int {
	return p[i].Age - p[j].Age
}

func (p AgeCompareInterface) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

var _ CompareInterface = (*AgeCompareInterface)(nil)

type NameCompareInterface []Person

func (p NameCompareInterface) Len() int {
	return len(p)
}

func (p NameCompareInterface) Compare(i, j int) int {
	return strings.Compare(p[i].Name, p[j].Name)
}

func (p NameCompareInterface) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

var _ CompareInterface = (*NameCompareInterface)(nil)

func BenchmarkInts(t *testing.B) {
	dataCases := []struct {
		name        string
		prepareFunc func([]int)
	}{
		{"Random", prepareRandomInts},
		{"Xor", prepareXorInts},
	}

	for _, size := range benchmarkSizes {
		var data = make([]int, size)

		for _, d := range dataCases {
			d.prepareFunc(data)
			name := d.name

			t.Run(fmt.Sprintf("DpsSpecializedInts-%s-%d", name, size), func(t *testing.B) {
				for i := 0; i < t.N; i++ {
					t.StopTimer()
					dup := copyInts(data)

					t.StartTimer()
					Ints(dup)
				}
			})

			t.Run(fmt.Sprintf("DpsSlice-%s-%d", name, size), func(t *testing.B) {
				for i := 0; i < t.N; i++ {
					t.StopTimer()
					dup := copyInts(data)

					t.StartTimer()
					Slice(dup, func(o1, o2 interface{}) int {
						return o1.(int) - o2.(int)
					})
				}
			})

			t.Run(fmt.Sprintf("TimSortSlice-%s-%d", name, size), func(t *testing.B) {
				for i := 0; i < t.N; i++ {
					t.StopTimer()
					dup := copyInts(data)

					t.StartTimer()
					SliceStable(dup, func(o1, o2 interface{}) int {
						return o1.(int) - o2.(int)
					})
				}
			})

			t.Run(fmt.Sprintf("TimSortSliceInterface-%s-%d", name, size), func(t *testing.B) {
				for i := 0; i < t.N; i++ {
					t.StopTimer()
					dup := copyInts(data)

					t.StartTimer()
					SliceInterface(IntCompareInterface(dup))
				}
			})

			t.Run(fmt.Sprintf("BuiltinSortInts-%s-%d", name, size), func(t *testing.B) {
				for i := 0; i < t.N; i++ {
					t.StopTimer()
					dup := copyInts(data)

					t.StartTimer()
					sort.Ints(dup)
				}
			})

			t.Run(fmt.Sprintf("BuiltinSortSpecializedInts-%s-%d", name, size), func(t *testing.B) {
				for i := 0; i < t.N; i++ {
					t.StopTimer()
					dup := copyInts(data)

					t.StartTimer()
					BuiltinSortSpecializedInts(dup)
				}
			})

			t.Run(fmt.Sprintf("BuiltinSortSlice-%s-%d", name, size), func(t *testing.B) {
				for i := 0; i < t.N; i++ {
					t.StopTimer()
					dup := copyInts(data)

					t.StartTimer()
					sort.Slice(dup, func(i, j int) bool {
						return dup[i] < dup[j]
					})
				}
			})

			t.Run(fmt.Sprintf("BuiltinSortSliceStable-%s-%d", name, size), func(t *testing.B) {
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
}

func copyPersonSlice(src []Person) []Person {
	dup := make([]Person, len(src))
	copy(dup, src)

	return dup
}

func BenchmarkStructSliceAge(t *testing.B) {
	dataCases := []struct {
		name        string
		prepareFunc func([]Person)
	}{
		{"Random", testdata.PrepareRandomAges},
		{"Xor", testdata.PrepareXorAges},
	}

	for _, size := range benchmarkSizes {
		var data = make([]Person, size)

		for _, c := range dataCases {
			c.prepareFunc(data)
			name := c.name

			t.Run(fmt.Sprintf("DpsSlice-%s-%d", name, size), func(t *testing.B) {
				for i := 0; i < t.N; i++ {
					t.StopTimer()
					dup := copyPersonSlice(data)

					t.StartTimer()
					Slice(dup, func(o1, o2 interface{}) int {
						return o1.(Person).Age - o2.(Person).Age
					})
				}
			})

			t.Run(fmt.Sprintf("TimSortSlice-%s-%d", name, size), func(t *testing.B) {
				for i := 0; i < t.N; i++ {
					t.StopTimer()
					dup := copyPersonSlice(data)

					t.StartTimer()
					SliceStable(dup, func(o1, o2 interface{}) int {
						return o1.(Person).Age - o2.(Person).Age
					})
				}
			})

			t.Run(fmt.Sprintf("TimSortSliceInterface-%s-%d", name, size), func(t *testing.B) {
				for i := 0; i < t.N; i++ {
					t.StopTimer()
					dup := copyPersonSlice(data)

					t.StartTimer()
					SliceInterface(AgeCompareInterface(dup))
				}
			})

			t.Run(fmt.Sprintf("BuiltinSortSlice-%s-%d", name, size), func(t *testing.B) {
				for i := 0; i < t.N; i++ {
					t.StopTimer()
					dup := copyPersonSlice(data)

					t.StartTimer()
					sort.Slice(dup, func(i, j int) bool {
						return dup[i].Age < dup[j].Age
					})
				}
			})

			t.Run(fmt.Sprintf("BuiltinSortSliceStable-%s-%d", name, size), func(t *testing.B) {
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
}

func BenchmarkStructSliceName(t *testing.B) {
	dataCases := []struct {
		name        string
		prepareFunc func([]Person)
	}{
		{"Random", testdata.PrepareRandomNames},
		{"Xor", testdata.PrepareXorNames},
	}

	for _, size := range benchmarkSizes {
		var data = make([]Person, size)

		for _, c := range dataCases {
			c.prepareFunc(data)
			name := c.name

			t.Run(fmt.Sprintf("DpsSlice-%s-%d", name, size), func(t *testing.B) {
				for i := 0; i < t.N; i++ {
					t.StopTimer()
					dup := copyPersonSlice(data)

					t.StartTimer()
					Slice(dup, func(o1, o2 interface{}) int {
						return strings.Compare(o1.(Person).Name, o2.(Person).Name)
					})
				}
			})

			t.Run(fmt.Sprintf("TimSortSlice-%s-%d", name, size), func(t *testing.B) {
				for i := 0; i < t.N; i++ {
					t.StopTimer()
					dup := copyPersonSlice(data)

					t.StartTimer()
					SliceStable(dup, func(o1, o2 interface{}) int {
						return strings.Compare(o1.(Person).Name, o2.(Person).Name)
					})
				}
			})

			t.Run(fmt.Sprintf("TimSortSliceInterface-%s-%d", name, size), func(t *testing.B) {
				for i := 0; i < t.N; i++ {
					t.StopTimer()
					dup := copyPersonSlice(data)

					t.StartTimer()
					SliceInterface(NameCompareInterface(dup))
				}
			})

			t.Run(fmt.Sprintf("BuiltinSortSlice-%s-%d", name, size), func(t *testing.B) {
				for i := 0; i < t.N; i++ {
					t.StopTimer()
					dup := copyPersonSlice(data)

					t.StartTimer()
					sort.Slice(dup, func(i, j int) bool {
						return dup[i].Name < dup[j].Name
					})
				}
			})

			t.Run(fmt.Sprintf("BuiltinSortSliceStable-%s-%d", name, size), func(t *testing.B) {
				for i := 0; i < t.N; i++ {
					t.StopTimer()
					dup := copyPersonSlice(data)

					t.StartTimer()
					sort.SliceStable(dup, func(i, j int) bool {
						return dup[i].Name < dup[j].Name
					})
				}
			})
		}
	}
}

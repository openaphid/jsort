package dualpivotsort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

var benchmarkSizes = []int{256, 1024, 4192, 16768}

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

type Person struct {
	age  int
	name string
}

func (p Person) String() string {
	return fmt.Sprintf("Person(%d, %s)", p.age, p.name)
}

func prepare(a []Person) {
	for i, _ := range a {
		a[i] = Person{
			age:  rand.Int(),
			name: fmt.Sprintf("n-%d", rand.Int()),
		}
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
					return o1.(Person).age - o2.(Person).age
				})
			}
		})

		t.Run(fmt.Sprintf("TimSortSlice-%d", size), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
				t.StopTimer()
				dup := copyPersonSlice(data)

				t.StartTimer()
				SliceStable(dup, func(o1, o2 interface{}) int {
					return o1.(Person).age - o2.(Person).age
				})
			}
		})

		t.Run(fmt.Sprintf("BuiltinSortSlice-%d", size), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
				t.StopTimer()
				dup := copyPersonSlice(data)

				t.StartTimer()
				sort.Slice(dup, func(i, j int) bool {
					return dup[i].age < dup[j].age
				})
			}
		})

		t.Run(fmt.Sprintf("BuiltinSortSliceStable-%d", size), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
				t.StopTimer()
				dup := copyPersonSlice(data)

				t.StartTimer()
				sort.SliceStable(dup, func(i, j int) bool {
					return dup[i].age < dup[j].age
				})
			}
		})
	}
}

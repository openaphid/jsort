package sort_slice_dps

import (
	"fmt"
	"log"
	"math/rand"
	builtinsort "sort"
	"testing"
)

type Person struct {
	age  int
	name string
}

func (p Person) String() string {
	return fmt.Sprintf("Person(%d, %s)", p.age, p.name)
}

type PersonSlice []Person

var invokeCount = make(map[string]int)

func resetCount() {
	invokeCount = make(map[string]int)
}

func (p PersonSlice) cmp(compare CompareFunc, i, j int) int {
	invokeCount["cmp"]++
	return compare(p[i], p[j])
}

var _ SliceInterface = (*PersonSlice)(nil)

func (p PersonSlice) make(n int) SliceInterface {
	invokeCount["make"]++
	return make(PersonSlice, n)
}

func (p PersonSlice) copy(src SliceInterface) {
	invokeCount["copy"]++
	copy(p, src.(PersonSlice))
}

func (p PersonSlice) get(i int) interface{} {
	invokeCount["get"]++
	return p[i]
}

func (p PersonSlice) set(i int, v interface{}) {
	invokeCount["set"]++
	p[i] = v.(Person)
}

func (p PersonSlice) swap(i, j int) {
	invokeCount["swap"]++
	p[i], p[j] = p[j], p[i]
}

func (p PersonSlice) len() int {
	invokeCount["len"]++
	return len(p)
}

func (p PersonSlice) slice(i int) SliceInterface {
	invokeCount["slice"]++
	return p[i:]
}

func (p PersonSlice) slice2(i, j int) SliceInterface {
	invokeCount["slice2"]++
	return p[i:j]
}

func prepare(a []Person) {
	for i, _ := range a {
		a[i] = Person{
			age:  rand.Int(),
			name: fmt.Sprintf("n-%d", rand.Int()),
		}
	}
}

func isSorted(a []Person) bool {
	n := len(a)
	for i := n - 1; i > 0; i-- {
		if a[i].age-a[i-1].age < 0 {
			return false
		}
	}
	return true
}

var byAge = func(o1, o2 interface{}) int {
	invokeCount["byAge"]++
	return o1.(Person).age - o2.(Person).age
}

func TestByAge(t *testing.T) {
	for i := 1; i <= 1024; i++ {
		persons := make(PersonSlice, i)
		prepare(persons)

		Sort(persons, byAge)

		sorted := IsSorted(persons, byAge)

		if !sorted {
			log.Panicf("should be sorted: %d", i)
		}
	}
}

const invokeDataSize = 1024

func TestDpsInvokeCount(t *testing.T) {
	persons := make([]Person, invokeDataSize)
	prepare(persons)

	{
		resetCount()
		persons2 := make(PersonSlice, invokeDataSize)
		for i, _ := range persons {
			persons2[i] = persons[i]
		}

		Sort(persons2, byAge)
		log.Printf("dps invoke count: %v", invokeCount)
	}

	{
		resetCount()
		persons2 := make(PersonBuiltinSlice, invokeDataSize)
		for i, _ := range persons {
			persons2[i] = persons[i]
		}

		builtinsort.Sort(persons2)
		log.Printf("builtin invoke count: %v", invokeCount)
	}
}

type PersonBuiltinSlice []Person

var _ builtinsort.Interface = (*PersonBuiltinSlice)(nil)

func (p PersonBuiltinSlice) Len() int {
	invokeCount["Len"]++
	return len(p)
}

func (p PersonBuiltinSlice) Less(i, j int) bool {
	invokeCount["Less"]++
	return p[i].age < p[j].age
}

func (p PersonBuiltinSlice) Swap(i, j int) {
	invokeCount["Swap"]++
	p[i], p[j] = p[j], p[i]
}

var benchmarkSizes = []int{256, 1024, 4192, 16768}

func BenchmarkDpsSortInterface(t *testing.B) {
	for _, size := range benchmarkSizes {
		t.Run(fmt.Sprintf("%d", size), func(t *testing.B) {
			var data = make(PersonSlice, size)
			prepare(data)

			t.ResetTimer()

			for i := 0; i < t.N; i++ {
				t.StopTimer()
				dup := make(PersonSlice, size)
				copy(dup, data)
				t.StartTimer()
				Sort(dup, byAge)
			}
		})
	}
}

func BenchmarkBuiltinSortSlice(t *testing.B) {
	for _, size := range benchmarkSizes {
		t.Run(fmt.Sprintf("%d", size), func(t *testing.B) {
			var data = make(PersonSlice, size)
			prepare(data)

			t.ResetTimer()

			for i := 0; i < t.N; i++ {
				t.StopTimer()
				dup := make(PersonSlice, size)
				copy(dup, data)
				t.StartTimer()
				builtinsort.Slice(dup, func(i, j int) bool {
					return dup[i].age < dup[j].age
				})
			}
		})
	}
}

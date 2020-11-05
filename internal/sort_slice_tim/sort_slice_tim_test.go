package sort_slice_tim

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

func (p PersonSlice) cmp(compare CompareFunc, i, j int) int {
	return compare(p[i], p[j])
}

var _ SliceInterface = (*PersonSlice)(nil)

func (p PersonSlice) make(n int) SliceInterface {
	return make(PersonSlice, n)
}

func (p PersonSlice) copy(src SliceInterface) {
	copy(p, src.(PersonSlice))
}

func (p PersonSlice) get(i int) interface{} {
	return p[i]
}

func (p PersonSlice) set(i int, v interface{}) {
	p[i] = v.(Person)
}

func (p PersonSlice) swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PersonSlice) len() int {
	return len(p)
}

func (p PersonSlice) slice(i int) SliceInterface {
	return p[i:]
}

func (p PersonSlice) slice2(i, j int) SliceInterface {
	return p[i:j]
}

type PersonSliceWithStat []Person

var invokeCount = make(map[string]int)

func resetCount() {
	invokeCount = make(map[string]int)
}

func (p PersonSliceWithStat) cmp(compare CompareFunc, i, j int) int {
	invokeCount["cmp"]++
	return compare(p[i], p[j])
}

var _ SliceInterface = (*PersonSliceWithStat)(nil)

func (p PersonSliceWithStat) make(n int) SliceInterface {
	invokeCount["make"]++
	return make(PersonSliceWithStat, n)
}

func (p PersonSliceWithStat) copy(src SliceInterface) {
	invokeCount["copy"]++
	copy(p, src.(PersonSliceWithStat))
}

func (p PersonSliceWithStat) get(i int) interface{} {
	invokeCount["get"]++
	return p[i]
}

func (p PersonSliceWithStat) set(i int, v interface{}) {
	invokeCount["set"]++
	p[i] = v.(Person)
}

func (p PersonSliceWithStat) swap(i, j int) {
	invokeCount["swap"]++
	p[i], p[j] = p[j], p[i]
}

func (p PersonSliceWithStat) len() int {
	invokeCount["len"]++
	return len(p)
}

func (p PersonSliceWithStat) slice(i int) SliceInterface {
	invokeCount["slice"]++
	return p[i:]
}

func (p PersonSliceWithStat) slice2(i, j int) SliceInterface {
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

var byAge = func(o1, o2 interface{}) int {
	return o1.(Person).age - o2.(Person).age
}

var byAgeWithStat = func(o1, o2 interface{}) int {
	invokeCount["byAge"]++
	return o1.(Person).age - o2.(Person).age
}

func TestByAge(t *testing.T) {
	for i := 1; i <= 1024*10; i++ {
		persons := make(PersonSlice, i)
		prepare(persons)

		Sort(persons, byAge)

		sorted := IsSorted(persons, byAge)

		if !sorted {
			log.Panicf("should be sorted: %d", i)
		}
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

const invokeDataSize = 1024

func TestInvokeCount(t *testing.T) {
	persons := make([]Person, invokeDataSize)
	prepare(persons)

	{
		resetCount()
		persons2 := make(PersonSliceWithStat, invokeDataSize)
		for i, _ := range persons {
			persons2[i] = persons[i]
		}

		Sort(persons2, byAgeWithStat)
		log.Printf("tim invoke count: %v", invokeCount)
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

var benchmarkSizes = []int{256, 1024, 4192, 16768}

func BenchmarkVSSort(t *testing.B) {
	for _, size := range benchmarkSizes {
		var data = make(PersonSlice, size)
		prepare(data)

		t.Run(fmt.Sprintf("TimSort-%d", size), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
				t.StopTimer()
				dup := make(PersonSlice, size)
				copy(dup, data)
				t.StartTimer()
				Sort(dup, func(o1, o2 interface{}) int {
					return o1.(Person).age - o2.(Person).age
				})
			}
		})

		t.Run(fmt.Sprintf("BuiltinSort-%d", size), func(t *testing.B) {
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

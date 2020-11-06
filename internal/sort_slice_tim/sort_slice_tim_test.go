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

var benchmarkSizes = []int{256, 1024, 4192, 16768}

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
					return o1.(Person).age - o2.(Person).age
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
					return o1.(Person).age - o2.(Person).age
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
					return dup[i].age < dup[j].age
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
					return dup[i].age < dup[j].age
				})
			}
		})
	}
}

package sort_slice_reflect

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
	return o1.(Person).age - o2.(Person).age
}

func TestByAge(t *testing.T) {
	for i := 1; i <= 1024; i++ {
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

func BenchmarkDpsSortReflect(t *testing.B) {
	for _, size := range benchmarkSizes {
		t.Run(fmt.Sprintf("%d", size), func(t *testing.B) {
			var data = make([]Person, size)
			prepare(data)

			t.ResetTimer()

			for i := 0; i < t.N; i++ {
				t.StopTimer()
				dup := make([]Person, size)
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
			var data = make([]Person, size)
			prepare(data)

			t.ResetTimer()

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
	}
}

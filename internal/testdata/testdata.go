package testdata

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Person struct {
	Age  int
	Name string
}

func (p Person) String() string {
	return fmt.Sprintf("Person(%d, %s)", p.Age, p.Name)
}

func PrepareRandomAges(a []Person) {
	rand.Seed(time.Now().Unix())
	for i, _ := range a {
		a[i] = Person{
			Age:  rand.Int(),
			Name: fmt.Sprintf("n-%d", i),
		}
	}
}

func PrepareRandomNames(a []Person) {
	rand.Seed(time.Now().Unix())
	for i, _ := range a {
		a[i] = Person{
			Age:  i,
			Name: fmt.Sprintf("n-%d", rand.Int()),
		}
	}
}

func PrepareXorAges(a []Person) {
	for i, _ := range a {
		a[i] = Person{
			Age:  i ^ 0x2cc,
			Name: fmt.Sprintf("n-%d", i),
		}
	}
}

func PrepareXorNames(a []Person) {
	for i, _ := range a {
		a[i] = Person{
			Age:  i,
			Name: fmt.Sprintf("n-%d", i^0x2cc),
		}
	}
}

func PrepareShuffledSeq(a []Person) {
	rand.Seed(time.Now().Unix())
	for i, _ := range a {
		a[i] = Person{
			Age:  i,
			Name: fmt.Sprintf("n-%d", i),
		}
	}

	for i, _ := range a {
		j := rand.Intn(len(a))
		a[i], a[j] = a[j], a[i]
	}
}

func DumpData(a []Person) {
	ages := make([]string, len(a))
	for i, _ := range ages {
		ages[i] = strconv.Itoa(a[i].Age)
	}
	ioutil.WriteFile("/tmp/data.txt", []byte(strings.Join(ages, ",")), 0644)
}

func LoadData() []Person {
	bytes, _ := ioutil.ReadFile("/tmp/data.txt")
	data := string(bytes)

	var ret []Person
	for _, token := range strings.Split(data, ",") {
		age, _ := strconv.Atoi(token)
		ret = append(ret, Person{
			Age:  age,
			Name: fmt.Sprintf("n-%d", len(ret)),
		})
	}

	return ret
}

func GenBenchmarkSizes(init, multiplier, length int) []int {
	var sizes = []int{}

	s := init

	for i := 0; i < length; i++ {
		sizes = append(sizes, s)
		s *= multiplier
	}

	return sizes
}

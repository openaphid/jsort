// Code generated from sort_primitive.go using sort_primitive_gen.go; DO NOT EDIT.

package sort_uint

import (
	"math/rand"
	builtinsort "sort"
	"testing"
	"time"
)

func prepareRandom(a []primitive) {
	rand.Seed(time.Now().Unix())
	for i := range a {
		a[i] = primitive(rand.Int63())
	}
}

func prepareShuffledSeq(a []primitive) {
	for i := range a {
		a[i] = primitive(i)
	}

	for i, _ := range a {
		j := rand.Intn(len(a))
		a[i], a[j] = a[j], a[i]
	}
}

func duplicate(a []primitive) []primitive {
	r := make([]primitive, len(a))

	for i, v := range a {
		r[i] = v
	}

	return r
}

func TestDpsUints(t *testing.T) {
	for i := 1; i <= 1024*10; i++ {
		data := make([]primitive, i)
		prepareRandom(data)
		Sort(data)

		if !IsSorted(data) {
			t.Fatal("Test #d, should be sorted", i)
		}
	}
}

func TestDpsSeq(t *testing.T) {
	for i := 1; i <= 1024*10; i++ {
		data := make([]primitive, i)
		prepareShuffledSeq(data)

		ref := duplicate(data)
		builtinsort.Slice(ref, func(i, j int) bool {
			return ref[i] < ref[j]
		})

		Sort(data)

		for j := 0; j < len(data); j++ {
			if data[j] != ref[j] {
				t.Fatalf("Test %d, #%d %v should be %v", i, j, data[j], ref[j])
			}
		}
	}
}

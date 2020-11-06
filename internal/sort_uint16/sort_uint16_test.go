// Code generated from sort_primitive.go using sort_primitive_gen.go; DO NOT EDIT.
//go:generate go run genprimitives.go

package sort_uint16

import (
	"math/rand"
	"testing"
	"time"
)

func prepare(src []primitive) {
	rand.Seed(time.Now().Unix())
	for i := range src {
		src[i] = primitive(rand.Int63())
	}
}

func TestDpsUint16s(t *testing.T) {
	for i := 1; i <= 1024*10; i++ {
		data := make([]primitive, i)
		prepare(data)
		Sort(data)

		if !IsSorted(data) {
			t.Fatal("should be sorted")
		}
	}
}

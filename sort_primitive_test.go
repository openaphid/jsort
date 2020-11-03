// +build ignore
//go:generate go run genprimitives.go

package dualpivotsort

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

func TestDpsPrimitive(t *testing.T) {
	for i := 1; i <= 1024*10; i++ {
		data := make([]primitive, i)
		prepare(data)
		Sort(data)

		if !IsSorted(data) {
			t.Fatal("should be sorted")
		}
	}
}

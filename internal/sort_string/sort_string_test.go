package sort_string

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func prepare(src []primitive) {
	rand.Seed(time.Now().Unix())
	for i := range src {
		src[i] = strconv.Itoa(rand.Int())
	}
}

func TestDpsStrings(t *testing.T) {
	for i := 1; i <= 1024*10; i++ {
		data := make([]primitive, i)
		prepare(data)
		Sort(data)

		if !IsSorted(data) {
			t.Fatal("should be sorted")
		}
	}
}

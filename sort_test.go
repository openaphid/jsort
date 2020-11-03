package dualpivotsort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func prepareInts(src []int) {
	rand.Seed(time.Now().Unix())
	for i := range src {
		src[i] = rand.Int()
	}
}

var benchmarkSizes = []int{256, 1024, 4192, 16768}

func BenchmarkDpsInts(t *testing.B) {
	for _, size := range benchmarkSizes {
		t.Run(fmt.Sprintf("%d", size), func(t *testing.B) {
			var data = make([]int, size)
			prepareInts(data)

			t.ResetTimer()

			for i := 0; i < t.N; i++ {
				t.StopTimer()
				dup := make([]int, size)
				copy(dup, data)
				t.StartTimer()
				Ints(dup)
			}
		})
	}
}

func BenchmarkBuiltinInts(t *testing.B) {
	for _, size := range benchmarkSizes {
		t.Run(fmt.Sprintf("%d", size), func(t *testing.B) {
			var data = make([]int, size)
			prepareInts(data)

			t.ResetTimer()

			for i := 0; i < t.N; i++ {
				t.StopTimer()
				dup := make([]int, size)
				copy(dup, data)
				t.StartTimer()
				sort.Ints(dup)
			}
		})
	}
}

// Code generated by go2go; DO NOT EDIT.


//line sort_slice_tim_go2.go2:1
package sort_slice_go2

//line sort_slice_tim_go2.go2:1
import (
//line sort_slice_tim_go2.go2:1
 "fmt"
//line sort_slice_tim_go2.go2:1
 "math/rand"
//line sort_slice_tim_go2.go2:1
 "sort"
//line sort_slice_tim_go2.go2:1
 "strings"
//line sort_slice_tim_go2.go2:1
 "testing"
//line sort_slice_tim_go2.go2:1
 "time"
//line sort_slice_tim_go2.go2:1
)

//line sort_slice_tim_go2.go2:33
const MIN_MERGE = 32

//line sort_slice_tim_go2.go2:39
const MIN_GALLOP = 7

//line sort_slice_tim_go2.go2:48
const INITIAL_TMP_STORAGE_LENGTH = 256

//line sort_slice_tim_go2.go2:388
func minRunLength(n int) int {

	var r = 0
	for n >= MIN_MERGE {
		r |= (n & 1)
		n >>= 1
	}
	return n + r
}

//line sort_slice_tim_go2.go2:1044
func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}

//line sort_slice_tim_go2.go2:1050
var _ = fmt.Errorf
//line sort_slice_tim_go2.go2:1050
var _ = rand.ExpFloat64

//line sort_slice_tim_go2.go2:1050
type _ sort.Float64Slice
//line sort_slice_tim_go2.go2:1050
type _ strings.Builder

//line sort_slice_tim_go2.go2:1050
var _ = testing.AllocsPerRun

//line sort_slice_tim_go2.go2:1050
const _ = time.ANSIC
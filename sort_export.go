// Code generated from sort_primitive.go using sort_primitive_gen.go; DO NOT EDIT.

package jsort

import (
	"github.com/openaphid/jsort/internal/sort_float32"
	"github.com/openaphid/jsort/internal/sort_float64"
	"github.com/openaphid/jsort/internal/sort_int"
	"github.com/openaphid/jsort/internal/sort_int16"
	"github.com/openaphid/jsort/internal/sort_int32"
	"github.com/openaphid/jsort/internal/sort_int64"
	"github.com/openaphid/jsort/internal/sort_int8"
	"github.com/openaphid/jsort/internal/sort_slice_tim_interface"
	"github.com/openaphid/jsort/internal/sort_string"
	"github.com/openaphid/jsort/internal/sort_uint"
	"github.com/openaphid/jsort/internal/sort_uint16"
	"github.com/openaphid/jsort/internal/sort_uint32"
	"github.com/openaphid/jsort/internal/sort_uint64"
	"github.com/openaphid/jsort/internal/sort_uint8"
	"sort"
)

// int
var Ints func([]int) = sort_int.Sort
var IntsAreSorted func([]int) bool = sort_int.IsSorted

// int8
var Int8s func([]int8) = sort_int8.Sort
var Int8sAreSorted func([]int8) bool = sort_int8.IsSorted

// int16
var Int16s func([]int16) = sort_int16.Sort
var Int16sAreSorted func([]int16) bool = sort_int16.IsSorted

// int32
var Int32s func([]int32) = sort_int32.Sort
var Int32sAreSorted func([]int32) bool = sort_int32.IsSorted

// int64
var Int64s func([]int64) = sort_int64.Sort
var Int64sAreSorted func([]int64) bool = sort_int64.IsSorted

// uint
var Uints func([]uint) = sort_uint.Sort
var UintsAreSorted func([]uint) bool = sort_uint.IsSorted

// uint8
var Uint8s func([]uint8) = sort_uint8.Sort
var Uint8sAreSorted func([]uint8) bool = sort_uint8.IsSorted

// uint16
var Uint16s func([]uint16) = sort_uint16.Sort
var Uint16sAreSorted func([]uint16) bool = sort_uint16.IsSorted

// uint32
var Uint32s func([]uint32) = sort_uint32.Sort
var Uint32sAreSorted func([]uint32) bool = sort_uint32.IsSorted

// uint64
var Uint64s func([]uint64) = sort_uint64.Sort
var Uint64sAreSorted func([]uint64) bool = sort_uint64.IsSorted

// float32
var Float32s func([]float32) = sort_float32.Sort
var Float32sAreSorted func([]float32) bool = sort_float32.IsSorted

// float64
var Float64s func([]float64) = sort_float64.Sort
var Float64sAreSorted func([]float64) bool = sort_float64.IsSorted

// string
var Strings func([]string) = sort_string.Sort
var StringsAreSorted func([]string) bool = sort_string.IsSorted

// byte
var Bytes func([]byte) = sort_uint8.Sort
var BytesAreSorted func([]byte) bool = sort_uint8.IsSorted

// rune
var Runes func([]rune) = sort_int32.Sort
var RunesAreSorted func([]rune) bool = sort_int32.IsSorted

// The following APIs are compatible with the ones in the built-in `sort` package
// One difference is that all sort functions are stable by using timsort
var Sort func(data sort.Interface) = sort_slice_tim_interface.Sort
var Stable func(data sort.Interface) = sort_slice_tim_interface.Sort
var IsSorted func(data sort.Interface) bool = sort_slice_tim_interface.IsSorted
var Slice func(slice interface{}, less func(i, j int) bool) = sort_slice_tim_interface.Slice
var SliceStable func(slice interface{}, less func(i, j int) bool) = sort_slice_tim_interface.Slice
var SliceIsSorted func(slice interface{}, less func(i, j int) bool) bool = sort_slice_tim_interface.SliceIsSorted

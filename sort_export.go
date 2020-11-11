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
)

// int
var Ints = sort_int.Sort
var IntsAreSorted = sort_int.IsSorted

// int8
var Int8s = sort_int8.Sort
var Int8sAreSorted = sort_int8.IsSorted

// int16
var Int16s = sort_int16.Sort
var Int16sAreSorted = sort_int16.IsSorted

// int32
var Int32s = sort_int32.Sort
var Int32sAreSorted = sort_int32.IsSorted

// int64
var Int64s = sort_int64.Sort
var Int64sAreSorted = sort_int64.IsSorted

// uint
var Uints = sort_uint.Sort
var UintsAreSorted = sort_uint.IsSorted

// uint8
var Uint8s = sort_uint8.Sort
var Uint8sAreSorted = sort_uint8.IsSorted

// uint16
var Uint16s = sort_uint16.Sort
var Uint16sAreSorted = sort_uint16.IsSorted

// uint32
var Uint32s = sort_uint32.Sort
var Uint32sAreSorted = sort_uint32.IsSorted

// uint64
var Uint64s = sort_uint64.Sort
var Uint64sAreSorted = sort_uint64.IsSorted

// float32
var Float32s = sort_float32.Sort
var Float32sAreSorted = sort_float32.IsSorted

// float64
var Float64s = sort_float64.Sort
var Float64sAreSorted = sort_float64.IsSorted

// string
var Strings = sort_string.Sort
var StringsAreSorted = sort_string.IsSorted

// byte
var Bytes = sort_uint8.Sort
var BytesAreSorted = sort_uint8.IsSorted

// rune
var Runes = sort_int32.Sort
var RunesAreSorted = sort_int32.IsSorted

// The following APIs are compatible with the ones in the built-in `sort` package
var Sort = sort_slice_tim_interface.Sort
var Stable = sort_slice_tim_interface.Sort
var Slice = sort_slice_tim_interface.Slice
var SliceStable = sort_slice_tim_interface.Slice
var SliceIsSorted = sort_slice_tim_interface.SliceIsSorted

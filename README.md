# Dual-pivot sort and timsort in Go

Please refer to the blog post for more details: [Porting Dual-Pivot Sort and Timsort from Java to Go
](https://hubo.dev/2020-11-15-porting-dual-pivot-sort-and-timsort-from-java-to-go/)

# Installation

```
go get -u github.com/openaphid/jsort
```

# Quick Start

The [public APIs](https://github.com/openaphid/jsort/blob/main/sort_export.go#L85) for generic slice types are compatible with Go's built-in `sort` package. The sorting algorithm is timsort.

```go
type Person struct {
	Age  int
	Name string
}

type ByAgeInterface []Person

func (p ByAgeInterface) Len() int {
	return len(p)
}

func (p ByAgeInterface) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}

func (p ByAgeInterface) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

var _ sort.Interface = (*ByAgeInterface)(nil)

func main() {
    data := make([]Person, 256)

    jsort.Sort(ByAgeInterface(data)) // sort.Sort(ByAgeInterface(data))
    jsort.Stable(ByAgeInterface(data)) // sort.Stable(ByAgeInterface(data))
    jsort.IsSorted(ByAgeInterface(data)) // sort.IsSorted(ByAgeInterface(data))

    less := func(i, j int) bool {
    	return data[i].Age < data[j].Age
    }

    jsort.Slice(data, less) // sort.Slice(data, less)
    jsort.SliceStable(data, less) // sort.SliceStable(data, less)

    _ = jsort.SliceIsSorted(data, less) // sort.SliceIsSorted(data, less)
}
```

There are additional [specialized APIs](https://github.com/openaphid/jsort/blob/main/sort_export.go#L23) for primitive slice types. The algorithm is dual-pivot sort.

```go
data := []int64{3, 2, 1, 4}
jsort.Int64s(data)
```

# Performance

There is a benchmark section in the [blog post](https://hubo.dev/2020-11-15-porting-dual-pivot-sort-and-timsort-from-java-to-go/). Basically speaking, the timsort APIs can be 2-5x times faster than Go's `sort.Stable` in the cost of extra space. The dual-pivot APIs can be 3-5x times faster than `sort.Sort` without extra space.

![Benchmark: Sort By Name](https://hubo.dev/assets/datawrapper/datawrapper-sort-by-name.png)
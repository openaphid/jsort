// Code generated by go2go; DO NOT EDIT.


//line sort_slice_tim_test.go2:1
package sort_slice_tim_go2

//line sort_slice_tim_test.go2:1
import (
//line sort_slice_tim_test.go2:7
 builtinsort "sort"
//line sort_slice_tim_test.go2:7
 "fmt"
//line sort_slice_tim_test.go2:7
 "log"
//line sort_slice_tim_test.go2:7
 "math/rand"
//line sort_slice_tim_test.go2:7
 "testing"
//line sort_slice_tim_test.go2:7
)

//line sort_slice_tim_test.go2:11
type Person struct {
	age  int
	name string
}

func (p Person) String() string {
	return fmt.Sprintf("Person(%d, %s)", p.age, p.name)
}

func prepare(a []Person) {
	for i, _ := range a {
		a[i] = Person{
			age:  rand.Int(),
			name: fmt.Sprintf("n-%d", rand.Int()),
		}
	}
}

var byAge = func(o1, o2 Person) int {
	return o1.age - o2.age
}

func TestByAge(t *testing.T) {
	for i := 1; i <= 1024*10; i++ {
						persons := make([]Person, i)
						prepare(persons)
//line sort_slice_tim_test.go2:36
  instantiate୦୦Sort୦sort_slice_tim_go2୮aPerson(persons, byAge)

//line sort_slice_tim_test.go2:40
  sorted := instantiate୦୦IsSorted୦sort_slice_tim_go2୮aPerson(persons, byAge)

		if !sorted {
			log.Panicf("should be sorted: %d", i)
		}
	}
}

var benchmarkSizes = []int{256, 1024, 4192, 16768}

func BenchmarGo2kVSSort(t *testing.B) {
	for _, size := range benchmarkSizes {
		var data = make([]Person, size)
		prepare(data)

		t.Run(fmt.Sprintf("TimSort-%d", size), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
								t.StopTimer()
								dup := make([]Person, size)
								copy(dup, data)
								t.StartTimer()
//line sort_slice_tim_test.go2:60
    instantiate୦୦Sort୦sort_slice_tim_go2୮aPerson(dup, byAge)
//line sort_slice_tim_test.go2:62
   }
		})

		t.Run(fmt.Sprintf("BuiltinSort-%d", size), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
				t.StopTimer()
				dup := make([]Person, size)
				copy(dup, data)
				t.StartTimer()
				builtinsort.Slice(dup, func(i, j int) bool {
					return dup[i].age < dup[j].age
				})
			}
		})
	}
}

func prepareInts(a []int) {
	for i, _ := range a {
		a[i] = rand.Int()
	}
}

func BenchmarkGo2VSSortInts(t *testing.B) {
	for _, size := range benchmarkSizes {
		var data = make([]int, size)
		prepareInts(data)

		t.Run(fmt.Sprintf("TimSort-%d", size), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
								t.StopTimer()
								dup := make([]int, size)
								copy(dup, data)
								t.StartTimer()
//line sort_slice_tim_test.go2:95
    instantiate୦୦Sort୦int(dup, func(i, j int) int { return i - j })
//line sort_slice_tim_test.go2:97
   }
		})

		t.Run(fmt.Sprintf("BuiltinSort-%d", size), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
				t.StopTimer()
				dup := make([]int, size)
				copy(dup, data)
				t.StartTimer()
				builtinsort.Ints(dup)
			}
		})
	}
}
//line sort_slice_tim.go2:3
func instantiate୦୦Sort୦sort_slice_tim_go2୮aPerson(a []Person, compare func(i, j Person,) int) {
//line sort_slice_tim.go2:3
 instantiate୦୦sort୦sort_slice_tim_go2୮aPerson(a, 0, len(a), compare, nil, 0, 0)
//line sort_slice_tim.go2:5
}

func instantiate୦୦IsSorted୦sort_slice_tim_go2୮aPerson(a []Person, compare func(i, j Person,) int) bool {
	for i := len(a) - 1; i > 0; i-- {
		if compare(a[i], a[i-1]) < 0 {
			return false
		}
	}
	return true
}
//line sort_slice_tim.go2:3
func instantiate୦୦Sort୦int(a []int, compare func(i, j int,) int) {
//line sort_slice_tim.go2:3
 instantiate୦୦sort୦int(a, 0, len(a), compare, nil, 0, 0)
//line sort_slice_tim.go2:5
}

//line sort_slice_tim.go2:177
func instantiate୦୦sort୦sort_slice_tim_go2୮aPerson(a []Person, lo int, hi int, c func(i, j Person,) int, work []Person, workBase int, workLen int) {

//line sort_slice_tim.go2:180
 var nRemaining = hi - lo
	if nRemaining < 2 {
		return
	}

//line sort_slice_tim.go2:186
 if nRemaining < MIN_MERGE {
					var initRunLen = instantiate୦୦countRunAndMakeAscending୦sort_slice_tim_go2୮aPerson(a, lo, hi, c)
//line sort_slice_tim.go2:187
  instantiate୦୦binarySort୦sort_slice_tim_go2୮aPerson(a, lo, hi, lo+initRunLen, c)
//line sort_slice_tim.go2:189
  return
	}

//line sort_slice_tim.go2:197
 var ts = instantiate୦୦newTimSort୦sort_slice_tim_go2୮aPerson(a, c, work, workBase, workLen)
	var minRun = minRunLength(nRemaining)
	for {

					var runLen = instantiate୦୦countRunAndMakeAscending୦sort_slice_tim_go2୮aPerson(a, lo, hi, c)

//line sort_slice_tim.go2:204
  if runLen < minRun {
			force := 0
			if nRemaining <= minRun {
				force = nRemaining
			} else {
				force = minRun
			}
//line sort_slice_tim.go2:210
   instantiate୦୦binarySort୦sort_slice_tim_go2୮aPerson(a, lo, lo+force, lo+runLen, c)
//line sort_slice_tim.go2:212
   runLen = force
		}

//line sort_slice_tim.go2:216
  ts.pushRun(lo, runLen)
					ts.mergeCollapse()

//line sort_slice_tim.go2:220
  lo += runLen
		nRemaining -= runLen
		if nRemaining != 0 {
		} else {
			break
		}
	}

//line sort_slice_tim.go2:230
 ts.mergeForceCollapse()

}
//line sort_slice_tim.go2:177
func instantiate୦୦sort୦int(a []int, lo int, hi int, c func(i, j int,) int, work []int, workBase int, workLen int) {

//line sort_slice_tim.go2:180
 var nRemaining = hi - lo
	if nRemaining < 2 {
		return
	}

//line sort_slice_tim.go2:186
 if nRemaining < MIN_MERGE {
					var initRunLen = instantiate୦୦countRunAndMakeAscending୦int(a, lo, hi, c)
//line sort_slice_tim.go2:187
  instantiate୦୦binarySort୦int(a, lo, hi, lo+initRunLen, c)
//line sort_slice_tim.go2:189
  return
	}

//line sort_slice_tim.go2:197
 var ts = instantiate୦୦newTimSort୦int(a, c, work, workBase, workLen)
	var minRun = minRunLength(nRemaining)
	for {

					var runLen = instantiate୦୦countRunAndMakeAscending୦int(a, lo, hi, c)

//line sort_slice_tim.go2:204
  if runLen < minRun {
			force := 0
			if nRemaining <= minRun {
				force = nRemaining
			} else {
				force = minRun
			}
//line sort_slice_tim.go2:210
   instantiate୦୦binarySort୦int(a, lo, lo+force, lo+runLen, c)
//line sort_slice_tim.go2:212
   runLen = force
		}

//line sort_slice_tim.go2:216
  ts.pushRun(lo, runLen)
					ts.mergeCollapse()

//line sort_slice_tim.go2:220
  lo += runLen
		nRemaining -= runLen
		if nRemaining != 0 {
		} else {
			break
		}
	}

//line sort_slice_tim.go2:230
 ts.mergeForceCollapse()

}

//line sort_slice_tim.go2:327
func instantiate୦୦countRunAndMakeAscending୦sort_slice_tim_go2୮aPerson(a []Person, lo int, hi int, c func(i, j Person,) int) int {

	var runHi = lo + 1
	if runHi == hi {
		return 1
	}

//line sort_slice_tim.go2:335
 runHi++
	if c(a[runHi-1], a[lo]) < 0 {
		for runHi < hi && c(a[runHi], a[runHi-1]) < 0 {
			runHi++
		}
//line sort_slice_tim.go2:339
  instantiate୦୦reverseRange୦sort_slice_tim_go2୮aPerson(a, lo, runHi)
//line sort_slice_tim.go2:341
 } else {
		for runHi < hi && c(a[runHi], a[runHi-1]) >= 0 {
			runHi++
		}
	}

	return runHi - lo
}
//line sort_slice_tim.go2:252
func instantiate୦୦binarySort୦sort_slice_tim_go2୮aPerson(a []Person, lo int, hi int, start int, c func(i, j Person,) int) {

	if start == lo {
		start++
	}

	for ; start < hi; start++ {
					var pivot = a[start]

//line sort_slice_tim.go2:262
  var left = lo
					var right = start

//line sort_slice_tim.go2:270
  for left < right {
			var mid = (left + right) >> 1
			if c(pivot, a[mid]) < 0 {
				right = mid
			} else {
				left = mid + 1
			}
		}

//line sort_slice_tim.go2:287
  var n = start - left

		switch n {
		case 2:
			a[left+2] = a[left+1]
		case 1:
			a[left+1] = a[left]
		default:
			copy(a[left+1:], a[left:left+n])

		}
		a[left] = pivot
	}
}
//line sort_slice_tim.go2:101
func instantiate୦୦newTimSort୦sort_slice_tim_go2୮aPerson(a []Person, c func(i, j Person,) int, work []Person, workBase int, workLen int) *instantiate୦୦TimSort୦sort_slice_tim_go2୮aPerson {
	this := &instantiate୦୦TimSort୦sort_slice_tim_go2୮aPerson{
		a:         a,
		c:         c,
		minGallop: MIN_GALLOP,
	}

//line sort_slice_tim.go2:109
 l := len(a)
	tlen := 0
	if l < 2*INITIAL_TMP_STORAGE_LENGTH {
		tlen = l >> 1
	} else {
		tlen = INITIAL_TMP_STORAGE_LENGTH
	}
	if len(work) == 0 || workLen < tlen || workBase+tlen > len(work) {
		this.tmp = make([]Person, tlen)
		this.tmpBase = 0
		this.tmpLen = tlen
	} else {
		this.tmp = work
		this.tmpBase = workBase
		this.tmpLen = workLen
	}

//line sort_slice_tim.go2:140
 stackLen := 0
	if l < 120 {
		stackLen = 5
	} else if l < 1542 {
		stackLen = 10
	} else if l < 119151 {
		stackLen = 24
	} else {
		stackLen = 49
	}

	this.runBase = make([]int, stackLen)
	this.runLen = make([]int, stackLen)
	return this
}

//line sort_slice_tim.go2:327
func instantiate୦୦countRunAndMakeAscending୦int(a []int, lo int, hi int, c func(i, j int,) int) int {

	var runHi = lo + 1
	if runHi == hi {
		return 1
	}

//line sort_slice_tim.go2:335
 runHi++
	if c(a[runHi-1], a[lo]) < 0 {
		for runHi < hi && c(a[runHi], a[runHi-1]) < 0 {
			runHi++
		}
//line sort_slice_tim.go2:339
  instantiate୦୦reverseRange୦int(a, lo, runHi)
//line sort_slice_tim.go2:341
 } else {
		for runHi < hi && c(a[runHi], a[runHi-1]) >= 0 {
			runHi++
		}
	}

	return runHi - lo
}
//line sort_slice_tim.go2:252
func instantiate୦୦binarySort୦int(a []int, lo int, hi int, start int, c func(i, j int,) int) {

	if start == lo {
		start++
	}

	for ; start < hi; start++ {
					var pivot = a[start]

//line sort_slice_tim.go2:262
  var left = lo
					var right = start

//line sort_slice_tim.go2:270
  for left < right {
			var mid = (left + right) >> 1
			if c(pivot, a[mid]) < 0 {
				right = mid
			} else {
				left = mid + 1
			}
		}

//line sort_slice_tim.go2:287
  var n = start - left

		switch n {
		case 2:
			a[left+2] = a[left+1]
		case 1:
			a[left+1] = a[left]
		default:
			copy(a[left+1:], a[left:left+n])

		}
		a[left] = pivot
	}
}
//line sort_slice_tim.go2:101
func instantiate୦୦newTimSort୦int(a []int, c func(i, j int,) int, work []int, workBase int, workLen int) *instantiate୦୦TimSort୦int {
	this := &instantiate୦୦TimSort୦int{
		a:         a,
		c:         c,
		minGallop: MIN_GALLOP,
	}

//line sort_slice_tim.go2:109
 l := len(a)
	tlen := 0
	if l < 2*INITIAL_TMP_STORAGE_LENGTH {
		tlen = l >> 1
	} else {
		tlen = INITIAL_TMP_STORAGE_LENGTH
	}
	if len(work) == 0 || workLen < tlen || workBase+tlen > len(work) {
		this.tmp = make([]int, tlen)
		this.tmpBase = 0
		this.tmpLen = tlen
	} else {
		this.tmp = work
		this.tmpBase = workBase
		this.tmpLen = workLen
	}

//line sort_slice_tim.go2:140
 stackLen := 0
	if l < 120 {
		stackLen = 5
	} else if l < 1542 {
		stackLen = 10
	} else if l < 119151 {
		stackLen = 24
	} else {
		stackLen = 49
	}

	this.runBase = make([]int, stackLen)
	this.runLen = make([]int, stackLen)
	return this
}

//line sort_slice_tim.go2:357
func instantiate୦୦reverseRange୦sort_slice_tim_go2୮aPerson(a []Person, lo int, hi int) {
	hi--
	for lo < hi {
		var t = a[lo]
		a[lo] = a[hi]
		lo++
		a[hi] = t
		hi--
	}
}

//line sort_slice_tim.go2:366
type instantiate୦୦TimSort୦sort_slice_tim_go2୮aPerson struct {
//line sort_slice_tim.go2:54
 a []Person

//line sort_slice_tim.go2:59
 c func(i, j Person,) int

//line sort_slice_tim.go2:66
 minGallop int

//line sort_slice_tim.go2:73
 tmp []Person
				tmpBase int
				tmpLen  int

//line sort_slice_tim.go2:87
 stackSize int
	runBase []int
	runLen  []int
}

//line sort_slice_tim.go2:401
func (this *instantiate୦୦TimSort୦sort_slice_tim_go2୮aPerson,) pushRun(runBase int, runLen int) {
	this.runBase[this.stackSize] = runBase
	this.runLen[this.stackSize] = runLen
	this.stackSize++
}

//line sort_slice_tim.go2:423
func (this *instantiate୦୦TimSort୦sort_slice_tim_go2୮aPerson,) mergeCollapse() {
	for this.stackSize > 1 {
		var n = this.stackSize - 2
		if n > 0 && this.runLen[n-1] <= this.runLen[n]+this.runLen[n+1] || n > 1 && this.runLen[n-2] <= this.runLen[n]+this.runLen[n-1] {
			if this.runLen[n-1] < this.runLen[n+1] {
				n--
			}
		} else if n < 0 || this.runLen[n] > this.runLen[n+1] {
			break
		}
		this.mergeAt(n)
	}
}

//line sort_slice_tim.go2:441
func (this *instantiate୦୦TimSort୦sort_slice_tim_go2୮aPerson,) mergeForceCollapse() {
	for this.stackSize > 1 {
		var n = this.stackSize - 2
		if n > 0 && this.runLen[n-1] < this.runLen[n+1] {
			n--
		}
		this.mergeAt(n)
	}
}

//line sort_slice_tim.go2:458
func (this *instantiate୦୦TimSort୦sort_slice_tim_go2୮aPerson,) mergeAt(i int) {

//line sort_slice_tim.go2:463
 var base1 = this.runBase[i]
				var len1 = this.runLen[i]
				var base2 = this.runBase[i+1]
				var len2 = this.runLen[i+1]

//line sort_slice_tim.go2:475
 this.runLen[i] = len1 + len2
	if i == this.stackSize-3 {
		this.runBase[i+1] = this.runBase[i+2]
		this.runLen[i+1] = this.runLen[i+2]
	}
				this.stackSize--

//line sort_slice_tim.go2:486
 var k = instantiate୦୦gallopRight୦sort_slice_tim_go2୮aPerson(this.a[base2], this.a, base1, len1, 0, this.c)

	base1 += k
	len1 -= k
	if len1 == 0 {
		return
	}

//line sort_slice_tim.go2:498
 len2 = instantiate୦୦gallopLeft୦sort_slice_tim_go2୮aPerson(this.a[base1+len1-1], this.a, base2, len2, len2-1, this.c)

	if len2 == 0 {
		return
	}

//line sort_slice_tim.go2:505
 if len1 <= len2 {
		this.mergeLo(base1, len1, base2, len2)
	} else {
		this.mergeHi(base1, len1, base2, len2)
	}
}

//line sort_slice_tim.go2:682
func (this *instantiate୦୦TimSort୦sort_slice_tim_go2୮aPerson,) mergeLo(base1 int, len1 int, base2 int, len2 int) {

//line sort_slice_tim.go2:686
 var a = this.a
				var tmp []Person = this.ensureCapacity(len1)
				var cursor1 = this.tmpBase
				var cursor2 = base2
				var dest = base1

				copy(tmp[cursor1:], a[base1:base1+len1])

//line sort_slice_tim.go2:696
 a[dest] = a[cursor2]
	dest++
	cursor2++

	len2--
	if len2 == 0 {
		copy(a[dest:], tmp[cursor1:cursor1+len1])

		return
	}
	if len1 == 1 {
		copy(a[dest:], a[cursor2:cursor2+len2])

		a[dest+len2] = tmp[cursor1]
		return
	}

	var c = this.c
	var minGallop = this.minGallop
outer:
	for {
					var count1 = 0
					var count2 = 0

//line sort_slice_tim.go2:724
  for {

			if c(a[cursor2], tmp[cursor1]) < 0 {
				a[dest] = a[cursor2]
				dest++
				cursor2++

				count2++
				count1 = 0
				len2--
				if len2 == 0 {
					break outer
				}
			} else {
				a[dest] = tmp[cursor1]
				dest++
				cursor1++

				count1++
				count2 = 0
				len1--
				if len1 == 1 {
					break outer
				}
			}

			if (count1 | count2) < minGallop {
			} else {
				break
			}
		}

//line sort_slice_tim.go2:761
  for {

			count1 = instantiate୦୦gallopRight୦sort_slice_tim_go2୮aPerson(a[cursor2], tmp, cursor1, len1, 0, c)
			if count1 != 0 {
				copy(a[dest:], tmp[cursor1:cursor1+count1])

				dest += count1
				cursor1 += count1
				len1 -= count1
				if len1 <= 1 {
					break outer
				}
			}
			a[dest] = a[cursor2]
			dest++
			cursor2++

			len2--
			if len2 == 0 {
				break outer
			}

			count2 = instantiate୦୦gallopLeft୦sort_slice_tim_go2୮aPerson(tmp[cursor1], a, cursor2, len2, 0, c)
			if count2 != 0 {
				copy(a[dest:], a[cursor2:cursor2+count2])

				dest += count2
				cursor2 += count2
				len2 -= count2
				if len2 == 0 {
					break outer
				}
			}
			a[dest] = tmp[cursor1]
			dest++
			cursor1++

			len1--
			if len1 == 1 {
				break outer
			}
			minGallop--

			if count1 >= MIN_GALLOP || count2 >= MIN_GALLOP {
			} else {
				break
			}
		}
		if minGallop < 0 {
			minGallop = 0
		}
		minGallop += 2
	}
	if minGallop < 1 {
		this.minGallop = 1
	} else {
		this.minGallop = minGallop
	}

	if len1 == 1 {

		copy(a[dest:], a[cursor2:cursor2+len2])

		a[dest+len2] = tmp[cursor1]
	} else if len1 == 0 {
		panic("Comparison method violates its general contract!")
	} else {

//line sort_slice_tim.go2:830
  copy(a[dest:], tmp[cursor1:cursor1+len1])

	}
}

//line sort_slice_tim.go2:846
func (this *instantiate୦୦TimSort୦sort_slice_tim_go2୮aPerson,) mergeHi(base1, len1, base2, len2 int) {

//line sort_slice_tim.go2:850
 var a = this.a
				var tmp []Person = this.ensureCapacity(len2)
				var tmpBase = this.tmpBase

				copy(tmp[tmpBase:], a[base2:base2+len2])

//line sort_slice_tim.go2:857
 var cursor1 = base1 + len1 - 1
				var cursor2 = tmpBase + len2 - 1
				var dest = base2 + len2 - 1

//line sort_slice_tim.go2:862
 a[dest] = a[cursor1]
	dest--
	cursor1--

	len1--
	if len1 == 0 {
		copy(a[dest-(len2-1):], tmp[tmpBase:tmpBase+len2])

		return
	}
	if len2 == 1 {
		dest -= len1
		cursor1 -= len1
		copy(a[dest+1:], a[cursor1+1:cursor1+1+len1])

		a[dest] = tmp[cursor2]
		return
	}

	var c = this.c
	var minGallop = this.minGallop
outer:
	for true {
					var count1 = 0
					var count2 = 0

//line sort_slice_tim.go2:892
  for {

			if c(tmp[cursor2], a[cursor1]) < 0 {
				a[dest] = a[cursor1]
				dest--
				cursor1--

				count1++
				count2 = 0

				len1--
				if len1 == 0 {
					break outer
				}
			} else {
				a[dest] = tmp[cursor2]
				dest--
				cursor2--

				count2++
				count1 = 0

				len2--
				if len2 == 1 {
					break outer
				}
			}
			if (count1 | count2) < minGallop {
			} else {
				break
			}
		}

//line sort_slice_tim.go2:930
  for {

			count1 = len1 - instantiate୦୦gallopRight୦sort_slice_tim_go2୮aPerson(tmp[cursor2], a, base1, len1, len1-1, c)
			if count1 != 0 {
				dest -= count1
				cursor1 -= count1
				len1 -= count1
				copy(a[dest+1:], a[cursor1+1:cursor1+1+count1])

				if len1 == 0 {
					break outer
				}
			}
			a[dest] = tmp[cursor2]
			dest--
			cursor2--

			len2--
			if len2 == 1 {
				break outer
			}

			count2 = len2 - instantiate୦୦gallopLeft୦sort_slice_tim_go2୮aPerson(a[cursor1], tmp, tmpBase, len2, len2-1, c)
			if count2 != 0 {
				dest -= count2
				cursor2 -= count2
				len2 -= count2
				copy(a[dest+1:], tmp[cursor2+1:cursor2+1+count2])

				if len2 <= 1 {
					break outer
				}
			}
			a[dest] = a[cursor1]
			dest--
			cursor1--

			len1--
			if len1 == 0 {
				break outer
			}
			minGallop--
			if count1 >= MIN_GALLOP || count2 >= MIN_GALLOP {
			} else {
				break
			}
		}
		if minGallop < 0 {
			minGallop = 0
		}
		minGallop += 2
	}

//line sort_slice_tim.go2:984
 if minGallop < 1 {
		this.minGallop = 1
	} else {
		this.minGallop = minGallop
	}

	if len2 == 1 {

		dest -= len1
		cursor1 -= len1
		copy(a[dest+1:], a[cursor1+1:cursor1+1+len1])

		a[dest] = tmp[cursor2]
	} else if len2 == 0 {
		panic(
			"Comparison method violates its general contract!")
	} else {

//line sort_slice_tim.go2:1003
  copy(a[dest-(len2-1):], tmp[tmpBase:tmpBase+len2])

	}
}

//line sort_slice_tim.go2:1016
func (this *instantiate୦୦TimSort୦sort_slice_tim_go2୮aPerson,) ensureCapacity(minCapacity int) []Person {
	if this.tmpLen < minCapacity {

//line sort_slice_tim.go2:1020
  newSize := minCapacity
		newSize |= newSize >> 1
		newSize |= newSize >> 2
		newSize |= newSize >> 4
		newSize |= newSize >> 8
		newSize |= newSize >> 16
		newSize++

		if newSize < 0 {
			newSize = minCapacity
		} else {
			newSize = min(newSize, len(this.a)>>1)
		}

		this.tmp = make([]Person, newSize)
		this.tmpLen = newSize
		this.tmpBase = 0
	}
	return this.tmp
}
//line sort_slice_tim.go2:357
func instantiate୦୦reverseRange୦int(a []int, lo int, hi int) {
	hi--
	for lo < hi {
		var t = a[lo]
		a[lo] = a[hi]
		lo++
		a[hi] = t
		hi--
	}
}

//line sort_slice_tim.go2:366
type instantiate୦୦TimSort୦int struct {
//line sort_slice_tim.go2:54
 a []int

//line sort_slice_tim.go2:59
 c func(i, j int,) int

//line sort_slice_tim.go2:66
 minGallop int

//line sort_slice_tim.go2:73
 tmp []int
				tmpBase int
				tmpLen  int

//line sort_slice_tim.go2:87
 stackSize int
	runBase []int
	runLen  []int
}

//line sort_slice_tim.go2:401
func (this *instantiate୦୦TimSort୦int,) pushRun(runBase int, runLen int) {
	this.runBase[this.stackSize] = runBase
	this.runLen[this.stackSize] = runLen
	this.stackSize++
}

//line sort_slice_tim.go2:423
func (this *instantiate୦୦TimSort୦int,) mergeCollapse() {
	for this.stackSize > 1 {
		var n = this.stackSize - 2
		if n > 0 && this.runLen[n-1] <= this.runLen[n]+this.runLen[n+1] || n > 1 && this.runLen[n-2] <= this.runLen[n]+this.runLen[n-1] {
			if this.runLen[n-1] < this.runLen[n+1] {
				n--
			}
		} else if n < 0 || this.runLen[n] > this.runLen[n+1] {
			break
		}
		this.mergeAt(n)
	}
}

//line sort_slice_tim.go2:441
func (this *instantiate୦୦TimSort୦int,) mergeForceCollapse() {
	for this.stackSize > 1 {
		var n = this.stackSize - 2
		if n > 0 && this.runLen[n-1] < this.runLen[n+1] {
			n--
		}
		this.mergeAt(n)
	}
}

//line sort_slice_tim.go2:458
func (this *instantiate୦୦TimSort୦int,) mergeAt(i int) {

//line sort_slice_tim.go2:463
 var base1 = this.runBase[i]
				var len1 = this.runLen[i]
				var base2 = this.runBase[i+1]
				var len2 = this.runLen[i+1]

//line sort_slice_tim.go2:475
 this.runLen[i] = len1 + len2
	if i == this.stackSize-3 {
		this.runBase[i+1] = this.runBase[i+2]
		this.runLen[i+1] = this.runLen[i+2]
	}
				this.stackSize--

//line sort_slice_tim.go2:486
 var k = instantiate୦୦gallopRight୦int(this.a[base2], this.a, base1, len1, 0, this.c)

	base1 += k
	len1 -= k
	if len1 == 0 {
		return
	}

//line sort_slice_tim.go2:498
 len2 = instantiate୦୦gallopLeft୦int(this.a[base1+len1-1], this.a, base2, len2, len2-1, this.c)

	if len2 == 0 {
		return
	}

//line sort_slice_tim.go2:505
 if len1 <= len2 {
		this.mergeLo(base1, len1, base2, len2)
	} else {
		this.mergeHi(base1, len1, base2, len2)
	}
}

//line sort_slice_tim.go2:682
func (this *instantiate୦୦TimSort୦int,) mergeLo(base1 int, len1 int, base2 int, len2 int) {

//line sort_slice_tim.go2:686
 var a = this.a
				var tmp []int = this.ensureCapacity(len1)
				var cursor1 = this.tmpBase
				var cursor2 = base2
				var dest = base1

				copy(tmp[cursor1:], a[base1:base1+len1])

//line sort_slice_tim.go2:696
 a[dest] = a[cursor2]
	dest++
	cursor2++

	len2--
	if len2 == 0 {
		copy(a[dest:], tmp[cursor1:cursor1+len1])

		return
	}
	if len1 == 1 {
		copy(a[dest:], a[cursor2:cursor2+len2])

		a[dest+len2] = tmp[cursor1]
		return
	}

	var c = this.c
	var minGallop = this.minGallop
outer:
	for {
					var count1 = 0
					var count2 = 0

//line sort_slice_tim.go2:724
  for {

			if c(a[cursor2], tmp[cursor1]) < 0 {
				a[dest] = a[cursor2]
				dest++
				cursor2++

				count2++
				count1 = 0
				len2--
				if len2 == 0 {
					break outer
				}
			} else {
				a[dest] = tmp[cursor1]
				dest++
				cursor1++

				count1++
				count2 = 0
				len1--
				if len1 == 1 {
					break outer
				}
			}

			if (count1 | count2) < minGallop {
			} else {
				break
			}
		}

//line sort_slice_tim.go2:761
  for {

			count1 = instantiate୦୦gallopRight୦int(a[cursor2], tmp, cursor1, len1, 0, c)
			if count1 != 0 {
				copy(a[dest:], tmp[cursor1:cursor1+count1])

				dest += count1
				cursor1 += count1
				len1 -= count1
				if len1 <= 1 {
					break outer
				}
			}
			a[dest] = a[cursor2]
			dest++
			cursor2++

			len2--
			if len2 == 0 {
				break outer
			}

			count2 = instantiate୦୦gallopLeft୦int(tmp[cursor1], a, cursor2, len2, 0, c)
			if count2 != 0 {
				copy(a[dest:], a[cursor2:cursor2+count2])

				dest += count2
				cursor2 += count2
				len2 -= count2
				if len2 == 0 {
					break outer
				}
			}
			a[dest] = tmp[cursor1]
			dest++
			cursor1++

			len1--
			if len1 == 1 {
				break outer
			}
			minGallop--

			if count1 >= MIN_GALLOP || count2 >= MIN_GALLOP {
			} else {
				break
			}
		}
		if minGallop < 0 {
			minGallop = 0
		}
		minGallop += 2
	}
	if minGallop < 1 {
		this.minGallop = 1
	} else {
		this.minGallop = minGallop
	}

	if len1 == 1 {

		copy(a[dest:], a[cursor2:cursor2+len2])

		a[dest+len2] = tmp[cursor1]
	} else if len1 == 0 {
		panic("Comparison method violates its general contract!")
	} else {

//line sort_slice_tim.go2:830
  copy(a[dest:], tmp[cursor1:cursor1+len1])

	}
}

//line sort_slice_tim.go2:846
func (this *instantiate୦୦TimSort୦int,) mergeHi(base1, len1, base2, len2 int) {

//line sort_slice_tim.go2:850
 var a = this.a
				var tmp []int = this.ensureCapacity(len2)
				var tmpBase = this.tmpBase

				copy(tmp[tmpBase:], a[base2:base2+len2])

//line sort_slice_tim.go2:857
 var cursor1 = base1 + len1 - 1
				var cursor2 = tmpBase + len2 - 1
				var dest = base2 + len2 - 1

//line sort_slice_tim.go2:862
 a[dest] = a[cursor1]
	dest--
	cursor1--

	len1--
	if len1 == 0 {
		copy(a[dest-(len2-1):], tmp[tmpBase:tmpBase+len2])

		return
	}
	if len2 == 1 {
		dest -= len1
		cursor1 -= len1
		copy(a[dest+1:], a[cursor1+1:cursor1+1+len1])

		a[dest] = tmp[cursor2]
		return
	}

	var c = this.c
	var minGallop = this.minGallop
outer:
	for true {
					var count1 = 0
					var count2 = 0

//line sort_slice_tim.go2:892
  for {

			if c(tmp[cursor2], a[cursor1]) < 0 {
				a[dest] = a[cursor1]
				dest--
				cursor1--

				count1++
				count2 = 0

				len1--
				if len1 == 0 {
					break outer
				}
			} else {
				a[dest] = tmp[cursor2]
				dest--
				cursor2--

				count2++
				count1 = 0

				len2--
				if len2 == 1 {
					break outer
				}
			}
			if (count1 | count2) < minGallop {
			} else {
				break
			}
		}

//line sort_slice_tim.go2:930
  for {

			count1 = len1 - instantiate୦୦gallopRight୦int(tmp[cursor2], a, base1, len1, len1-1, c)
			if count1 != 0 {
				dest -= count1
				cursor1 -= count1
				len1 -= count1
				copy(a[dest+1:], a[cursor1+1:cursor1+1+count1])

				if len1 == 0 {
					break outer
				}
			}
			a[dest] = tmp[cursor2]
			dest--
			cursor2--

			len2--
			if len2 == 1 {
				break outer
			}

			count2 = len2 - instantiate୦୦gallopLeft୦int(a[cursor1], tmp, tmpBase, len2, len2-1, c)
			if count2 != 0 {
				dest -= count2
				cursor2 -= count2
				len2 -= count2
				copy(a[dest+1:], tmp[cursor2+1:cursor2+1+count2])

				if len2 <= 1 {
					break outer
				}
			}
			a[dest] = a[cursor1]
			dest--
			cursor1--

			len1--
			if len1 == 0 {
				break outer
			}
			minGallop--
			if count1 >= MIN_GALLOP || count2 >= MIN_GALLOP {
			} else {
				break
			}
		}
		if minGallop < 0 {
			minGallop = 0
		}
		minGallop += 2
	}

//line sort_slice_tim.go2:984
 if minGallop < 1 {
		this.minGallop = 1
	} else {
		this.minGallop = minGallop
	}

	if len2 == 1 {

		dest -= len1
		cursor1 -= len1
		copy(a[dest+1:], a[cursor1+1:cursor1+1+len1])

		a[dest] = tmp[cursor2]
	} else if len2 == 0 {
		panic(
			"Comparison method violates its general contract!")
	} else {

//line sort_slice_tim.go2:1003
  copy(a[dest-(len2-1):], tmp[tmpBase:tmpBase+len2])

	}
}

//line sort_slice_tim.go2:1016
func (this *instantiate୦୦TimSort୦int,) ensureCapacity(minCapacity int) []int {
	if this.tmpLen < minCapacity {

//line sort_slice_tim.go2:1020
  newSize := minCapacity
		newSize |= newSize >> 1
		newSize |= newSize >> 2
		newSize |= newSize >> 4
		newSize |= newSize >> 8
		newSize |= newSize >> 16
		newSize++

		if newSize < 0 {
			newSize = minCapacity
		} else {
			newSize = min(newSize, len(this.a)>>1)
		}

		this.tmp = make([]int, newSize)
		this.tmpLen = newSize
		this.tmpBase = 0
	}
	return this.tmp
}
//line sort_slice_tim.go2:604
func instantiate୦୦gallopRight୦sort_slice_tim_go2୮aPerson(key Person, a []Person, base int, len int, hint int, c func(i, j Person,) int) int {

//line sort_slice_tim.go2:607
 var ofs = 1
	var lastOfs = 0
	if c(key, a[base+hint]) < 0 {

		var maxOfs = hint + 1
		for ofs < maxOfs && c(key, a[base+hint-ofs]) < 0 {
			lastOfs = ofs
			ofs = (ofs << 1) + 1
			if ofs <= 0 {
				ofs = maxOfs
			}
		}
		if ofs > maxOfs {
			ofs = maxOfs
		}

//line sort_slice_tim.go2:624
  var tmp = lastOfs
		lastOfs = hint - ofs
		ofs = hint - tmp
	} else {

		var maxOfs = len - hint
		for ofs < maxOfs && c(key, a[base+hint+ofs]) >= 0 {
			lastOfs = ofs
			ofs = (ofs << 1) + 1
			if ofs <= 0 {
				ofs = maxOfs
			}
		}
		if ofs > maxOfs {
			ofs = maxOfs
		}

//line sort_slice_tim.go2:642
  lastOfs += hint
		ofs += hint
	}

//line sort_slice_tim.go2:652
 lastOfs++
	for lastOfs < ofs {
		var m = lastOfs + ((ofs - lastOfs) >> 1)

		if c(key, a[base+m]) < 0 {
			ofs = m
		} else {
			lastOfs = m + 1
		}
	}

	return ofs
}
//line sort_slice_tim.go2:530
func instantiate୦୦gallopLeft୦sort_slice_tim_go2୮aPerson(key Person, a []Person, base int, len int, hint int, c func(i, j Person,) int) int {

	var lastOfs = 0
	var ofs = 1
	if c(key, a[base+hint]) > 0 {

		var maxOfs = len - hint
		for ofs < maxOfs && c(key, a[base+hint+ofs]) > 0 {
			lastOfs = ofs
			ofs = (ofs << 1) + 1
			if ofs <= 0 {
				ofs = maxOfs
			}
		}
		if ofs > maxOfs {
			ofs = maxOfs
		}

//line sort_slice_tim.go2:549
  lastOfs += hint
		ofs += hint
	} else {

		var maxOfs = hint + 1
		for ofs < maxOfs && c(key, a[base+hint-ofs]) <= 0 {
			lastOfs = ofs
			ofs = (ofs << 1) + 1
			if ofs <= 0 {
				ofs = maxOfs
			}
		}
		if ofs > maxOfs {
			ofs = maxOfs
		}

//line sort_slice_tim.go2:566
  var tmp = lastOfs
		lastOfs = hint - ofs
		ofs = hint - tmp
	}

//line sort_slice_tim.go2:577
 lastOfs++
	for lastOfs < ofs {
		var m = lastOfs + ((ofs - lastOfs) >> 1)

		if c(key, a[base+m]) > 0 {
			lastOfs = m + 1
		} else {
			ofs = m
		}
	}

	return ofs
}

//line sort_slice_tim.go2:604
func instantiate୦୦gallopRight୦int(key int, a []int, base int, len int, hint int, c func(i, j int,) int) int {

//line sort_slice_tim.go2:607
 var ofs = 1
	var lastOfs = 0
	if c(key, a[base+hint]) < 0 {

		var maxOfs = hint + 1
		for ofs < maxOfs && c(key, a[base+hint-ofs]) < 0 {
			lastOfs = ofs
			ofs = (ofs << 1) + 1
			if ofs <= 0 {
				ofs = maxOfs
			}
		}
		if ofs > maxOfs {
			ofs = maxOfs
		}

//line sort_slice_tim.go2:624
  var tmp = lastOfs
		lastOfs = hint - ofs
		ofs = hint - tmp
	} else {

		var maxOfs = len - hint
		for ofs < maxOfs && c(key, a[base+hint+ofs]) >= 0 {
			lastOfs = ofs
			ofs = (ofs << 1) + 1
			if ofs <= 0 {
				ofs = maxOfs
			}
		}
		if ofs > maxOfs {
			ofs = maxOfs
		}

//line sort_slice_tim.go2:642
  lastOfs += hint
		ofs += hint
	}

//line sort_slice_tim.go2:652
 lastOfs++
	for lastOfs < ofs {
		var m = lastOfs + ((ofs - lastOfs) >> 1)

		if c(key, a[base+m]) < 0 {
			ofs = m
		} else {
			lastOfs = m + 1
		}
	}

	return ofs
}
//line sort_slice_tim.go2:530
func instantiate୦୦gallopLeft୦int(key int, a []int, base int, len int, hint int, c func(i, j int,) int) int {

	var lastOfs = 0
	var ofs = 1
	if c(key, a[base+hint]) > 0 {

		var maxOfs = len - hint
		for ofs < maxOfs && c(key, a[base+hint+ofs]) > 0 {
			lastOfs = ofs
			ofs = (ofs << 1) + 1
			if ofs <= 0 {
				ofs = maxOfs
			}
		}
		if ofs > maxOfs {
			ofs = maxOfs
		}

//line sort_slice_tim.go2:549
  lastOfs += hint
		ofs += hint
	} else {

		var maxOfs = hint + 1
		for ofs < maxOfs && c(key, a[base+hint-ofs]) <= 0 {
			lastOfs = ofs
			ofs = (ofs << 1) + 1
			if ofs <= 0 {
				ofs = maxOfs
			}
		}
		if ofs > maxOfs {
			ofs = maxOfs
		}

//line sort_slice_tim.go2:566
  var tmp = lastOfs
		lastOfs = hint - ofs
		ofs = hint - tmp
	}

//line sort_slice_tim.go2:577
 lastOfs++
	for lastOfs < ofs {
		var m = lastOfs + ((ofs - lastOfs) >> 1)

		if c(key, a[base+m]) > 0 {
			lastOfs = m + 1
		} else {
			ofs = m
		}
	}

	return ofs
}

//line sort_slice_tim.go2:589
type _ builtinsort.Float64Slice

//line sort_slice_tim.go2:589
var _ = fmt.Errorf
//line sort_slice_tim.go2:589
var _ = log.Fatal
//line sort_slice_tim.go2:589
var _ = rand.ExpFloat64
//line sort_slice_tim.go2:589
var _ = testing.AllocsPerRun

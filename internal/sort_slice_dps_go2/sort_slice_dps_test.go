// Code generated by go2go; DO NOT EDIT.


//line sort_slice_dps_test.go2:1
package sort_slice_dps

//line sort_slice_dps_test.go2:1
import (
//line sort_slice_dps_test.go2:7
 builtinsort "sort"
//line sort_slice_dps_test.go2:7
 "fmt"
//line sort_slice_dps_test.go2:7
 "log"
//line sort_slice_dps_test.go2:7
 "math/rand"
//line sort_slice_dps_test.go2:7
 "testing"
//line sort_slice_dps_test.go2:7
)

//line sort_slice_dps_test.go2:11
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
//line sort_slice_dps_test.go2:36
  instantiate୦୦Sort୦sort_slice_dps୮aPerson(persons, byAge)

//line sort_slice_dps_test.go2:40
  sorted := instantiate୦୦IsSorted୦sort_slice_dps୮aPerson(persons, byAge)

		if !sorted {
			log.Panicf("should be sorted: %d", i)
		}
	}
}

var benchmarkSizes = []int{256, 1024, 4192, 16768}

func BenchmarkGo2VSSort(t *testing.B) {
	for _, size := range benchmarkSizes {
		var data = make([]Person, size)
		prepare(data)

		t.Run(fmt.Sprintf("DPS-%d", size), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
								t.StopTimer()
								dup := make([]Person, size)
								copy(dup, data)
								t.StartTimer()
//line sort_slice_dps_test.go2:60
    instantiate୦୦Sort୦sort_slice_dps୮aPerson(dup, func(o1, o2 Person) int {
//line sort_slice_dps_test.go2:63
     return o1.age - o2.age
				})
			}
		})

		t.Run(fmt.Sprintf("BuiltinSore-%d", size), func(t *testing.B) {
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

		t.Run(fmt.Sprintf("DPS-%d", size), func(t *testing.B) {
			for i := 0; i < t.N; i++ {
								t.StopTimer()
								dup := make([]int, size)
								copy(dup, data)
								t.StartTimer()
//line sort_slice_dps_test.go2:99
    instantiate୦୦Sort୦int(dup, func(i, j int) int { return i - j })
//line sort_slice_dps_test.go2:101
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
//line sort_slice_dps.go2:20
func instantiate୦୦Sort୦sort_slice_dps୮aPerson(a []Person, compare func(o1, o2 Person,) int) {
//line sort_slice_dps.go2:20
 instantiate୦୦sort୦sort_slice_dps୮aPerson(compare, a, 0, len(a)-1, nil, 0, 0)
//line sort_slice_dps.go2:22
}

func instantiate୦୦IsSorted୦sort_slice_dps୮aPerson(a []Person, compare func(o1, o2 Person,) int) bool {
	n := len(a)
	for i := n - 1; i > 0; i-- {
		if compare(a[i], a[i-1]) < 0 {
			return false
		}
	}
	return true
}
//line sort_slice_dps.go2:20
func instantiate୦୦Sort୦int(a []int, compare func(o1, o2 int,) int) {
//line sort_slice_dps.go2:20
 instantiate୦୦sort୦int(compare, a, 0, len(a)-1, nil, 0, 0)
//line sort_slice_dps.go2:22
}

//line sort_slice_dps.go2:45
func instantiate୦୦sort୦sort_slice_dps୮aPerson(compare func(o1, o2 Person,) int, a []Person, left int, right int, work []Person, workBase int, workLen int) {

	if right-left < quicksortThreshold {
//line sort_slice_dps.go2:47
  instantiate୦୦sortInternal୦sort_slice_dps୮aPerson(compare, a, left, right, true)
//line sort_slice_dps.go2:49
  return
	}

//line sort_slice_dps.go2:56
 var run = make([]int, maxRunCount+1)
				var count = 0
				run[0] = left

//line sort_slice_dps.go2:61
 for k := left; k < right; run[count] = k {

		for k < right && compare(a[k], a[k+1]) == 0 {
			k++
		}
		if k == right {
			break
		}
		if compare(a[k], a[k+1]) < 0 {
			for {
				k++
				if k <= right && compare(a[k-1], a[k]) <= 0 {
				} else {
					break
				}
			}

		} else if compare(a[k], a[k+1]) > 0 {
			for {
				k++
				if k <= right && compare(a[k-1], a[k]) >= 0 {
				} else {
					break
				}
			}

			lo := run[count] - 1
			for hi := k; lo+1 < hi-1; {
				lo++
				hi--
				a[lo], a[hi] = a[hi], a[lo]
			}
		}

//line sort_slice_dps.go2:97
  if run[count] > left && compare(a[run[count]], a[run[count]-1]) >= 0 {
			count--
		}

//line sort_slice_dps.go2:105
  count++
		if count == maxRunCount {
//line sort_slice_dps.go2:106
   instantiate୦୦sortInternal୦sort_slice_dps୮aPerson(compare, a, left, right, true)
//line sort_slice_dps.go2:108
   return
		}
	}

//line sort_slice_dps.go2:116
 if count == 0 {

		return
	} else if count == 1 && run[count] > right {

//line sort_slice_dps.go2:123
  return
	}
	right++
	if run[count] < right {

//line sort_slice_dps.go2:131
  count++
		run[count] = right
	}

//line sort_slice_dps.go2:136
 var odd byte = 0
	for n := 1; ; {
		n <<= 1
		if n < count {
			odd ^= 1
		} else {
			break
		}
	}

//line sort_slice_dps.go2:147
 var b []Person
	var ao, bo int
	var blen = right - left
	if len(work) == 0 || workLen < blen || workBase+blen > len(work) {
		work = make([]Person, blen)
		workBase = 0
	}
	if odd == 0 {
		copy(work[workBase:], a[left:left+blen])

		b = a
		bo = 0
		a = work
		ao = workBase - left
	} else {
		b = work
		ao = 0
		bo = workBase - left
	}

//line sort_slice_dps.go2:168
 for last := 0; count > 1; count = last {
		last = 0
		for k := 2; k <= count; k += 2 {
			var hi = run[k]
			var mi = run[k-1]

			i := run[k-2]
			p := i
			q := mi
			for ; i < hi; i++ {
				if q >= hi || p < mi && compare(a[p+ao], a[q+ao]) <= 0 {

					b[i+bo] = a[p+ao]
					p++
				} else {

					b[i+bo] = a[q+ao]
					q++
				}
			}
			last++
			run[last] = hi
		}
		if (count & 1) != 0 {
			i := right
			lo := run[count-1]
			for ; i-1 >= lo; b[i+bo] = a[i+ao] {
				i--
			}
			last++
			run[last] = right
		}
		var t = a
		a = b
		b = t
		var o = ao
		ao = bo
		bo = o
	}
}
//line sort_slice_dps.go2:45
func instantiate୦୦sort୦int(compare func(o1, o2 int,) int, a []int, left int, right int, work []int, workBase int, workLen int) {

	if right-left < quicksortThreshold {
//line sort_slice_dps.go2:47
  instantiate୦୦sortInternal୦int(compare, a, left, right, true)
//line sort_slice_dps.go2:49
  return
	}

//line sort_slice_dps.go2:56
 var run = make([]int, maxRunCount+1)
				var count = 0
				run[0] = left

//line sort_slice_dps.go2:61
 for k := left; k < right; run[count] = k {

		for k < right && compare(a[k], a[k+1]) == 0 {
			k++
		}
		if k == right {
			break
		}
		if compare(a[k], a[k+1]) < 0 {
			for {
				k++
				if k <= right && compare(a[k-1], a[k]) <= 0 {
				} else {
					break
				}
			}

		} else if compare(a[k], a[k+1]) > 0 {
			for {
				k++
				if k <= right && compare(a[k-1], a[k]) >= 0 {
				} else {
					break
				}
			}

			lo := run[count] - 1
			for hi := k; lo+1 < hi-1; {
				lo++
				hi--
				a[lo], a[hi] = a[hi], a[lo]
			}
		}

//line sort_slice_dps.go2:97
  if run[count] > left && compare(a[run[count]], a[run[count]-1]) >= 0 {
			count--
		}

//line sort_slice_dps.go2:105
  count++
		if count == maxRunCount {
//line sort_slice_dps.go2:106
   instantiate୦୦sortInternal୦int(compare, a, left, right, true)
//line sort_slice_dps.go2:108
   return
		}
	}

//line sort_slice_dps.go2:116
 if count == 0 {

		return
	} else if count == 1 && run[count] > right {

//line sort_slice_dps.go2:123
  return
	}
	right++
	if run[count] < right {

//line sort_slice_dps.go2:131
  count++
		run[count] = right
	}

//line sort_slice_dps.go2:136
 var odd byte = 0
	for n := 1; ; {
		n <<= 1
		if n < count {
			odd ^= 1
		} else {
			break
		}
	}

//line sort_slice_dps.go2:147
 var b []int
	var ao, bo int
	var blen = right - left
	if len(work) == 0 || workLen < blen || workBase+blen > len(work) {
		work = make([]int, blen)
		workBase = 0
	}
	if odd == 0 {
		copy(work[workBase:], a[left:left+blen])

		b = a
		bo = 0
		a = work
		ao = workBase - left
	} else {
		b = work
		ao = 0
		bo = workBase - left
	}

//line sort_slice_dps.go2:168
 for last := 0; count > 1; count = last {
		last = 0
		for k := 2; k <= count; k += 2 {
			var hi = run[k]
			var mi = run[k-1]

			i := run[k-2]
			p := i
			q := mi
			for ; i < hi; i++ {
				if q >= hi || p < mi && compare(a[p+ao], a[q+ao]) <= 0 {

					b[i+bo] = a[p+ao]
					p++
				} else {

					b[i+bo] = a[q+ao]
					q++
				}
			}
			last++
			run[last] = hi
		}
		if (count & 1) != 0 {
			i := right
			lo := run[count-1]
			for ; i-1 >= lo; b[i+bo] = a[i+ao] {
				i--
			}
			last++
			run[last] = right
		}
		var t = a
		a = b
		b = t
		var o = ao
		ao = bo
		bo = o
	}
}

//line sort_slice_dps.go2:217
func instantiate୦୦sortInternal୦sort_slice_dps୮aPerson(compare func(o1, o2 Person,) int, a []Person, left int, right int, leftmost bool) {
				var length = right - left + 1

//line sort_slice_dps.go2:221
 if length < insertionSortThreshold {
		if leftmost {

//line sort_slice_dps.go2:228
   i := left
			j := i
			for i < right {
				var ai = a[i+1]
				for compare(ai, a[j]) < 0 {
					a[j+1] = a[j]
					j--
					if j+1 == left {
						break
					}
				}
				a[j+1] = ai

				i++
				j = i
			}
		} else {

//line sort_slice_dps.go2:249
   for {
				if left >= right {
					return
				}

				left++
				if compare(a[left], a[left-1]) >= 0 {
				} else {
					break
				}
			}

//line sort_slice_dps.go2:269
   k := left
			for {
				left++
				if left <= right {

				} else {
					break
				}

				var a1 = a[k]
				var a2 = a[left]

				if compare(a1, a2) < 0 {
					a2 = a1
					a1 = a[left]
				}

				for {
					k--
					if compare(a1, a[k]) < 0 {
					} else {
						break
					}
					a[k+2] = a[k]
				}
				k++
				a[k+1] = a1

				for {
					k--
					if compare(a2, a[k]) < 0 {
					} else {
						break
					}
					a[k+1] = a[k]
				}
				a[k+1] = a2
				left++
				k = left
			}
			var last = a[right]

			for {
				right--
				if compare(last, a[right]) < 0 {
				} else {
					break
				}
				a[right+1] = a[right]
			}
			a[right+1] = last
		}
		return
	}

//line sort_slice_dps.go2:325
 var seventh = (length >> 3) + (length >> 6) + 1

//line sort_slice_dps.go2:334
 var e3 = (left + right) >> 1
				var e2 = e3 - seventh
				var e1 = e2 - seventh
				var e4 = e3 + seventh
				var e5 = e4 + seventh

//line sort_slice_dps.go2:341
 if compare(a[e2], a[e1]) < 0 {
		a[e2], a[e1] = a[e1], a[e2]
	}

	if compare(a[e3], a[e2]) < 0 {
		var t = a[e3]
		a[e3], a[e2] = a[e2], a[e3]
		if compare(t, a[e1]) < 0 {
			a[e2] = a[e1]
			a[e1] = t
		}
	}
	if compare(a[e4], a[e3]) < 0 {
		var t = a[e4]
		a[e4], a[e3] = a[e3], a[e4]
		if compare(t, a[e2]) < 0 {
			a[e3] = a[e2]
			a[e2] = t
			if compare(t, a[e1]) < 0 {
				a[e2] = a[e1]
				a[e1] = t
			}
		}
	}
	if compare(a[e5], a[e4]) < 0 {
		var t = a[e5]
		a[e5], a[e4] = a[e4], a[e5]
		if compare(t, a[e3]) < 0 {
			a[e4] = a[e3]
			a[e3] = t
			if compare(t, a[e2]) < 0 {
				a[e3] = a[e2]
				a[e2] = t
				if compare(t, a[e1]) < 0 {
					a[e2] = a[e1]
					a[e1] = t
				}
			}
		}
	}

//line sort_slice_dps.go2:383
 var less = left
	var great = right

	if compare(a[e1], a[e2]) != 0 && compare(a[e2], a[e3]) != 0 && compare(a[e3], a[e4]) != 0 && compare(a[e4], a[e5]) != 0 {

//line sort_slice_dps.go2:392
  var pivot1 = a[e2]
					var pivot2 = a[e4]

//line sort_slice_dps.go2:401
  a[e2] = a[left]
					a[e4] = a[right]

//line sort_slice_dps.go2:407
  for {
			less++
			if compare(a[less], pivot1) < 0 {
			} else {
				break
			}
		}
		for {
			great--
			if compare(a[great], pivot2) > 0 {
			} else {
				break
			}
		}

//line sort_slice_dps.go2:441
 Outer:
		for k := less - 1; ; {
			k++
			if k <= great {
			} else {
				break
			}
			var ak = a[k]
			if compare(ak, pivot1) < 0 {
							a[k] = a[less]

//line sort_slice_dps.go2:455
    a[less] = ak
				less++
			} else if compare(ak, pivot2) > 0 {
				for compare(a[great], pivot2) > 0 {
					great--
					if great+1 == k {
						break Outer
					}
				}
				if compare(a[great], pivot1) < 0 {
					a[k] = a[less]
					a[less] = a[great]
					less++
				} else {
					a[k] = a[great]
				}

//line sort_slice_dps.go2:475
    a[great] = ak
				great--
			}
		}

//line sort_slice_dps.go2:481
  a[left] = a[less-1]
					a[less-1] = pivot1
					a[right] = a[great+1]
					a[great+1] = pivot2
//line sort_slice_dps.go2:484
  instantiate୦୦sortInternal୦sort_slice_dps୮aPerson(compare, a, left, less-2, leftmost)
//line sort_slice_dps.go2:487
  instantiate୦୦sortInternal୦sort_slice_dps୮aPerson(compare, a, great+2, right, false)

//line sort_slice_dps.go2:494
  if less < e1 && e5 < great {

//line sort_slice_dps.go2:498
   for compare(a[less], pivot1) == 0 {
				less++
			}

			for compare(a[great], pivot2) == 0 {
				great--
			}

//line sort_slice_dps.go2:525
  outer2:
			for k := less - 1; ; {
				k++
				if k <= great {
				} else {
					break
				}
				var ak = a[k]
				if compare(ak, pivot1) == 0 {
					a[k] = a[less]
					a[less] = ak
					less++
				} else if compare(ak, pivot2) == 0 {
					for compare(a[great], pivot2) == 0 {
						great--
						if great+1 == k {
							break outer2
						}
					}
					if compare(a[great], pivot1) == 0 {
									a[k] = a[less]

//line sort_slice_dps.go2:554
      a[less] = pivot1
						less++
					} else {
						a[k] = a[great]
					}
					a[great] = ak
					great--
				}
			}
		}
//line sort_slice_dps.go2:563
  instantiate୦୦sortInternal୦sort_slice_dps୮aPerson(compare, a, less, great, false)

//line sort_slice_dps.go2:568
 } else {

//line sort_slice_dps.go2:573
  var pivot = a[e3]

//line sort_slice_dps.go2:595
  for k := less; k <= great; k++ {
			if compare(a[k], pivot) == 0 {
				continue
			}
			var ak = a[k]
			if compare(ak, pivot) < 0 {
				a[k] = a[less]
				a[less] = ak
				less++
			} else {
				for compare(a[great], pivot) > 0 {
					great--
				}
				if compare(a[great], pivot) < 0 {
					a[k] = a[less]
					a[less] = a[great]
					less++
				} else {

//line sort_slice_dps.go2:621
     a[k] = pivot
				}
				a[great] = ak
				great--
			}
		}
//line sort_slice_dps.go2:626
  instantiate୦୦sortInternal୦sort_slice_dps୮aPerson(compare, a, left, less-1, leftmost)
//line sort_slice_dps.go2:633
  instantiate୦୦sortInternal୦sort_slice_dps୮aPerson(compare, a, great+1, right, false)
//line sort_slice_dps.go2:635
 }
}
//line sort_slice_dps.go2:217
func instantiate୦୦sortInternal୦int(compare func(o1, o2 int,) int, a []int, left int, right int, leftmost bool) {
				var length = right - left + 1

//line sort_slice_dps.go2:221
 if length < insertionSortThreshold {
		if leftmost {

//line sort_slice_dps.go2:228
   i := left
			j := i
			for i < right {
				var ai = a[i+1]
				for compare(ai, a[j]) < 0 {
					a[j+1] = a[j]
					j--
					if j+1 == left {
						break
					}
				}
				a[j+1] = ai

				i++
				j = i
			}
		} else {

//line sort_slice_dps.go2:249
   for {
				if left >= right {
					return
				}

				left++
				if compare(a[left], a[left-1]) >= 0 {
				} else {
					break
				}
			}

//line sort_slice_dps.go2:269
   k := left
			for {
				left++
				if left <= right {

				} else {
					break
				}

				var a1 = a[k]
				var a2 = a[left]

				if compare(a1, a2) < 0 {
					a2 = a1
					a1 = a[left]
				}

				for {
					k--
					if compare(a1, a[k]) < 0 {
					} else {
						break
					}
					a[k+2] = a[k]
				}
				k++
				a[k+1] = a1

				for {
					k--
					if compare(a2, a[k]) < 0 {
					} else {
						break
					}
					a[k+1] = a[k]
				}
				a[k+1] = a2
				left++
				k = left
			}
			var last = a[right]

			for {
				right--
				if compare(last, a[right]) < 0 {
				} else {
					break
				}
				a[right+1] = a[right]
			}
			a[right+1] = last
		}
		return
	}

//line sort_slice_dps.go2:325
 var seventh = (length >> 3) + (length >> 6) + 1

//line sort_slice_dps.go2:334
 var e3 = (left + right) >> 1
				var e2 = e3 - seventh
				var e1 = e2 - seventh
				var e4 = e3 + seventh
				var e5 = e4 + seventh

//line sort_slice_dps.go2:341
 if compare(a[e2], a[e1]) < 0 {
		a[e2], a[e1] = a[e1], a[e2]
	}

	if compare(a[e3], a[e2]) < 0 {
		var t = a[e3]
		a[e3], a[e2] = a[e2], a[e3]
		if compare(t, a[e1]) < 0 {
			a[e2] = a[e1]
			a[e1] = t
		}
	}
	if compare(a[e4], a[e3]) < 0 {
		var t = a[e4]
		a[e4], a[e3] = a[e3], a[e4]
		if compare(t, a[e2]) < 0 {
			a[e3] = a[e2]
			a[e2] = t
			if compare(t, a[e1]) < 0 {
				a[e2] = a[e1]
				a[e1] = t
			}
		}
	}
	if compare(a[e5], a[e4]) < 0 {
		var t = a[e5]
		a[e5], a[e4] = a[e4], a[e5]
		if compare(t, a[e3]) < 0 {
			a[e4] = a[e3]
			a[e3] = t
			if compare(t, a[e2]) < 0 {
				a[e3] = a[e2]
				a[e2] = t
				if compare(t, a[e1]) < 0 {
					a[e2] = a[e1]
					a[e1] = t
				}
			}
		}
	}

//line sort_slice_dps.go2:383
 var less = left
	var great = right

	if compare(a[e1], a[e2]) != 0 && compare(a[e2], a[e3]) != 0 && compare(a[e3], a[e4]) != 0 && compare(a[e4], a[e5]) != 0 {

//line sort_slice_dps.go2:392
  var pivot1 = a[e2]
					var pivot2 = a[e4]

//line sort_slice_dps.go2:401
  a[e2] = a[left]
					a[e4] = a[right]

//line sort_slice_dps.go2:407
  for {
			less++
			if compare(a[less], pivot1) < 0 {
			} else {
				break
			}
		}
		for {
			great--
			if compare(a[great], pivot2) > 0 {
			} else {
				break
			}
		}

//line sort_slice_dps.go2:441
 Outer:
		for k := less - 1; ; {
			k++
			if k <= great {
			} else {
				break
			}
			var ak = a[k]
			if compare(ak, pivot1) < 0 {
							a[k] = a[less]

//line sort_slice_dps.go2:455
    a[less] = ak
				less++
			} else if compare(ak, pivot2) > 0 {
				for compare(a[great], pivot2) > 0 {
					great--
					if great+1 == k {
						break Outer
					}
				}
				if compare(a[great], pivot1) < 0 {
					a[k] = a[less]
					a[less] = a[great]
					less++
				} else {
					a[k] = a[great]
				}

//line sort_slice_dps.go2:475
    a[great] = ak
				great--
			}
		}

//line sort_slice_dps.go2:481
  a[left] = a[less-1]
					a[less-1] = pivot1
					a[right] = a[great+1]
					a[great+1] = pivot2
//line sort_slice_dps.go2:484
  instantiate୦୦sortInternal୦int(compare, a, left, less-2, leftmost)
//line sort_slice_dps.go2:487
  instantiate୦୦sortInternal୦int(compare, a, great+2, right, false)

//line sort_slice_dps.go2:494
  if less < e1 && e5 < great {

//line sort_slice_dps.go2:498
   for compare(a[less], pivot1) == 0 {
				less++
			}

			for compare(a[great], pivot2) == 0 {
				great--
			}

//line sort_slice_dps.go2:525
  outer2:
			for k := less - 1; ; {
				k++
				if k <= great {
				} else {
					break
				}
				var ak = a[k]
				if compare(ak, pivot1) == 0 {
					a[k] = a[less]
					a[less] = ak
					less++
				} else if compare(ak, pivot2) == 0 {
					for compare(a[great], pivot2) == 0 {
						great--
						if great+1 == k {
							break outer2
						}
					}
					if compare(a[great], pivot1) == 0 {
									a[k] = a[less]

//line sort_slice_dps.go2:554
      a[less] = pivot1
						less++
					} else {
						a[k] = a[great]
					}
					a[great] = ak
					great--
				}
			}
		}
//line sort_slice_dps.go2:563
  instantiate୦୦sortInternal୦int(compare, a, less, great, false)

//line sort_slice_dps.go2:568
 } else {

//line sort_slice_dps.go2:573
  var pivot = a[e3]

//line sort_slice_dps.go2:595
  for k := less; k <= great; k++ {
			if compare(a[k], pivot) == 0 {
				continue
			}
			var ak = a[k]
			if compare(ak, pivot) < 0 {
				a[k] = a[less]
				a[less] = ak
				less++
			} else {
				for compare(a[great], pivot) > 0 {
					great--
				}
				if compare(a[great], pivot) < 0 {
					a[k] = a[less]
					a[less] = a[great]
					less++
				} else {

//line sort_slice_dps.go2:621
     a[k] = pivot
				}
				a[great] = ak
				great--
			}
		}
//line sort_slice_dps.go2:626
  instantiate୦୦sortInternal୦int(compare, a, left, less-1, leftmost)
//line sort_slice_dps.go2:633
  instantiate୦୦sortInternal୦int(compare, a, great+1, right, false)
//line sort_slice_dps.go2:635
 }
}

//line sort_slice_dps.go2:636
type _ builtinsort.Float64Slice

//line sort_slice_dps.go2:636
var _ = fmt.Errorf
//line sort_slice_dps.go2:636
var _ = log.Fatal
//line sort_slice_dps.go2:636
var _ = rand.ExpFloat64
//line sort_slice_dps.go2:636
var _ = testing.AllocsPerRun
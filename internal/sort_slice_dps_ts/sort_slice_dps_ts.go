package sort_slice_dps_ts

import "reflect"

/**
 * The maximum number of runs in merge sort.
 */
const maxRunCount = 67

/**
 * If the length of an array to be sorted is less than this
 * constant, Quicksort is used in preference to merge sort.
 */
const quicksortThreshold = 286

/**
 * If the length of an array to be sorted is less than this
 * constant, insertion sort is used in preference to Quicksort.
 */
const insertionSortThreshold = 47

type CompareFunc = func(o1, o2 interface{}) int

func Sort(a interface{}, compare CompareFunc) {
	rt := reflect.TypeOf(a)
	rv := reflect.ValueOf(a)
	elm := rt.Elem()

	n := rv.Len()

	interfaces := make([]interface{}, n)
	for i := 0; i < n; i++ {
		interfaces[i] = rv.Index(i).Interface()
	}

	SortInterfaces(interfaces, compare)

	for i := 0; i < n; i++ {
		rv.Index(i).Set(reflect.ValueOf(interfaces[i]).Convert(elm))
	}
}

func SortInterfaces(a []interface{}, compare CompareFunc) {
	sort(compare, a, 0, len(a)-1, nil, 0, 0)
}

func IsSorted(a interface{}, compare CompareFunc) bool {
	rv := reflect.ValueOf(a)
	n := rv.Len()
	for i := n - 1; i > 0; i-- {
		if compare(rv.Index(i).Interface(), rv.Index(i-1).Interface()) < 0 {
			return false
		}
	}
	return true
}

/**
 * Sorts the specified range of the array using the given
 * workspace array slice if possible for merging
 *
 * @param a the array to be sorted
 * @param left the index of the first element, inclusive, to be sorted
 * @param right the index of the last element, inclusive, to be sorted
 * @param work a workspace array (slice)
 * @param workBase origin of usable space in work array
 * @param workLen usable size of work array
 */
func sort(compare CompareFunc, a []interface{}, left int, right int, work []interface{}, workBase int, workLen int) {
	// Use Quicksort on small arrays
	if right-left < quicksortThreshold {
		sortInternal(compare, a, left, right, true)
		return
	}

	/*
	 * Index run[i] is the start of i-th run
	 * (ascending or descending sequence).
	 */
	var run = make([]int, maxRunCount+1)
	var count = 0
	run[0] = left

	// Check if the array is nearly sorted
	for k := left; k < right; run[count] = k {
		// Equal items in the beginning of the sequence
		for k < right && compare(a[k], a[k+1]) == 0 {
			k++
		}
		if k == right {
			break
		} // Sequence finishes with equal items
		if compare(a[k], a[k+1]) < 0 { // ascending
			for {
				k++
				if k <= right && compare(a[k-1], a[k]) <= 0 {
				} else {
					break
				}
			}

		} else if compare(a[k], a[k+1]) > 0 { // descending
			for {
				k++
				if k <= right && compare(a[k-1], a[k]) >= 0 {
				} else {
					break
				}
			}
			// Transform into an ascending sequence
			lo := run[count] - 1
			for hi := k; lo+1 < hi-1; {
				lo++
				hi--
				a[lo], a[hi] = a[hi], a[lo]
			}
		}

		// Merge a transformed descending sequence followed by an
		// ascending sequence
		if run[count] > left && compare(a[run[count]], a[run[count]-1]) >= 0 {
			count--
		}

		/*
		 * The array is not highly structured,
		 * use Quicksort instead of merge sort.
		 */
		count++
		if count == maxRunCount {
			sortInternal(compare, a, left, right, true)
			return
		}
	}

	// These invariants should hold true:
	//    run[0] = 0
	//    run[<last>] = right + 1; (terminator)

	if count == 0 {
		// A single equal run
		return
	} else if count == 1 && run[count] > right {
		// Either a single ascending or a transformed descending run.
		// Always check that a final run is a proper terminator, otherwise
		// we have an unterminated trailing run, to handle downstream.
		return
	}
	right++
	if run[count] < right {
		// Corner case: the final run is not a terminator. This may happen
		// if a final run is an equals run, or there is a single-element run
		// at the end. Fix up by adding a proper terminator at the end.
		// Note that we terminate with (right + 1), incremented earlier.
		count++
		run[count] = right
	}

	// Determine alternation base for merge
	var odd byte = 0
	for n := 1; ; {
		n <<= 1
		if n < count {
			odd ^= 1
		} else {
			break
		}
	}

	// Use or create temporary array b for merging
	var b []interface{}     // temp array; alternates with a
	var ao, bo int          // array offsets from 'left'
	var blen = right - left // space needed for b
	if len(work) == 0 || workLen < blen || workBase+blen > len(work) {
		work = make([]interface{}, blen)
		workBase = 0
	}
	if odd == 0 {
		copy(work[workBase:], a[left:left+blen])
		//System.arraycopy(a, left, work, workBase, blen);
		b = a
		bo = 0
		a = work
		ao = workBase - left
	} else {
		b = work
		ao = 0
		bo = workBase - left
	}

	// Merging
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
					//b[i + bo] = a[p++ + ao];
					b[i+bo] = a[p+ao]
					p++
				} else {
					//b[i + bo] = a[q++ + ao];
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

/**
 * Sorts the specified range of the array by Dual-Pivot Quicksort.
 *
 * @param a the array to be sorted
 * @param left the index of the first element, inclusive, to be sorted
 * @param right the index of the last element, inclusive, to be sorted
 * @param leftmost indicates if this part is the leftmost in the range
 */
func sortInternal(compare CompareFunc, a []interface{}, left int, right int, leftmost bool) {
	var length = right - left + 1

	// Use insertion sort on tiny arrays
	if length < insertionSortThreshold {
		if leftmost {
			/*
			 * Traditional (without sentinel) insertion sort,
			 * optimized for server VM, is used in case of
			 * the leftmost part.
			 */
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
			/*
			 * Skip the longest ascending sequence.
			 */

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

			/*
			 * Every element from adjoining part plays the role
			 * of sentinel, therefore this allows us to avoid the
			 * left range check on each iteration. Moreover, we use
			 * the more optimized algorithm, so called pair insertion
			 * sort, which is faster (in the context of Quicksort)
			 * than traditional implementation of insertion sort.
			 */
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

	// Inexpensive approximation of length / 7
	var seventh = (length >> 3) + (length >> 6) + 1

	/*
	 * Sort five evenly spaced elements around (and including) the
	 * center element in the range. These elements will be used for
	 * pivot selection as described below. The choice for spacing
	 * these elements was empirically determined to work well on
	 * a wide variety of inputs.
	 */
	var e3 = int(uint(left+right) >> 1) // The midpoint
	var e2 = e3 - seventh
	var e1 = e2 - seventh
	var e4 = e3 + seventh
	var e5 = e4 + seventh

	// Sort these elements using insertion sort
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

	// Pointers
	var less = left   // The index of the first element of center part
	var great = right // The index before the first element of right part

	if compare(a[e1], a[e2]) != 0 && compare(a[e2], a[e3]) != 0 && compare(a[e3], a[e4]) != 0 && compare(a[e4], a[e5]) != 0 {
		/*
		 * Use the second and fourth of the five sorted elements as pivots.
		 * These values are inexpensive approximations of the first and
		 * second terciles of the array. Note that pivot1 <= pivot2.
		 */
		var pivot1 = a[e2]
		var pivot2 = a[e4]

		/*
		 * The first and the last elements to be sorted are moved to the
		 * locations formerly occupied by the pivots. When partitioning
		 * is complete, the pivots are swapped back into their final
		 * positions, and excluded from subsequent sorting.
		 */
		a[e2] = a[left]
		a[e4] = a[right]

		/*
		 * Skip elements, which are less or greater than pivot values.
		 */
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

		/*
		 * Partitioning:
		 *
		 *   left part           center part                   right part
		 * +--------------------------------------------------------------+
		 * |  < pivot1  |  pivot1 <= && <= pivot2  |    ?    |  > pivot2  |
		 * +--------------------------------------------------------------+
		 *               ^                          ^       ^
		 *               |                          |       |
		 *              less                        k     great
		 *
		 * Invariants:
		 *
		 *              all in (left, less)   < pivot1
		 *    pivot1 <= all in [less, k)     <= pivot2
		 *              all in (great, right) > pivot2
		 *
		 * Pointer k is the first index of ?-part.
		 */
	Outer:
		for k := less - 1; ; {
			k++
			if k <= great {
			} else {
				break
			}
			var ak = a[k]
			if compare(ak, pivot1) < 0 { // Move a[k] to left part
				a[k] = a[less]
				/*
				 * Here and below we use "a[i] = b; i++;" instead
				 * of "a[i++] = b;" due to performance issue.
				 */
				a[less] = ak
				less++
			} else if compare(ak, pivot2) > 0 { // Move a[k] to right part
				for compare(a[great], pivot2) > 0 {
					great--
					if great+1 == k {
						break Outer
					}
				}
				if compare(a[great], pivot1) < 0 { // a[great] <= pivot2
					a[k] = a[less]
					a[less] = a[great]
					less++
				} else { // pivot1 <= a[great] <= pivot2
					a[k] = a[great]
				}
				/*
				 * Here and below we use "a[i] = b; i--;" instead
				 * of "a[i--] = b;" due to performance issue.
				 */
				a[great] = ak
				great--
			}
		}

		// Swap pivots into their final positions
		a[left] = a[less-1]
		a[less-1] = pivot1
		a[right] = a[great+1]
		a[great+1] = pivot2

		// Sort left and right parts recursively, excluding known pivots
		sortInternal(compare, a, left, less-2, leftmost)
		sortInternal(compare, a, great+2, right, false)

		/*
		 * If center part is too large (comprises > 4/7 of the array),
		 * swap internal pivot values to ends.
		 */
		if less < e1 && e5 < great {
			/*
			 * Skip elements, which are equal to pivot values.
			 */
			for compare(a[less], pivot1) == 0 {
				less++
			}

			for compare(a[great], pivot2) == 0 {
				great--
			}

			/*
			 * Partitioning:
			 *
			 *   left part         center part                  right part
			 * +----------------------------------------------------------+
			 * | == pivot1 |  pivot1 < && < pivot2  |    ?    | == pivot2 |
			 * +----------------------------------------------------------+
			 *              ^                        ^       ^
			 *              |                        |       |
			 *             less                      k     great
			 *
			 * Invariants:
			 *
			 *              all in (*,  less) == pivot1
			 *     pivot1 < all in [less,  k)  < pivot2
			 *              all in (great, *) == pivot2
			 *
			 * Pointer k is the first index of ?-part.
			 */
		outer2:
			for k := less - 1; ; {
				k++
				if k <= great {
				} else {
					break
				}
				var ak = a[k]
				if compare(ak, pivot1) == 0 { // Move a[k] to left part
					a[k] = a[less]
					a[less] = ak
					less++
				} else if compare(ak, pivot2) == 0 { // Move a[k] to right part
					for compare(a[great], pivot2) == 0 {
						great--
						if great+1 == k {
							break outer2
						}
					}
					if compare(a[great], pivot1) == 0 { // a[great] < pivot2
						a[k] = a[less]
						/*
						 * Even though a[great] equals to pivot1, the
						 * assignment a[less] = pivot1 may be incorrect,
						 * if a[great] and pivot1 are floating-point zeros
						 * of different signs. Therefore in float and
						 * double sorting methods we have to use more
						 * accurate assignment a[less] = a[great].
						 */
						a[less] = pivot1
						less++
					} else { // pivot1 < a[great] < pivot2
						a[k] = a[great]
					}
					a[great] = ak
					great--
				}
			}
		}

		// Sort center part recursively
		sortInternal(compare, a, less, great, false)

	} else { // Partitioning with one pivot
		/*
		 * Use the third of the five sorted elements as pivot.
		 * This value is inexpensive approximation of the median.
		 */
		var pivot = a[e3]

		/*
		 * Partitioning degenerates to the traditional 3-way
		 * (or "Dutch National Flag") schema:
		 *
		 *   left part    center part              right part
		 * +-------------------------------------------------+
		 * |  < pivot  |   == pivot   |     ?    |  > pivot  |
		 * +-------------------------------------------------+
		 *              ^              ^        ^
		 *              |              |        |
		 *             less            k      great
		 *
		 * Invariants:
		 *
		 *   all in (left, less)   < pivot
		 *   all in [less, k)     == pivot
		 *   all in (great, right) > pivot
		 *
		 * Pointer k is the first index of ?-part.
		 */
		for k := less; k <= great; k++ {
			if compare(a[k], pivot) == 0 {
				continue
			}
			var ak = a[k]
			if compare(ak, pivot) < 0 { // Move a[k] to left part
				a[k] = a[less]
				a[less] = ak
				less++
			} else { // a[k] > pivot - Move a[k] to right part
				for compare(a[great], pivot) > 0 {
					great--
				}
				if compare(a[great], pivot) < 0 { // a[great] <= pivot
					a[k] = a[less]
					a[less] = a[great]
					less++
				} else { // a[great] == pivot
					/*
					 * Even though a[great] equals to pivot, the
					 * assignment a[k] = pivot may be incorrect,
					 * if a[great] and pivot are floating-point
					 * zeros of different signs. Therefore in float
					 * and double sorting methods we have to use
					 * more accurate assignment a[k] = a[great].
					 */
					a[k] = pivot
				}
				a[great] = ak
				great--
			}
		}

		/*
		 * Sort left and right parts recursively.
		 * All elements from center part are equal
		 * and, therefore, already sorted.
		 */
		sortInternal(compare, a, left, less-1, leftmost)
		sortInternal(compare, a, great+1, right, false)
	}
}

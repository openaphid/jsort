package sort_slice_dps

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

type SliceInterface interface {
	make(n int) SliceInterface
	copy(src SliceInterface)
	get(i int) interface{}
	set(i int, v interface{})
	swap(i, j int)
	len() int
	slice(i int) SliceInterface
	slice2(i, j int) SliceInterface
	cmp(compare CompareFunc, i, j int) int
}

type CompareFunc = func(o1, o2 interface{}) int

func Sort(a SliceInterface, compare CompareFunc) {
	sort(compare, a, 0, a.len()-1, nil, 0, 0)
}

func IsSorted(a SliceInterface, compare CompareFunc) bool {
	for i := a.len() - 1; i > 0; i-- {
		if a.cmp(compare, i, i-1) < 0 {
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
func sort(compare CompareFunc, a SliceInterface, left int, right int, work SliceInterface, workBase int, workLen int) {
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
		for k < right && a.cmp(compare, k, k+1) == 0 {
			k++
		}
		if k == right {
			break
		} // Sequence finishes with equal items
		if a.cmp(compare, k, k+1) < 0 { // ascending
			for {
				k++
				if k <= right && a.cmp(compare, k-1, k) <= 0 {
				} else {
					break
				}
			}

		} else if a.cmp(compare, k, k+1) > 0 { // descending
			for {
				k++
				if k <= right && a.cmp(compare, k-1, k) >= 0 {
				} else {
					break
				}
			}
			// Transform into an ascending sequence
			lo := run[count] - 1
			for hi := k; lo+1 < hi-1; {
				lo++
				hi--
				a.swap(lo, hi)
			}
		}

		// Merge a transformed descending sequence followed by an
		// ascending sequence
		if run[count] > left && a.cmp(compare, run[count], run[count]-1) >= 0 {
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
	var b SliceInterface    // temp array; alternates with a
	var ao, bo int          // array offsets from 'left'
	var blen = right - left // space needed for b
	if work == nil || workLen < blen || workBase+blen > work.len() {
		work = a.make(blen)
		workBase = 0
	}
	if odd == 0 {
		work.slice(workBase).copy(a.slice2(left, left+blen))
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
				if q >= hi || p < mi && a.cmp(compare, p+ao, q+ao) <= 0 {
					//b[i + bo] = a[p++ + ao];
					b.set(i+bo, a.get(p+ao))
					p++
				} else {
					//b[i + bo] = a[q++ + ao];
					b.set(i+bo, a.get(q+ao))
					q++
				}
			}
			last++
			run[last] = hi
		}
		if (count & 1) != 0 {
			i := right
			lo := run[count-1]
			for ; i-1 >= lo; b.set(i+bo, a.get(i+ao)) {
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
func sortInternal(compare CompareFunc, a SliceInterface, left int, right int, leftmost bool) {
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
				var ai = a.get(i + 1)
				for compare(ai, a.get(j)) < 0 {
					a.set(j+1, a.get(j))
					j--
					if j+1 == left {
						break
					}
				}
				a.set(j+1, ai)

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
				if a.cmp(compare, left, left-1) >= 0 {
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

				var a1 = a.get(k)
				var a2 = a.get(left)

				if compare(a1, a2) < 0 {
					a2 = a1
					a1 = a.get(left)
				}

				for {
					k--
					if compare(a1, a.get(k)) < 0 {
					} else {
						break
					}
					a.set(k+2, a.get(k))
				}
				k++
				a.set(k+1, a1)

				for {
					k--
					if compare(a2, a.get(k)) < 0 {
					} else {
						break
					}
					a.set(k+1, a.get(k))
				}
				a.set(k+1, a2)
				left++
				k = left
			}
			var last = a.get(right)

			for {
				right--
				if compare(last, a.get(right)) < 0 {
				} else {
					break
				}
				a.set(right+1, a.get(right))
			}
			a.set(right+1, last)
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
	var e3 = (left + right) >> 1 // The midpoint TODO check
	var e2 = e3 - seventh
	var e1 = e2 - seventh
	var e4 = e3 + seventh
	var e5 = e4 + seventh

	// Sort these elements using insertion sort
	if a.cmp(compare, e2, e1) < 0 {
		a.swap(e2, e1)
	}

	if a.cmp(compare, e3, e2) < 0 {
		var t = a.get(e3)
		a.swap(e3, e2)
		if compare(t, a.get(e1)) < 0 {
			a.set(e2, a.get(e1))
			a.set(e1, t)
		}
	}
	if a.cmp(compare, e4, e3) < 0 {
		var t = a.get(e4)
		a.swap(e4, e3)
		if compare(t, a.get(e2)) < 0 {
			a.set(e3, a.get(e2))
			a.set(e2, t)
			if compare(t, a.get(e1)) < 0 {
				a.set(e2, a.get(e1))
				a.set(e1, t)
			}
		}
	}
	if a.cmp(compare, e5, e4) < 0 {
		var t = a.get(e5)
		a.swap(e5, e4)
		if compare(t, a.get(e3)) < 0 {
			a.set(e4, a.get(e3))
			a.set(e3, t)
			if compare(t, a.get(e2)) < 0 {
				a.set(e3, a.get(e2))
				a.set(e2, t)
				if compare(t, a.get(e1)) < 0 {
					a.set(e2, a.get(e1))
					a.set(e1, t)
				}
			}
		}
	}

	// Pointers
	var less = left   // The index of the first element of center part
	var great = right // The index before the first element of right part

	if a.cmp(compare, e1, e2) != 0 && a.cmp(compare, e2, e3) != 0 && a.cmp(compare, e3, e4) != 0 && a.cmp(compare, e4, e5) != 0 {
		/*
		 * Use the second and fourth of the five sorted elements as pivots.
		 * These values are inexpensive approximations of the first and
		 * second terciles of the array. Note that pivot1 <= pivot2.
		 */
		var pivot1 = a.get(e2)
		var pivot2 = a.get(e4)

		/*
		 * The first and the last elements to be sorted are moved to the
		 * locations formerly occupied by the pivots. When partitioning
		 * is complete, the pivots are swapped back into their final
		 * positions, and excluded from subsequent sorting.
		 */
		a.set(e2, a.get(left))
		a.set(e4, a.get(right))

		/*
		 * Skip elements, which are less or greater than pivot values.
		 */
		for {
			less++
			if compare(a.get(less), pivot1) < 0 {
			} else {
				break
			}
		}
		for {
			great--
			if compare(a.get(great), pivot2) > 0 {
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
			var ak = a.get(k)
			if compare(ak, pivot1) < 0 { // Move a[k] to left part
				a.set(k, a.get(less))
				/*
				 * Here and below we use "a[i] = b; i++;" instead
				 * of "a[i++] = b;" due to performance issue.
				 */
				a.set(less, ak)
				less++
			} else if compare(ak, pivot2) > 0 { // Move a[k] to right part
				for compare(a.get(great), pivot2) > 0 {
					great--
					if great+1 == k {
						break Outer
					}
				}
				if compare(a.get(great), pivot1) < 0 { // a[great] <= pivot2
					a.set(k, a.get(less))
					a.set(less, a.get(great))
					less++
				} else { // pivot1 <= a[great] <= pivot2
					a.set(k, a.get(great))
				}
				/*
				 * Here and below we use "a[i] = b; i--;" instead
				 * of "a[i--] = b;" due to performance issue.
				 */
				a.set(great, ak)
				great--
			}
		}

		// Swap pivots into their final positions
		a.set(left, a.get(less-1))
		a.set(less-1, pivot1)
		a.set(right, a.get(great+1))
		a.set(great+1, pivot2)

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
			for compare(a.get(less), pivot1) == 0 {
				less++
			}

			for compare(a.get(great), pivot2) == 0 {
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
				var ak = a.get(k)
				if ak == pivot1 { // Move a[k] to left part
					a.set(k, a.get(less))
					a.set(less, ak)
					less++
				} else if ak == pivot2 { // Move a[k] to right part
					for compare(a.get(great), pivot2) == 0 {
						great--
						if great+1 == k {
							break outer2
						}
					}
					if compare(a.get(great), pivot1) == 0 { // a[great] < pivot2
						a.set(k, a.get(less))
						/*
						 * Even though a[great] equals to pivot1, the
						 * assignment a[less] = pivot1 may be incorrect,
						 * if a[great] and pivot1 are floating-point zeros
						 * of different signs. Therefore in float and
						 * double sorting methods we have to use more
						 * accurate assignment a[less] = a[great].
						 */
						a.set(less, pivot1)
						less++
					} else { // pivot1 < a[great] < pivot2
						a.set(k, a.get(great))
					}
					a.set(great, ak)
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
		var pivot = a.get(e3)

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
			if compare(a.get(k), pivot) == 0 {
				continue
			}
			var ak = a.get(k)
			if compare(ak, pivot) < 0 { // Move a[k] to left part
				a.set(k, a.get(less))
				a.set(less, ak)
				less++
			} else { // a[k] > pivot - Move a[k] to right part
				for compare(a.get(great), pivot) > 0 {
					great--
				}
				if compare(a.get(great), pivot) < 0 { // a[great] <= pivot
					a.set(k, a.get(less))
					a.set(less, a.get(great))
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
					a.set(k, pivot)
				}
				a.set(great, ak)
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

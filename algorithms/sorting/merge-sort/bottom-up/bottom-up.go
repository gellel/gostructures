// Package bottomup provides MergeSort
// using the bottom-up merge sort scheme.
// This implementation subsets the original
// sequence by base two per cycle. Per each division
// these subsets are then merged sequentially to create
// the sorted sequence until the whole Slice
// or Array is sorted.
package bottomup

// MergeSort algorithm using bottom-up merge sort scheme.
func MergeSort(a []int) []int {
	return sort(a, clone(a), len(a))
}

// Clone duplicates Slice or Array A
// and spreads A's contents into Slice B.
// Effectively builtins.copy but simplified
// to a single returnable line.
func clone(a []int) []int {
	return append(make([]int, 0), a...)
}

// Merge performs the core bottom up merge sort method.
// Merge takes sequence A and B and performs
// re-ordering operations.
func merge(a []int, b []int, floor int, ceiling int, n int) []int {
	// set base index for shifting
	// lower bounds comparision.
	i := floor
	// set base index for shifting
	// upper bounds.
	j := ceiling
	// iterate across range for
	// current partition size.
	// this is determine by the size
	// of (i + (2 * width)).
	for k := floor; k < n; k++ {
		// confirm whether lower bounds
		// has available positions to check in the
		// subset sequence. additionally confirm
		// that upper bounds has available positions.
		// j is required to check whether it matches
		// or exceeds the size of the sequence.
		// then compare whether the duplicate sequence
		// at the current indexes contains an element
		// requiring re-positioning.
		if (i < ceiling) && ((j >= n) || (a[i] <= a[j])) {
			// set the duplicate sequence to contain
			// the upper bounds value contained in
			// the original sequence.
			b[k] = a[i]
			// increment lower bounds reference.
			i = i + 1
		} else {
			// set the duplicate sequence to contain
			// the lower bounds value contained in
			// the original sequence.
			b[k] = a[j]
			// increment upper bounds reference.
			j = j + 1
		}
	}
	// return working-copy Array.
	return b
}

// Sort partitions original sequence
// into two subsets width adjacent values;
// each set containing one element each.
// These values are then sorted and merged
// before a square of the original partition
// index is performed. For base case, this
// scales the subsetting and merging index from
// one to two; generating two subsets of
// adjacent values with two elements each.
// This operation occurs N times, scaling
// by base 2 (i.e [1, 2, 4, 8 ... N]) before
// all subsets are ordered and merged.
func sort(a []int, b []int, n int) []int {
	// generate stable subsets of of base
	// two. perform sequencing until the
	// sequence size cannot be scaled.
	for width := 1; width < n; width = (2 * width) {
		// iterate across current partition
		// setting i to be current partition
		// relative to the number of adjacent
		// elements required to be sorted.
		// it's important to scale the range
		// before incrementing the index or
		// the sequence will overflow.
		for i := 0; i < n; i = (i + (2 * width)) {
			// merge sorted vales from temporary
			// sequence B into sequence A. uses the
			// shifted range (i.e [0,1,2])
			merge(a, b, i, (i + width), (i + (2 * width)))
		}
		// clone entries from B to A for next iteration
		// as current B sequence contains previous partitions
		// from last iteration cycle. if this step
		// is omitted A remains incorrectly ordered.
		a = clone(b)
	}
	// sequence is sorted.
	return a
}

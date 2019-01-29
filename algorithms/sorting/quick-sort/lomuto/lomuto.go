// Package lomuto provides QuickSort 
// using the Lomuto partitioning scheme.
// This implementation chooses a pivot that is the
// last element in a Slice or Array. 
package lomuto


// QuickSort algorithm using Lomuto partition scheme.
func QuickSort(a []int) []int {
    return quicksort(a, 0, len(a)-1)
}

// QuickSort algorithm using Lomuto partition scheme.
// Uses floor: 0 and ceiling: n-1 to divide
// and organise the unsorted Slice or Array.
func quicksort(a []int, floor int, ceiling int) []int {
    // checks whether the recursive floor position
    // has been shifted n times to match the ceiling.
    // a sorted Slice or Array would have a floor equal to
    // ceiling, or simply N.
	if floor < ceiling {
        // perform core Quicksort algorithm.
		p := partition(a, floor, ceiling)
        // organise unsorted Slice or Array
        // that is on the left side of the
        // generated pivot. 
        // shifts ceiling lower.
        quicksort(a, floor, (p - 1))
        // organise unsorted Slice or Array
        // that is on the right side of the
        // generated pivot. 
        // shifts floor higher.
		quicksort(a, (p + 1), ceiling)
    }
    // Slice or Array is organised.
	return a
}

// Partition using Lomuto partition scheme.
// Uses ceiling sum as the partition pivot.
func partition(a []int, floor int, ceiling int) int {
    // define pivot to operate as condition
    // for assigning items that are misplaced.
    // uses sum of pivot to be ordering field.
    // this allows operation to naturally
    // organise by ascending order when the.
    // value of element J is less than P. 
    p := a[ceiling]
    
    // define base left swap position.
    // required for shifting a reliable
    // index J indexes right. usage
    // of floor - 1 creates a convention 
    // where incrimenting i before swapping
    // is required. 
    i := (floor - 1)

    // scans across Slice or Array J number of
    // times relative to the distance of P from J.
    // can be expressed as if floor + 1 is
    // not ceiling - 1.
    for j := floor; j <= (ceiling - 1); j++ {
        // confirms whether the current
        // element at position J requires
        // a swap. element at J should be a lower
        // sum than P.
        if a[j] <= p {
            // update next swap position.
            i = i + 1
            // perform the swap.
            swap(a, i, j)
        }
    }
    // swap the element value found
    // that exceeds the current ceiling.    
    swap(a, i + 1, ceiling)
    // previous position of highest
    // found element in Slice or Array becomes
    // new recursive pivot. 
    return i + 1
}

// Swap reassigns elements within
// a Slice or Array.
func swap(a []int, i int, j int) {
    a[i], a[j] = a[j], a[i]
}

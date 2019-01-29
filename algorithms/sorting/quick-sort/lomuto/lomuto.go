// Package lomuto exports Quicksort function using the partition scheme designed by Nico Lomuto.
// Quicksort using this implementation offers average case
// Omega(n log(n)) but degrades to O(n2) for a sorted slice.
package lomuto

// QuickSort algorithm using Lomuto partition scheme.
func QuickSort(a []int) []int {
    return quicksort(a, 0, len(a)-1)
}

// QuickSort algorithm using Nico Lomuto partition scheme.
// Uses floor as N(0) and ceiling N(-1) to divide
// and organise the unsorted slice.
func quicksort(a []int, floor int, ceiling int) []int {
	if floor < ceiling {
		p := partition(a, floor, ceiling)
        quicksort(a, floor, (p - 1))
		quicksort(a, (p + 1), ceiling)
    }
	return a
}

// Partition using Lomuto partition scheme.
func partition(a []int, floor int, ceiling int) int {
    /* define pivot to operate as condition
     * for assigning items that are misplaced.
     * uses sum of pivot to be ordering field.
     * this allows operation to naturally
     * organise by ascending order when the.
     * value of element j is less than p. 
     */
    p := a[ceiling]
    
    /* define base left swap position. 
     * required for shifting a reliable
     * index j indexes right. usage of (floor - 1) 
     * creates a convention where incrimenting
     * i before swapping is required. 
     */
    i := (floor - 1)

    /* iterate across sequence j number of times, relative
     * to the distance of p from j. can be expressed
     * as perform operation when (floor + 1) is not (ceiling - 1).
     */
    for j := floor; j <= (ceiling - 1); j++ {
        /* confirms whether the current element 
         * at position j requires a swap. element
         * at j should be a lower sum than p. 
         */
        if a[j] <= p {
            /* increment i to insure that the current
             * sequence does not underflow and 
             * throw an out of provided range exception.
             */
            i = (i + 1)
            /* swap the current indexes. */
            swap(a, i, j)
        }
    }
    /* swap the element value found that exceeds the current ceiling. */
    swap(a, (i + 1), ceiling)
    /* (i + 1) becomes the new partition index. */
    return i + 1
}

// Swap reassigns elements within a slice.
func swap(a []int, i int, j int) {
    a[i], a[j] = a[j], a[i]
}

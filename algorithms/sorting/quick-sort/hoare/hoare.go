// Package hoare exports Quicksort function using the partition
// scheme designed by British computer scientist Tony Hoare.
// Quicksort using this implementation offers average case
// Omega(n log(n)) but degrades to O(n2) for a sorted slice.
package hoare

// QuickSort algorithm using Hoare partition scheme.
func QuickSort(a []int) []int {
    return quicksort(a, 0, len(a)-1)
}

// QuickSort algorithm using Tony Hoare partition scheme.
// Uses floor as N(0) and ceiling N(-1) to divide
// and organise the unsorted slice.
func quicksort(a []int, floor int, ceiling int) []int {
    if floor < ceiling {
        p := partition(a, floor, ceiling)
        quicksort(a, floor, p)
        quicksort(a, p + 1, ceiling)
    }
    return a
}

// Partition using Tony Hoare partition scheme.
func partition(a []int, floor int, ceiling int) int {
    /* define pivot to operate as condition
     * for assigning items that are misplaced.
     * uses sum of pivot to be ordering field. 
     */
    p := a[floor]

    /* define left swap index position. this is required
     * for determining whether the value of element at a(i)
     * belongs on the right side of the created pivot (p).
     * setting floor as (floor-1) creates a convention where
     * incrementing i before swapping index positions is
     * required to prevent an underflow.
     */
    i := (floor - 1)

    /* define right swap index position. this is required
     * for determining whether the value of element at a(j)
     * belongs on the left side of the created pivot (p).
     * setting ceiling as (ceiling + 1) creates a convention where
     * incrementing i before swapping index positions is
     * required to prevent an underflow.
     */
    j := (ceiling + 1)

    /* continiously perform shifting operation
     * until all elements in the slice have been
     * sequentially sorted.
     */
    for {
        /* while the element at position a(j) does not
         * exceed the pivot, continue iterating through
         * the sequence up to the partition limit.
         */
        for {
            j = (j - 1)
            if a[j] <= p {
                break
            }
        }
        /* while the element at position a(i) is smaller
         * than the current the pivot, continue iterating through
         * the sequence up to the partition limit.
         */
        for {
            i = (i + 1)
            if a[i] >= p {
                break
            }
        }
        /* after an inversion of i and j has occurred the 
         * sequence has been exhausted of operations. 
         */
        if i >= j {
            break
        }
        /* swap the current elements at a(i) and a(j) */
        swap(a, i, j)
    }
    /* j becomes the new partition index. */
    return j
}

// Swap reassigns elements within a slice.
func swap(a []int, i int, j int) {
    a[i], a[j] = a[j], a[i]
}

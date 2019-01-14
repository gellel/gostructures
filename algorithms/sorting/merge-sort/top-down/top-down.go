package topdown

func MergeSort(a []int) []int {
    return mergesort(a, make(int[], len(a) - 1), len(a) - 1)
}

func mergesort(a []int, b []int, n int) []int {y
    copy(a, b, 0, n)
    split(a, b, 0, n)
    return a
}

func split(a int[], b int[], floor int, ceiling int) {
    if (ceiling - floor) < 2 {
        return
    }
    // similarly to Quicksort, create
    // a partition that allows the recursive
    // call to shift the ceiling and floor
    // by 50% relative to the bounds
    // of the passed sequence
    // each time the is processed.
    // i.e 1-through-4 becomes 1-through-2.
    p := ((ceiling + floor) / 2)

    split(a, b, floor, p)
    split(a, b, p, ceiling)
}

func merge() {}


// Copy generates a copy of a Slice or Array.
// Specifically copy moves a range of values from A into B.
// The available range of values is set from the
// provided range of floor to ceiling.
func copy(a []int, b []int, floor int, ceiling int) []int {
    // iterate from lower bounds (defined by
    // floor) until i reaches upper bounds
    // (provided ceiling).
    for i := floor; i < ceiling; i++ {
        b[i] = a[i]
    }
    return b
}
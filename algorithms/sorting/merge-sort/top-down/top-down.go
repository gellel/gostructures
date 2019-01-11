package topdown

func MergeSort() []int {

}

func mergesort(a []int, b []int, n int) []int {

}

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
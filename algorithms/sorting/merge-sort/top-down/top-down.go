// Package topdown provides Mergesort 
// using the top-down merge sort scheme.
// This implementation recursively subdivides
// N number of sequences until subset reaches 
// maximum a length of 1. These subsets
// are then merged sequentially to create
// the sorted sequence.
package topdown

// Mergesort algorithm using top-down merge sort scheme.
func Mergesort(a []int) []int {
    return sort(a, clone(a), 0, len(a))
}

// Clone duplicates Slice or Array A
// and spreads A's contents into Slice B.
// Effectively builtins.copy but simplified
// to a single returnable line.
func clone(a []int) []int {    
    return append(make([]int, 0), a...)
}

// Merge performs the core top down merge sort method.
// Merge takes sequence A and B and performs
// re-ordering operations. Argument A is generally
// provided as the temporary Array B in most
// implementations but has been re-ordered here
// for personal clarity. This helps identifying how
// the original Slice or Array (A) is mutated across
// function calls.
func merge(a []int, b []int, floor int, partition int, ceiling int) []int {
    //
    i := floor
    //
    j := partition
    
    for k := floor; k < ceiling; k++ {
        // confirm i has not exceeded the provided
        // upper bounds 
        if (i < partition) && ((j >= ceiling) || (b[i] <= b[j])) {
            a[k] = b[i]
            i = i + 1
        } else {
            a[k] = b[j]
            j = j + 1
        }
    }
    return a
}

// Sort subsets
func sort(a []int, b []int, floor int, ceiling int) []int {
    if !((ceiling - floor) < 2) {

        partition := ((floor + ceiling) / 2)

        sort(b, a, floor, partition)
        sort(b, a, partition, ceiling)
        merge(a, b, floor, partition, ceiling)
    }
    return a
}
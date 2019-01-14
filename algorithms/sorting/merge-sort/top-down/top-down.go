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
    // set base index for shifting
    // lower bounds comparision.
    i := floor
    // set base index for shifting
    // upper bounds.
    j := partition
    
    // iterate across sequence using shifted floor 
    // and ceiling from recursive sorted calls.
    // floor being either floor from
    // subdivision left (i.e, 0 -> n/2) or
    // partition for subdivision right (n/2 -> n)
    for k := floor; k < ceiling; k++ {
        // confirm whether lower bounds
        // has available positions to check in the
        // subset sequence. additionally confirm
        // that upper bounds has available positions.
        // j is required to check whether it matches
        // or exceeds the ceiling as upper bounds of subset
        // must be the highest ordered sum of the current
        // subset sequence. i.e [j=1, ceiling=2, then, j=2, ceiling=2]
        // then compare whether the duplicate sequence
        // at the current indexes contains an element
        // requiring re-positioning.
        if (i < partition) && ((j >= ceiling) || (b[i] <= b[j])) {
            // set the original sequence to contain
            // the upper bounds value contained in
            // the duplicate sequence. this
            // should substitute at the matching
            // value of a[k]. 
            // i.e a[1, 1], b[2, 1] becomes a[1, a[k]->2<-b[i]].
            a[k] = b[i]
            // increment lower bounds reference.
            i = i + 1
        } else {
            // set the original sequence to contain
            // the lower bounds value from the
            // duplicate sequence. generates a sorted
            // array. i.e a[2, 1] becomes a[1, 1].
            a[k] = b[j]
            // increment upper bounds reference.
            j = j + 1
        }
    }
    return a
}

// Sort subsets floor and ceiling into
// n/2 sets of original sequence size.
// Repeats operations for logN iterations until
// the sum of logN can no longer be evenly halved.
func sort(a []int, b []int, floor int, ceiling int) []int {
    // confirm whether current range of
    // subset can be further divided.
    // general implementation does not negate
    // condition.
    if !((ceiling - floor) < 2) {
        // set partition index for left
        // and right sides of current subset.
        partition := ((floor + ceiling) / 2)
        // recurively generate new partitions
        // and subsets for left and right
        // sequence ranges. in the recursion call case,
        // B takes the position of A as the
        // "original" sequence as merge updates
        // A from sorted values from B. if passed in
        // in the A B case, sorting does not work.
        sort(b, a, floor, partition)
        sort(b, a, partition, ceiling)
        merge(a, b, floor, partition, ceiling)
    }
    return a
}
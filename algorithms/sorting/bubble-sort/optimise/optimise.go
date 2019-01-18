// Package optimise exports the optimised implementation of the Bubble Sort algorithm. 
// The base Bubble Sort iterates across both dimensions of the provided
// Slice or Array, continously comparing for O(n2) operations until the argument
// sequence has completed all iteration cycles across i and j. Unlike it's
// unoptimised counterpart, the optimised algorithm can avoid looking at the last n − 1 
// items when running for the n-th time.
package normal

// BubbleSort algorithm using base sort scheme.
func BubbleSort(a []int) []int {
    return bubble(a, len(a)-1)
}

// Bubble performs the core Bubble Sort sorting operation. Algorithm begins by
// iterating across the n - 1 size of the argument Slice or Array. The iteration
// across i for each cycle generates a n - i loop continuously comparing
// elements in the sequeuence at adjacent sides and swapping the values
// should the pivot (j) exceed the adjacent pivot (j + 1). An example of
// the positions for a N of 4 would resemble [0, 3] for 3 iterations.
func bubble(a []int, n int) []int {
    var swapped bool
    for i := 0; i < n; i++ {
        swapped = false
        for j := 0; j < (n-i); j++ {
            if a[j] > a[j + 1] {
                swap(a, j, j + 1)
                swapped = true
            }
        }
        if swapped == false {
            break
        }
    }
    return a
}

func swap(a []int, i int, j int) {
    a[i], a[j] = a[j], a[i]
}
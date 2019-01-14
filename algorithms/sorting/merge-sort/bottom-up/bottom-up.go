package bottomup

// Mergesort algorithm using bottom-up merge sort scheme.
func Mergesort(a []int) []int {
	return sort(a, clone(a), len(a))
}

// Clone duplicates Slice or Array A
// and spreads A's contents into Slice B.
// Effectively builtins.copy but simplified
// to a single returnable line.
func clone(a []int) []int {    
	return append(make([]int, 0), a...)
}

func merge(a []int, b []int, left int, right int, end int) []int {
    i := left
    j := right
    for k := left; k < end; k++ {
        if ((i < right) && ((j >= end) || (a[i] <= a[j]))) {
            b[k] = a[i]
            i = i + 1
        } else {
            b[k] = a[j]
            j = j + 1
        }
    }
    return b
}

func min(a int, b int) int {
    if a < b {
        return a
    }
    return b
}

func sort(a []int, b []int, n int) []int {

    for width := 1; width < n; width = (2 * width) {
        for i := 0; i < n; i = ((i + 2) * width) {
            merge(a, b, i, min((i + width), n), min(((i + 2) * width), n))
        }
        a = clone(b)
    }
    return a
}
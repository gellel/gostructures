package bottomup

// Mergesort algorithm using bottom-up merge sort scheme.
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

func merge(a []int, b []int, floor int, ceiling int) []int {
    i := floor
    j := ceiling
    for k := floor; k < ceiling; k++ {
        if ((i < ceiling) && ((j >= ceiling) || (b[i] <= b[j]))) {
            a[k] = b[i]
            i = i + 1
        } else {
            a[k] = b[j]
            j = j + 1
        }
    }
    return a
}

func sort(a []int, b []int, n int) {
}
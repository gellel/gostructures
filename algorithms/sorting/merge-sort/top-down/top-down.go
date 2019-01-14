package topdown

func Mergesort(a []int) []int {
    return sort(a, clone(a), 0, len(a))
}

func clone(a []int) []int {    
    return append(make([]int, 0), a...)
}

func merge(a []int, b []int, floor int, partition int, ceiling int) []int {
    i := floor
    j := partition
    
    for k := floor; k < ceiling; k++ {
        if (i < partition) && ((j >= ceiling) || (a[i] <= a[j])) {
            b[k] = a[i]
            i = i + 1
        } else {
            b[k] = a[j]
            j = j + 1
        }
    }
    return b
}

func sort(a []int, b []int, floor int, ceiling int) []int {
    if !((ceiling - floor) < 2) {

        partition := ((floor + ceiling) / 2)

        sort(b, a, floor, partition)
        sort(b, a, partition, ceiling)
        merge(b, a, floor, partition, ceiling)
    }
    return a
}
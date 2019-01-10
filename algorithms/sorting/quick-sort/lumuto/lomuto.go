package lomuto


// Quicksort algorithm using Lomuto partition scheme.
// Uses floor: 0 and ceiling: n-1 to divide
// and organise the unsorted Slice or Array.
func Quicksort(a []int, floor int, ceiling int) []int {

	if floor < ceiling {
		partition := Partition(a, floor, ceiling)

		Quicksort(a, floor, (partition - 1))
		Quicksort(a, (partition + 1), ceiling)
	}
	return a
}

func Partition(a []int, floor int, ceiling int) int {
    
    p := a[ceiling]
    i := (floor - 1)

    for j := floor; j <= (ceiling - 1); j++ {
        if a[j] <= p {
            i = i + 1
            Swap(a, i, j)
        }
    }
    Swap(a, i + 1, ceiling)

    return i + 1
}

func Swap(a []int, i int, j int) {

    a[i], a[j] = a[j], a[i]
}
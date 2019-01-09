package lomuto

func Quicksort(unsorted *[]int) *[]int {
	return quicksort(unsorted, 0, len(*unsorted)-1)
}

func quicksort(A *[]int, low int, high int) *[]int {
	if low < high {
		p := partition(A, low, high)

		quicksort(A, low, p-1)
		quicksort(A, p+1, high)
	}
	return A
}

func partition(A *[]int, low int, high int) int {
	pivot := (*A)[high]
	i := low
	for j := low; j < high-1; j++ {
		if (*A)[j] < pivot {
			if i != j {

			}
		}
	}
	return i
}

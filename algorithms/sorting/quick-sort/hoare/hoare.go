package hoare

func Quicksort(a []int) []int {
	return quicksort(a, 0, len(a)-1)
}

func quicksort(a []int, floor int, ceiling int) []int {
	if floor < ceiling {
		p := partition(a, floor, ceiling)

		quicksort(a, floor, p)
		quicksort(a, p+1, ceiling)
	}
	return a
}

func partition(a []int, floor int, ceiling int) int {
	p := a[floor]
	i := floor - 1
	j := ceiling + 1

	for {
		for {
			j = j - 1
			if a[j] <= p {
				break
			}
		}
		for {
			i = i + 1
			if a[i] >= p {
				break
			}
		}
		if i < j {
			swap(a, i, j)
		} else {
			break
		}
	}
	return j
}

func swap(a []int, i int, j int) {
	a[i], a[j] = a[j], a[i]
}

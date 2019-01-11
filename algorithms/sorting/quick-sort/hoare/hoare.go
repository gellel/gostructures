// Package hoarse provides Quicksort 
// using the original Hoare partitioning scheme.
// This implementation chooses a pivot that is the
// last element in a Slice or Array. 
package hoare

// Quicksort algorithm using Hoare partition scheme.
func Quicksort(a []int) []int {
	return quicksort(a, 0, len(a)-1)
}

// Quicksort algorithm using Hoare partition scheme.
// Uses floor: 0 and ceiling: n-1 to divide
// and organise the unsorted Slice or Array.
func quicksort(a []int, floor int, ceiling int) []int {
	// checks whether the recursive floor position
    // has been shifted n times to match the ceiling.
    // a sorted Slice or Array would have a floor equal to
    // ceiling, or simply N.
	if floor < ceiling {
        // perform core Quicksort algorithm.
		p := partition(a, floor, ceiling)
        // organise unsorted Slice or Array
        // that is on the left side of the
        // generated pivot. 
        // use unmodified pivot as
        // the ceiling.
        quicksort(a, floor, p)
        // organise unsorted Slice or Array
        // that is on the right side of the
        // generated pivot. 
        // shifts floor higher.
		quicksort(a, p + 1, ceiling)
    }
    // Slice or Array is organised.
	return a
}

// Partition using Hoare partition scheme.
// Uses floor sum as the partition pivot.
func partition(a []int, floor int, ceiling int) int {
    // define pivot to operate as condition
    // for assigning items that are misplaced.
    // uses sum of pivot to be ordering field.
    p := a[floor]

    // define base left swap position.
    // required for determining whether
    // the value of element at a[i] belongs on
    // the right side of the pivot. usage
    // of floor - 1 creates a convention 
    // where incrementing i before swapping
    // is required. 
    i := floor - 1
    
    // define base right swap position.
    // required for determining whether
    // the value of element at a[j] belongs on
    // the left side of the pivot. usage
    // of ceiling + 1 creates a convention 
    // where decrementing j before swapping
    // is required. 
	j := ceiling + 1

    // set infinite loop
	for {
        // continue scanning across
        // a[n] until iteration j finds
        // value of a[j] is smaller or
        // equal to the defined pivot. 
		for {
			j = j - 1
			if a[j] <= p {
				break
			}
        }
        // continue scanning across
        // a[n] until iteration i finds
        // value of a[i] is larger or
        // equal to the defined pivot.
		for {
			i = i + 1
			if a[i] >= p {
				break
			}
		}
		if i >= j {
			break
        }
        // once i and j have reached an
        // inversion swap the values of
        // a[i] and a[j] as a[i] will
        // contain a value that is larger
        // than a[j].
		swap(a, i, j)
    }
    // j index becomes
    // new recursive pivot.
	return j
}

// Swap reassigns elements within
// a Slice or Array.
func swap(a []int, i int, j int) {
	a[i], a[j] = a[j], a[i]
}

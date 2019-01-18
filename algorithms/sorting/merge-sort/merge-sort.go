package mergesort

import (
	bottomup "github.com/gellel/gostructures/algorithms/sorting/merge-sort/bottom-up"
	topdown "github.com/gellel/gostructures/algorithms/sorting/merge-sort/top-down"
)

// BottomUp performs the MergeSort algorithm using the bottom-up scheme.
func BottomUp(a []int) []int {
	return bottomup.MergeSort(a)
}

// TopDown performs the MergeSort algorithm using the top-down scheme.
func TopDown(a []int) []int {
	return topdown.MergeSort(a)
}

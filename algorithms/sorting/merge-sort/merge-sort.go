package mergesort

import (
	bottomup "github.com/gellel/gostructures/algorithms/sorting/merge-sort/bottom-up"
	topdown "github.com/gellel/gostructures/algorithms/sorting/merge-sort/top-down"
)

// BottomUp merge-sort algorithm.
func BottomUp(a []int) []int {
	return bottomup.MergeSort(a)
}

// TopDown merge-sort algorithm
func TopDown(a []int) []int {
	return topdown.MergeSort(a)
}

// Package quicksort provides both
// Hoare and Lomuto QuickSort implementations.
// Faster operation is Hoare-partion-scheme.
package quicksort

import (

	"github.com/gellel/gostructures/algorithms/sorting/quick-sort/hoare"
	"github.com/gellel/gostructures/algorithms/sorting/quick-sort/lomuto"
)

// Hoare partition scheme QuickSort algorithm.
func Hoare(a []int) []int {
	return hoare.QuickSort(a)
}

// Lomuto partition scheme QuickSort algorithm.
func Lomuto(a []int) []int {
	return lomuto.QuickSort(a)
}
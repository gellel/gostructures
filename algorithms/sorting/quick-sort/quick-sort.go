// Package quicksort provides both
// Hoare and Lomuto Quicksort implementations.
// Faster operation is Hoare-partion-scheme.
package quicksort

import (

	"github.com/gellel/gostructures/algorithms/sorting/quick-sort/hoare"
	"github.com/gellel/gostructures/algorithms/sorting/quick-sort/lomuto"
)

// Hoare partition scheme Quicksort algorithm.
func Hoare(a []int) []int {
	return hoare.Quicksort(a)
}

// Lomuto partition scheme Quicksort algorithm.
func Lomuto(a []int) []int {
	return lomuto.Quicksort(a)
}
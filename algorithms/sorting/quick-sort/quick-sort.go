package quicksort

import (

	"github.com/gellel/gostructures/algorithms/sorting/quick-sort/hoare"
	"github.com/gellel/gostructures/algorithms/sorting/quick-sort/lomuto"
)

func Hoare(a []int) []int {
	return hoare.Quicksort(a)
}

func Lomuto(a []int) []int {
	return lomuto.Quicksort(a)
}
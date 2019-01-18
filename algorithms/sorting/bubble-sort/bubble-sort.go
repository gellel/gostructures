package bubblesort

import (
    "github.com/gellel/gostructures/algorithms/sorting/bubble-sort/normal"
    "github.com/gellel/gostructures/algorithms/sorting/bubble-sort/optimise"
)

// Normal performs the base BubbleSort algorithm.
func Normal(a []int) []int {
    return normal.BubbleSort(a)
}

// Optimise performs the optimised BubbleSort algorithm.
func Optimise(a []int) []int {
    return optimise.BubbleSort(a)
}
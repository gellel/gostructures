package hoare_test

import (
	"fmt"
	"testing"

	"github.com/gellel/gostructures/algorithms/sorting/quick-sort/hoare"
)

func Test(t *testing.T) {

	a := []int{1, 1, 2, 4, 56, 2, 19, 21, 6, 81, 51, 1, 3, 4, 6, 3}

	fmt.Println(hoare.Quicksort(a))
}

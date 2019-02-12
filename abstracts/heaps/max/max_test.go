package max_test

import (
	"fmt"
	"testing"

	"github.com/gellel/gostructures/abstracts/heaps/max"
)

func Test(t *testing.T) {

	heap := max.Heap{}

	heap.Push(1)

	heap.Push(3)

	heap.Push(2)

	fmt.Println(heap.Search(3), heap.Search(1), heap.Search(2), heap)
}

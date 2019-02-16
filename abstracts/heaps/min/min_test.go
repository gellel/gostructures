package min_test

import (
	"fmt"
	"testing"

	"github.com/gellel/gostructures/abstracts/heaps/min"
)

func Test(t *testing.T) {

	heap := min.Heap{}

	heap.Push(1)

	heap.Push(3)

	heap.Push(2)

	fmt.Println(heap.Search(3), heap.Search(1), heap.Search(2), heap)
}

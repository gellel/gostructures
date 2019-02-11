package max_test

import (
	"fmt"
	"testing"

	"github.com/gellel/gostructures/abstracts/heaps/max"
)

func Test(t *testing.T) {

	heap := max.Heap{}

	fmt.Println(heap.Length())

	fmt.Println(heap.Bounds(1))

	fmt.Println(heap.Violates(0))

	fmt.Println(heap.PeekLast())

	u := make([]uint, 1)

	fmt.Println(heap.Fill(u))
}

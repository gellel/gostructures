package max_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/gellel/gostructures/abstracts/heaps/max"
)

func Test(t *testing.T) {

	heap := max.Heap{}

	max := 10

	for i := 0; i != max; i++ {
		heap.Push(uint(rand.Intn(max-1) + 1))
	}

	fmt.Println("max.Heap:", heap)

	for heap.IsNotEmpty() {
		fmt.Println("value:", heap.Pop(), "n:", heap.Length())
	}
}

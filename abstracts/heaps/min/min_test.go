package min_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/gellel/gostructures/abstracts/heaps/min"
)

func Test(t *testing.T) {

	heap := min.Heap{}

	max := 10

	for i := 0; i != max; i++ {
		heap.Push(uint(rand.Intn(max-1) + 1))
	}

	fmt.Println("min.Heap:", heap)

	for heap.IsNotEmpty() {
		fmt.Println("value:", heap.Pop(), "n:", heap.Length())
	}
}

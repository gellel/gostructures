package max_test

import (
	"fmt"
	"math/rand"
	"testing"

	max "github.com/gellel/gostructures/abstracts/priority-queues/max/max-heap"
)

func Test(t *testing.T) {

	heap := max.New()

	max := 10

	for i := 0; i != max; i++ {
		heap.Push((rand.Intn(max-1) + 1), "a")
	}

	fmt.Println("max.Heap:", heap)

	for heap.IsNotEmpty() {
		fmt.Println("value:", heap.Poll(), "n:", heap.Length())
	}
}

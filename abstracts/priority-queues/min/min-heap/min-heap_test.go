package min_test

import (
	"fmt"
	"math/rand"
	"testing"

	min "github.com/gellel/gostructures/abstracts/priority-queues/min/min-heap"
)

func Test(t *testing.T) {

	heap := min.New()

	min := 10

	for i := 0; i != min; i++ {
		heap.Push((rand.Intn(min-1) + 1), "a")
	}

	fmt.Println("min.Heap:", heap)

	for heap.IsNotEmpty() {
		fmt.Println("value:", heap.Poll(), "n:", heap.Length())
	}
}

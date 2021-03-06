package min_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/gellel/gostructures/abstracts/heaps/heap/min"
)

func Test(t *testing.T) {

	heap := min.New()

	max := 10

	for i := 0; i != max; i++ {
		heap.Push(int(rand.Intn(max-1) + 1))
	}

	fmt.Println("min.Heap:", heap.Container)

	for heap.IsNotEmpty() {
		fmt.Println("value:", heap.Poll(), "n:", heap.Length())
	}
}

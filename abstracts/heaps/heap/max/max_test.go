package max_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/gellel/gostructures/abstracts/heaps/heap/max"
)

func Test(t *testing.T) {

	heap := max.New()

	max := 10

	for i := 0; i != max; i++ {
		heap.Push(int(rand.Intn(max-1) + 1))
	}

	fmt.Println("max.Heap:", heap.Container)

	for heap.IsNotEmpty() {
		fmt.Println("value:", heap.Poll(), "n:", heap.Length())
	}
}

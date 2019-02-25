package max_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/gellel/gostructures/abstracts/priority-queues/max"
)

func Test(t *testing.T) {

	priority := max.New()

	max := 10

	for i := 0; i != max; i++ {
		priority.Enqueue((rand.Intn(max-1) + 1), "a")
	}

	fmt.Println("priority.Max:", priority)

	for priority.IsNotEmpty() {
		fmt.Println("value:", priority.Dequeue(), "n:", priority.Length())
	}
}

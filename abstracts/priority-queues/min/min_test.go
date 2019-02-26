package min_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/gellel/gostructures/abstracts/priority-queues/min"
)

func Test(t *testing.T) {

	priority := min.New()

	max := 10

	for i := 0; i != max; i++ {
		priority.Enqueue((rand.Intn(max-1) + 1), "a")
	}

	fmt.Println("priority.Min:", priority)

	for priority.IsNotEmpty() {
		fmt.Println("value:", priority.Dequeue(), "n:", priority.Length())
	}
}

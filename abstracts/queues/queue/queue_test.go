package queue_test

import (
	"fmt"
	"testing"

	"github.com/gellel/gostructures/abstracts/queues/queue"
)

func Test(t *testing.T) {

	q := queue.New()

	fmt.Println(q)

	q.Enqueue(1).Enqueue(2).Enqueue(3)

	fmt.Println(q.Size())

	fmt.Println(q.Peek())

	fmt.Println(q.Poll())

	fmt.Println(q.Size())
}

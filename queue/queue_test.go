package queue_test

import (
	"fmt"
	"testing"

	"github.com/gellel/gostructures/queue"
)

func Test(t *testing.T) {

	q := queue.New(10)

	fmt.Println(q.Enqueue(1))

	fmt.Println(q.Dequeue())
}

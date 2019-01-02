package queue

import (
	"github.com/gellel/gostructures/queue/q"
)

// Queue of interfaces.
type Queue struct {
	q.Q
	queue []interface{}
}

// New instantiates a new Queue pointer.
// Size argument controls size of Queue.
func New(size int) *Queue {
	return &Queue{queue: make([]interface{}, 0)}
}

// SizeOf returns the length of the current Queue.queue.
func (queue *Queue) SizeOf() int {
	return len(queue.queue)
}

// InQueue checks whether index position falls outside of range.
func (queue *Queue) InQueue(i int) bool {
	return queue.SizeOf() >= i
}

// Enqueue appends a new item to the Queue.queue.
func (queue *Queue) Enqueue(property interface{}) (int, interface{}) {
	a := append(queue.queue, property)

	return queue.SizeOf(), a
}

package queue

import (
	"github.com/gellel/gostructures/queue/q"
)

// Queue of interfaces.
type Queue struct {
	q.Queue
	queue []interface{}
}

// New instantiates a new Queue pointer.
func New() *Queue {
	return &Queue{queue: make([]interface{}, 0)}
}

// Enqueue appends a new item to Queue.queue.
func (queue *Queue) Enqueue(property interface{}) int {
	queue.queue = append(queue.queue, property)

	return queue.SizeOf()
}

// Dequeue takes the first item from Queue.queue.
// On success, returns assigned property and remaining size.
// Returns nil if empty.
func (queue *Queue) Dequeue() (interface{}, int) {
	if queue.SizeOf() > 0 {
		return queue.queue[:1], queue.SizeOf()
	}
	return nil, -1
}

// InQueue checks whether index position falls outside of range.
func (queue *Queue) InQueue(i int) bool {
	return queue.SizeOf() >= i
}

// Peek shows the current item ad the Queue.queue index.
func (queue *Queue) Peek(i int) interface{} {
	if queue.InQueue(i) {
		return queue.queue[i]
	}
	return nil
}

// SizeOf returns the length of the current Queue.queue.
func (queue *Queue) SizeOf() int {
	return len(queue.queue)
}

// IsEmpty checks that Queue.queue does not contain entries.
func (queue *Queue) IsEmpty() bool {
	return queue.SizeOf() == 0
}

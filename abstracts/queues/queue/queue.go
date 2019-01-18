// Package queue exports a Queue. Queue struct contains a private reference to
// as Single-Linked-List to manage insertion and collection of provided data.
// Queue can contain any number of data types, but these can
// only be accessed and modified (in the context of the Queue)
// through using the Queue.Poll method.
// Package also exports a Queue pointer instantiation method (New).
package queue

import (
	"github.com/gellel/gostructures/abstracts/linked-lists/single"
)

// Queue declares the interface for a Queue.
type queue interface {
	Dequeue() *Queue
	Enqueue(value interface{}) *Queue
	IsEmpty() bool
	IsPopulated() bool
	Peek() interface{}
	Poll() interface{}
	Size() int
}

// Queue declares the struct for a Queue.
type Queue struct {
	queue *single.LinkedList // Private Single-Linked-List that manages the Queue's data.
	size  int                // Private integer representing the number of items stored in the Queue.
}

// New instantiates a new Queue pointer.
func New() *Queue {
	return &Queue{
		queue: single.New(),
		size:  0}
}

// Dequeue removes the value stored at the head of the Queue.
func (q *Queue) Dequeue() *Queue {
	q.queue.Shift()
	q.size--
	return q
}

// Enqueue adds a new value at the end of the Queue.
func (q *Queue) Enqueue(value interface{}) *Queue {
	q.queue.Append(value)
	q.size++
	return q
}

// IsEmpty checks whether the Queue is empty.
func (q *Queue) IsEmpty() bool {
	return q.queue.IsEmpty()
}

// IsPopulated checks whether the Queue contains a value.
func (q *Queue) IsPopulated() bool {
	return q.queue.IsPopulated()
}

// Peek returns the value stored at the head of the Queue. Peek does not modify the current Queue.
func (q *Queue) Peek() interface{} {
	if q.queue.HasHead() {
		return q.queue.Head.Value
	}
	return nil
}

// Poll removes the value stored at the head of the Queue and returns the value previously held by the Queue.
func (q *Queue) Poll() interface{} {
	q.size--
	return q.queue.ShiftGet()
}

// Size returns the sum of stored items held by the Queue.
func (q *Queue) Size() int {
	return q.size
}

var _ queue = (*Queue)(nil)

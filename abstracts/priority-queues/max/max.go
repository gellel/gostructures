// Package max exports a Max-Priorty Queue. Priority Queue holds
// collections of Priority Queue Nodes. Each Node contains an interface
// allowing data submission of any type and a priority integer.
package max

import (
	max "github.com/gellel/gostructures/abstracts/priority-queues/max/max-heap"
	node "github.com/gellel/gostructures/abstracts/priority-queues/max/max-node"
)

// Priority declares the implementation for a Max-Priority Queue.
type priority interface {
	Dequeue() interface{}
	Enqueue(priority int, value interface{}) int
	IsEmpty() bool
	IsNotEmpty() bool
	Length() int
	Peek() *node.Max
	PeekPriority() int
	PeekValue() interface{}
}

// Priority declares the Max-Priority Queue data structure.
type Priority struct {
	heap *max.Heap
}

// New instantiates a new Max-Priority Queue pointer.
func New() *Priority {
	return &Priority{
		heap: max.New()}
}

// Dequeue removes the element interface stored within the Queue's internal Heap.
func (p *Priority) Dequeue() interface{} {
	return p.heap.Poll()
}

// Enqueue pushes an element within the Queue's internal Heap and shuffles its position based on the argument priority.
func (p *Priority) Enqueue(priority int, value interface{}) int {
	return p.heap.Push(priority, value)
}

// IsEmpty checks whether the Max-Priority Queue is empty.
func (p *Priority) IsEmpty() bool {
	return p.heap.IsEmpty()
}

// IsNotEmpty checks whether the Max-Priorty Queue contains elements.
func (p *Priority) IsNotEmpty() bool {
	return p.heap.IsNotEmpty()
}

// Length computes the size of the Max-Priority Queue.
func (p *Priority) Length() int {
	return p.heap.Length()
}

// Peek accesses the first element interface stored in the Max-Priority Queue.
func (p *Priority) Peek() *node.Max {
	return p.heap.ToNode(p.heap.Peek())
}

// PeekPriority access the first element stored in the Max-Priorty Queue and returns its priority.
func (p *Priority) PeekPriority() int {
	n := p.Peek()
	if n != nil {
		return n.Priority
	}
	return -1
}

// PeekValue access the first element stored in the Max-Priorty Queue and returns its value.
func (p *Priority) PeekValue() interface{} {
	n := p.Peek()
	if n != nil {
		return n.Value
	}
	return nil
}

var _ priority = (*Priority)(nil)

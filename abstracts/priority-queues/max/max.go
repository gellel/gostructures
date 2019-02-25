package max

import (
	max "github.com/gellel/gostructures/abstracts/priority-queues/max/max-heap"
	node "github.com/gellel/gostructures/abstracts/priority-queues/max/max-node"
)

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

type Priority struct {
	heap *max.Heap
}

func New() *Priority {
	return &Priority{
		heap: max.New()}
}

func (p *Priority) Dequeue() interface{} {
	return p.heap.Poll()
}

func (p *Priority) Enqueue(priority int, value interface{}) int {
	return p.heap.Push(priority, value)
}

func (p *Priority) IsEmpty() bool {
	return p.heap.IsEmpty()
}

func (p *Priority) IsNotEmpty() bool {
	return p.heap.IsNotEmpty()
}

func (p *Priority) Length() int {
	return p.heap.Length()
}

func (p *Priority) Peek() *node.Max {
	return p.heap.ToNode(p.heap.Peek())
}

func (p *Priority) PeekPriority() int {
	n := p.Peek()
	if n != nil {
		return n.Priority
	}
	return -1
}

func (p *Priority) PeekValue() interface{} {
	n := p.Peek()
	if n != nil {
		return n.Value
	}
	return nil
}

// Package max exports a Max-Heap for use within a Max-Priority Queue.
// Uses the abstract Heap provided in the abstracts Heap package.
// Heap implements the abstract Heap and offers the BubbleUp and BubbleDown
// methods in the context of max-priority.
package max

import (
	"github.com/gellel/gostructures/abstracts/heaps/heap/container"
	"github.com/gellel/gostructures/abstracts/heaps/heap/core"
	node "github.com/gellel/gostructures/abstracts/priority-queues/max/max-node"
)

// Heap declares the interface for the Max-Heap.
type heap interface {
	BubbleDown(p int) int
	BubbleUp(p int) int
	Poll() interface{}
	Push(priority int, value interface{}) int
	PriorityOf(value interface{}) int
	ToNode(value interface{}) *node.Max
	ValueOf(value interface{}) interface{}
}

// Heap declares the structure for the Max-Heap abstract data structure.
type Heap struct {
	core.Heap
}

// New instantiates a new Max-Heap with the internal Heap container generated and assigned. Required otherwise methods will throw nil pointer exception.
func New() *Heap {
	h := &Heap{}
	h.Container = container.New()
	return h
}

// BubbleDown shuffles the element stored in the Heap at H[p] when the current element at H[Left] or H[Right] violates the Heap condition.
func (heap *Heap) BubbleDown(p int) int {
	x := p
	if heap.Container.BoundsLeftOf(p) && heap.PriorityOf(heap.Container.AccessLeftOf(p)) > heap.PriorityOf(heap.Container.Access(p)) {
		x = heap.Container.LeftOf(p)
	}
	if heap.Container.BoundsRightOf(p) && heap.PriorityOf(heap.Container.AccessRightOf(p)) > heap.PriorityOf(heap.Container.Access(x)) {
		x = heap.Container.RightOf(p)
	}
	if x != p {
		heap.Container.Swap(p, x)
		heap.BubbleDown(x)
	}
	return x
}

// BubbleUp shuffles the element stored in the Heap at H[P] when the current element at H[Parent] violates the Heap condition.
func (heap *Heap) BubbleUp(p int) int {
	if heap.Container.BoundsParentOf(p) && heap.PriorityOf(heap.Container.AccessParentOf(p)) < heap.PriorityOf(heap.Container.Access(p)) {
		heap.Container.Swap(heap.Container.ParentOf(p), p)
		heap.BubbleUp(heap.Container.ParentOf(p))
	}
	return p
}

// Poll accesses and returns the value held by the Max-Priority Node in the Max-Heap.
func (heap *Heap) Poll() interface{} {
	value := heap.ValueOf(heap.Heap.Poll())
	heap.BubbleDown(0)
	return value
}

// Push adds a new value into the Max-Heap, and assigns it position in the Heap container based on the priority argument.
func (heap *Heap) Push(priority int, value interface{}) int {
	return heap.BubbleUp(heap.Container.Push(node.New(priority, value)) - 1)
}

// PriorityOf accesses the priority of the Max-Heap Priority node.
func (heap *Heap) PriorityOf(value interface{}) int {
	return heap.ToNode(value).Priority
}

// ToNode casts the argument interface value to a Max-Heap node.
func (heap *Heap) ToNode(value interface{}) *node.Max {
	return value.(*node.Max)
}

// ValueOf accesses the value held by the Max-Heap Priority Node.
func (heap *Heap) ValueOf(value interface{}) interface{} {
	return heap.ToNode(value).Value
}

var _ heap = (*Heap)(nil)

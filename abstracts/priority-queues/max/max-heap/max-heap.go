package max

import (
	"github.com/gellel/gostructures/abstracts/heaps/heap/core"
	node "github.com/gellel/gostructures/abstracts/priority-queues/max/max-node"
)

type heap interface {
	BubbleDown(p int) int
	BubbleUp(p int) int
	Poll() interface{}
	Push(priority int, value interface{}) int
	PriorityOf(n *node.Max) int
	ToNode(value interface{}) *node.Max
	ValueOf(n *node.Max) interface{}
}

type Heap struct {
	core.Heap
}

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

func (heap *Heap) BubbleUp(p int) int {
	if heap.Container.BoundsParentOf(p) && heap.PriorityOf(heap.Container.AccessParentOf(p)) < heap.PriorityOf(heap.Container.Access(p)) {
		heap.Container.Swap(heap.Container.ParentOf(p), p)
		heap.BubbleUp(heap.Container.ParentOf(p))
	}
	return p
}

func (heap *Heap) Poll() interface{} {
	value := heap.ValueOf(heap.Heap.Poll())
	heap.BubbleDown(0)
	return value
}

func (heap *Heap) Push(priority int, value interface{}) int {
	return heap.BubbleUp(heap.Heap.Push(node.New(priority, value)) - 1)
}

func (heap *Heap) PriorityOf(value interface{}) int {
	return value.(node.Max).Priority
}

func (heap *Heap) ToNode(value interface{}) *node.Max {
	return value.(*node.Max)
}

func (heap *Heap) ValueOf(value interface{}) interface{} {
	return value.(node.Max).Value
}

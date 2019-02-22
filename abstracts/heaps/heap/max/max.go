package max

import (
	"github.com/gellel/gostructures/abstracts/heaps/heap/container"
	"github.com/gellel/gostructures/abstracts/heaps/heap/core"
)

type heap interface {
	BubbleDown(p int) int
	BubbleUp(p int) int
	Poll() int
	Push(value interface{}) int
	ToSlice() []int
}

type Heap struct {
	core.Heap
}

func New() *Heap {
	h := &Heap{}
	h.Container = container.New()
	return h
}

func (heap *Heap) BubbleDown(p int) int {
	x := p
	if heap.Container.BoundsLeftOf(p) && heap.ToInteger(heap.Container.AccessLeftOf(p)) > heap.ToInteger(heap.Container.Access(p)) {
		x = heap.Container.LeftOf(p)
	}
	if heap.Container.BoundsRightOf(p) && heap.ToInteger(heap.Container.AccessRightOf(p)) > heap.ToInteger(heap.Container.Access(x)) {
		x = heap.Container.RightOf(p)
	}
	if x != p {
		heap.Container.Swap(p, x)
		heap.BubbleDown(x)
	}
	return x
}

func (heap *Heap) BubbleUp(p int) int {
	if heap.Container.BoundsParentOf(p) && heap.ToInteger(heap.Container.AccessParentOf(p)) < heap.ToInteger(heap.Container.Access(p)) {
		heap.Container.Swap(heap.Container.ParentOf(p), p)
		heap.BubbleUp(heap.Container.ParentOf(p))
	}
	return p
}

func (heap *Heap) Poll() int {
	value := heap.Container.Poll()
	heap.BubbleDown(0)
	return value.(int)
}

func (heap *Heap) Push(value int) int {
	heap.Container.Push(value)
	return heap.BubbleUp(heap.Length() - 1)
}

func (heap *Heap) ToSlice() []int {
	a := make([]int, 0)
	for i := 0; i < heap.Length(); i++ {
		a = append(a, heap.Container.Access(i).(int))
	}
	return a
}

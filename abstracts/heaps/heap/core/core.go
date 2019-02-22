package core

import "github.com/gellel/gostructures/abstracts/heaps/heap/container"

type heap interface {
	CanMerge(h *Heap) bool
	CannotMerge(h *Heap) bool
	IsEmpty() bool
	IsNotEmpty() bool
	Length() int
	Merge(h *Heap) *Heap
	Peek() interface{}
	PeekAt(p int) interface{}
	PeekLast() interface{}
	Poll() interface{}
	Push(value interface{}) int
	Search(value interface{}) int
	ToSlice() []interface{}
}

type Heap struct {
	Container *container.Container
}

func New() *Heap {
	return &Heap{
		Container: container.New()}
}

func (heap *Heap) CanMerge(h *Heap) bool {
	return heap.Container.CanMerge(h.Container)
}

func (heap *Heap) CannotMerge(h *Heap) bool {
	return heap.Container.CannotMerge(h.Container)
}

func (heap *Heap) IsEmpty() bool {
	return heap.Container.IsEmpty()
}

func (heap *Heap) IsNotEmpty() bool {
	return heap.Container.IsNotEmpty()
}

func (heap *Heap) Length() int {
	return heap.Container.Length()
}

func (heap *Heap) Merge(h *Heap) *Heap {
	heap.Container.Merge(h.Container)
	return heap
}

func (heap *Heap) Peek() interface{} {
	return heap.Container.Peek()
}

func (heap *Heap) PeekAt(p int) interface{} {
	return heap.Container.PeekAt(p)
}

func (heap *Heap) PeekLast() interface{} {
	return heap.Container.PeekLast()
}

func (heap *Heap) Poll() interface{} {
	return heap.Container.Poll()
}

func (heap *Heap) Push(value interface{}) int {
	return heap.Container.Push(value)
}

func (heap *Heap) Search(value interface{}) int {
	return heap.Container.Search(value)
}

func (heap *Heap) ToSlice() []interface{} {
	return heap.Container.ToSlice()
}

var _ heap = (*Heap)(nil)

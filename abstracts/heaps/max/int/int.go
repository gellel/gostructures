package integer

type i interface {
	Empty() *Heap64
	Extract() uint64
	IsEmpty() bool
	IsLeaf(p int) bool
	LeftOf(p int) int
	Length() int
	Merge(h *Heap64) *Heap64
	NotEmpty() bool
	ParentOf(p int) int
	Peek(p int) int
	PeekFirst() int
	PeekLast() int
	Remove(value uint64) *Heap64
	RightOf(p int) int
	Set(a []uint64) *Heap64
	Unset() *Heap64
}

type Heap64 []uint64

func (heap *Heap64) Empty() *Heap64 {

	*heap = (*heap)[:0]

	return heap
}

func (heap *Heap64) Extract() uint64 {

	value := heap[0]
}
package integer

type i interface {
	Empty() *Heap64
	Extract() uint64
	Merge(h *Heap64) *Heap64
	Peek(p int) int
	PeekFirst() int
	PeekLast() int
	Set(a []uint64) *Heap64
	Unset() *Heap64
}

type Heap64 []uint64
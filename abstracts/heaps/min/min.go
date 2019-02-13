package min

type heap interface {
	Append(value uint) *Heap
	Bounds(p int) bool
	BoundsLeftOf(p int) bool
	BoundsRightOf(p int) bool
	BubbleDown(p int) *Heap
	BubbleUp(p int) *Heap
	Empty() *Heap
	Fill(a []uint) *Heap
	Filter(a []uint) []uint
	Insert(value int) *Heap
	IsEmpty() bool
	IsNotEmpty() bool
	LeftOf(p int) int
	Length() int
	Merge(h *Heap) *Heap
    ParentOf(p int) int
    Peek(p int) uint
    PeekFirst() uint
    PeekLast() uint
    PeekLeftOf(p int) uint
    PeekParentOf(p int) uint
	PeekRightOf(p int) uint
	Push(p uint) int
	Pop() uint
	RightOf(p int) int
	Search(value uint) int
	Swap(a int, b int) *Heap
}

func (heap *Heap) Append(value uint) *Heap {
	if value != 0 {
		*heap = append((*heap), value)
	}
	return heap
}

func (heap *Heap) Bounds(p int) bool {
	return ((p > -1) && p < len((*heap)))
}

func (heap *Heap) BoundsLeftOf(p int) bool {
	return heap.Bounds(heap.LeftOf(p))
}

func (heap *Heap) BoundsRightOf(p int) bool {
	return heap.Bounds(heap.RightOf(p))
}

func (heap *Heap) BubbleDown(p int) *Heap {
	if heap.BoundsLeftOf(p) && heap.PeekLeftOf(p) > heap.Peek(p) {
		
	}
}
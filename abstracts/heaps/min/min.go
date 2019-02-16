package min

type heap interface {
	Access(p int) uint
	Append(value uint) *Heap
	Bounds(p int) bool
	BoundsLeftOf(p int) bool
	BoundsParentOf(p int) bool
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

func (heap *Heap) Access(p int) uint {
	return (*heap)[p]
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

func (heap *Heap) BoundsParentOf(p int) bool {
	return heap.Bounds(heap.ParentOf(p))
}

func (heap *Heap) BoundsRightOf(p int) bool {
	return heap.Bounds(heap.RightOf(p))
}

func (heap *Heap) BubbleDown(p int) *Heap {
	x := p
	if heap.BoundsLeftOf(p) && heap.PeekLeftOf(p) > heap.Peek(p) {
		x = heap.LeftOf(p)
	}
	if heap.BoundsRightOf(p) && heap.PeekRightOf(p) > heap.Peek(x) {
		x = heap.RightOf(p)
	}
	if x != p {
		heap.Swap(p, x).BubbleDown(x)
	}
	return heap
}

func (heap *Heap) BubbleUp(p int) *Heap {
	if heap.BoundsParentOf(p) && heap.PeekParentOf(p) > heap.Peek(p) {
		heap.Swap(heap.ParentOf(p), p).BubbleUp(heap.ParentOf(p))
	}
	return heap
}

func (heap *Heap) Empty() *Heap {
	*heap = (*heap)[:0]
	return heap
}

func (heap *Heap) Fill(a []uint) *Heap {

	*heap = append((*heap)[:0], heap.Filter(a)...)

	for i := (heap.Length() - 1) / 2; i > -1; i-- {
		heap.BubbleDown(i)
	}
	return heap
}

func (heap *Heap) Filter(a []uint) []uint {
	b := make([]uint, 0)
	for _, n := range a {
		if n != 0 {
			b = append(b, n)
		}
	}
	return b
}

func (heap *Heap) IsEmpty() bool {
	return len((*heap)) == 0
}

func (heap *Heap) IsNotEmpty() bool {
	return len((*heap)) != 0
}

func (heap *Heap) LeftOf(p int) int {
	return ((p * 2) + 1)
}

func (heap *Heap) Merge(h *Heap) *Heap {
	
	*heap = append((*heap), (*h)...)

	return heap.BubbleUp(heap.Length() - 1)
}

func (heap *Heap) ParentOf(p int) int {
	return ((p - 1) / 2)
}

func (heap *Heap) Peek(p int) uint {
	if heap.Bounds(p) {
		return heap.Access(p)
	}
	return 0
}
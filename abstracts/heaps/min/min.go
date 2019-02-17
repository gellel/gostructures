package min

type heap interface {
	Access(p int) uint
	AccessLeftOf(p int) uint
	AccessParentOf(p int) uint
	AccessRightOf(p int) uint
	Append(value uint) *Heap
	Bounds(p int) bool
	BoundsLeftOf(p int) bool
	BoundsParentOf(p int) bool
	BoundsRightOf(p int) bool
	BubbleDown(p int) int
	BubbleUp(p int) int
	Empty() *Heap
	Fill(a []uint) *Heap
	Filter(a []uint) []uint
	Insert(value uint) *Heap
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
	ToSlice() []uint
}

type Heap []uint

func (heap *Heap) Access(p int) uint {
	return (*heap)[p]
}

func (heap *Heap) AccessLeftOf(p int) uint {
	return heap.Access(heap.LeftOf(p))
}

func (heap *Heap) AccessParentOf(p int) uint {
	return heap.Access(heap.ParentOf(p))
}

func (heap *Heap) AccessRightOf(p int) uint {
	return heap.Access(heap.RightOf(p))
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

func (heap *Heap) BubbleDown(p int) int {
	x := p
	if heap.BoundsLeftOf(p) && heap.AccessLeftOf(p) < heap.Access(p) {
		x = heap.LeftOf(p)
	}
	if heap.BoundsRightOf(p) && heap.AccessRightOf(p) < heap.Access(x) {
		x = heap.RightOf(p)
	}
	if x != p {
		heap.Swap(p, x).BubbleDown(x)
	}
	return x
}

func (heap *Heap) BubbleUp(p int) int {
	if heap.BoundsParentOf(p) && heap.AccessParentOf(p) > heap.Access(p) {
		heap.Swap(heap.ParentOf(p), p).BubbleUp(heap.ParentOf(p))
	}
	return p
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

func (heap *Heap) Insert(value uint) *Heap {

	heap.Push(value)

	return heap
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

func (heap *Heap) Length() int {
	return len((*heap))
}

func (heap *Heap) Merge(h *Heap) *Heap {

	*heap = append((*heap), (*h)...)

	heap.BubbleUp(heap.Length() - 1)

	return heap
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

func (heap *Heap) PeekFirst() uint {
	if heap.IsNotEmpty() {
		return heap.Access(0)
	}
	return 0
}

func (heap *Heap) PeekLast() uint {
	if heap.IsNotEmpty() {
		return heap.Access(heap.Length() - 1)
	}
	return 0
}

func (heap *Heap) PeekLeftOf(p int) uint {
	if heap.BoundsLeftOf(p) {
		return heap.Access(heap.LeftOf(p))
	}
	return 0
}

func (heap *Heap) PeekParentOf(p int) uint {
	if heap.BoundsParentOf(p) {
		return heap.Access(heap.ParentOf(p))
	}
	return 0
}

func (heap *Heap) PeekRightOf(p int) uint {
	if heap.BoundsRightOf(p) {
		return heap.Access(heap.RightOf(p))
	}
	return 0
}

func (heap *Heap) Push(value uint) int {

	heap.Append(value)

	return heap.BubbleUp(heap.Length() - 1)
}

func (heap *Heap) Pop() uint {
	if heap.IsEmpty() {
		return 0
	}

	k := heap.Access(0)

	*heap = (*heap)[1:]

	heap.BubbleDown(0)

	return k
}

func (heap *Heap) RightOf(p int) int {
	return ((p * 2) + 2)
}

func (heap *Heap) Search(value uint) int {
	for i := 0; i < heap.Length(); i++ {
		if heap.Access(i) == value {
			return i
		}
	}
	return -1
}

func (heap *Heap) Swap(a int, b int) *Heap {
	(*heap)[a], (*heap)[b] = (*heap)[b], (*heap)[a]
	return heap
}

func (heap *Heap) ToSlice() []uint {
	a := make([]uint, 0)
	for i := 0; i < heap.Length(); i++ {
		a = append(a, heap.Access(i))
	}
	return a
}

var _ heap = (*Heap)(nil)

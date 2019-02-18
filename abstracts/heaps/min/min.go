// Package min exports a Min-Heap. Heap is programmed to only support use
// of unsigned integers, and filters out any zero value arguments. Use of
// zero and negative integer is used for core functionality, such as being
// able to determine whether a value is absent in the Heap.
package min

// Heap declares the implementation for a Min-Heap
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

// Heap declares the Min-Heap data structure.
type Heap []uint

// Access accesses and returns the stored element in the Heap at position P.
func (heap *Heap) Access(p int) uint {
	return (*heap)[p]
}

// AccessLeftOf accesses and returns the stored element in the Heap at the left-of position P.
func (heap *Heap) AccessLeftOf(p int) uint {
	return heap.Access(heap.LeftOf(p))
}

// AccessParentOf accesses and returns the stored element in the Heap at the parent position of P.
func (heap *Heap) AccessParentOf(p int) uint {
	return heap.Access(heap.ParentOf(p))
}

// AccessRightOf accesses and returns the stored element in the Heap at the right-of position P.
func (heap *Heap) AccessRightOf(p int) uint {
	return heap.Access(heap.RightOf(p))
}

// Append adds a non-zero unsigned integer into the Heap. Does not BubbleUp the added value. Use either Heap.Insert or Heap.Push.
func (heap *Heap) Append(value uint) *Heap {
	if value != 0 {
		*heap = append((*heap), value)
	}
	return heap
}

// Bounds checks whether the argument position has a corresponding position within the Heap.
func (heap *Heap) Bounds(p int) bool {
	return ((p > -1) && p < len((*heap)))
}

// BoundsLeftOf checks whether the argument position has a corresponding position within the Heap when modified to be the left-of key for the argument key.
func (heap *Heap) BoundsLeftOf(p int) bool {
	return heap.Bounds(heap.LeftOf(p))
}

// BoundsParentOf checks whether the argument position has a corresponding position within the Heap when modified to be the parent key for the argument key.
func (heap *Heap) BoundsParentOf(p int) bool {
	return heap.Bounds(heap.ParentOf(p))
}

// BoundsRightOf checks whether the argument position has a corresponding position within the Heap when modified to be the right-of key for the argument key.
func (heap *Heap) BoundsRightOf(p int) bool {
	return heap.Bounds(heap.RightOf(p))
}

// BubbleDown shuffles the element stored in the Heap at H[p] when the current element at H[Left] or H[Right] violates the Heap condition.
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

// BubbleUp shuffles the element stored in the Heap at H[P] when the current element at H[Parent] violates the Heap condition.
func (heap *Heap) BubbleUp(p int) int {
	if heap.BoundsParentOf(p) && heap.AccessParentOf(p) > heap.Access(p) {
		heap.Swap(heap.ParentOf(p), p).BubbleUp(heap.ParentOf(p))
	}
	return p
}

// Empty removes all entries from the Heap.
func (heap *Heap) Empty() *Heap {
	*heap = (*heap)[:0]
	return heap
}

// Fill populates the the Heap with a Slice of unsigned integers and then sorts the Heap's children to satisfy the Heap condition.
func (heap *Heap) Fill(a []uint) *Heap {

	*heap = append((*heap)[:0], heap.Filter(a)...)

	for i := (heap.Length() - 1) / 2; i > -1; i-- {
		heap.BubbleDown(i)
	}
	return heap
}

// Filter empties zero-values from the argument unsigned integer Slice.
func (heap *Heap) Filter(a []uint) []uint {
	b := make([]uint, 0)
	for _, n := range a {
		if n != 0 {
			b = append(b, n)
		}
	}
	return b
}

// Insert pushes a non-zero unsigned integer into the Heap.
func (heap *Heap) Insert(value uint) *Heap {

	heap.Push(value)

	return heap
}

// IsEmpty checks whether the Heap is absent of elements.
func (heap *Heap) IsEmpty() bool {
	return len((*heap)) == 0
}

// IsNotEmpty checks whether the Heap contains at least on element.
func (heap *Heap) IsNotEmpty() bool {
	return len((*heap)) != 0
}

// LeftOf computes the position where the value left-of the argument position is stored in the Heap.
func (heap *Heap) LeftOf(p int) int {
	return ((p * 2) + 1)
}

// Length computes and returns the current number of elements contained within the Heap. Size represents the non-zero-index value of the Heap.
func (heap *Heap) Length() int {
	return len((*heap))
}

// Merge joins two Heaps into a single Heap.
func (heap *Heap) Merge(h *Heap) *Heap {

	*heap = append((*heap), (*h)...)

	heap.BubbleUp(heap.Length() - 1)

	return heap
}

// ParentOf computes the position where the value parent of the argument position is stored in the Heap.
func (heap *Heap) ParentOf(p int) int {
	return ((p - 1) / 2)
}

// Peek safely accesses an element stored in the Heap when the argument position is in the bounds of the Heap. If argument position cannot be accessed, a zero-value will be returned.
func (heap *Heap) Peek(p int) uint {
	if heap.Bounds(p) {
		return heap.Access(p)
	}
	return 0
}

// PeekFirst safely accesses the element stored at the beginning of the Heap when the Heap is not empty. Otherwise, returns a zero-value.
func (heap *Heap) PeekFirst() uint {
	if heap.IsNotEmpty() {
		return heap.Access(0)
	}
	return 0
}

// PeekLast safely accesses the element stored at the end of the Heap when the Heap is not empty. Otherwise, returns a zero-value.
func (heap *Heap) PeekLast() uint {
	if heap.IsNotEmpty() {
		return heap.Access(heap.Length() - 1)
	}
	return 0
}

// PeekLeftOf safely access the element stored left-of the argument position when the argument left-of is in the Heap bounds. Otherwise returns zero.
func (heap *Heap) PeekLeftOf(p int) uint {
	if heap.BoundsLeftOf(p) {
		return heap.Access(heap.LeftOf(p))
	}
	return 0
}

// PeekParentOf safely access the element stored at the parent of the argument position when the parent position is in the Heap bounds. Otherwise returns zero.
func (heap *Heap) PeekParentOf(p int) uint {
	if heap.BoundsParentOf(p) {
		return heap.Access(heap.ParentOf(p))
	}
	return 0
}

// PeekRightOf safely access the element stored right-of the argument position when the argument right-of is in the Heap bounds. Otherwise returns zero.
func (heap *Heap) PeekRightOf(p int) uint {
	if heap.BoundsRightOf(p) {
		return heap.Access(heap.RightOf(p))
	}
	return 0
}

// Push adds a new, non-zero unsigned integer in the Heap and shifts the new unsigned integer into its required position if it violates the Heap condition.
func (heap *Heap) Push(value uint) int {

	heap.Append(value)

	return heap.BubbleUp(heap.Length() - 1)
}

// Pop removes the first element from the Heap and shuffles the potential element at H[0] to its required position if it violates the Heap condition. If the Heap is emtpy, it returns zero.
func (heap *Heap) Pop() uint {
	if heap.IsEmpty() {
		return 0
	}

	k := heap.Access(0)

	*heap = (*heap)[1:]

	heap.BubbleDown(0)

	return k
}

// RightOf computes the position right-of the argument position.
func (heap *Heap) RightOf(p int) int {
	return ((p * 2) + 2)
}

// Search iterates across the entries of the Heap and attempts to find the index where the argument unsigned integer is stored.
func (heap *Heap) Search(value uint) int {
	for i := 0; i < heap.Length(); i++ {
		if heap.Access(i) == value {
			return i
		}
	}
	return -1
}

// Swap moves the element stored at Heap[A] to Heap[B] and Heap[B] to Heap[A].
func (heap *Heap) Swap(a int, b int) *Heap {
	(*heap)[a], (*heap)[b] = (*heap)[b], (*heap)[a]
	return heap
}

// ToSlice prodcues a new Slice of unsigned integers built from the values stored in the Heap.
func (heap *Heap) ToSlice() []uint {
	a := make([]uint, 0)
	for i := 0; i < heap.Length(); i++ {
		a = append(a, heap.Access(i))
	}
	return a
}

var _ heap = (*Heap)(nil)

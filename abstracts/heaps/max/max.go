// Package max exports a Max-Heap. Heap is programmed to only support use
// of unsigned integers, and filters out any zero value arguments. Use of
// zero and negative integer is used for core functionality, such as being
// able to determine whether a value is absent in the Heap. 
package max

import "fmt"

// Heap declares the implementation for a Max-Heap
type heap interface {
	Access(p int) uint
	Append(value uint) *Heap
	Arrange(p int) int
	Bounds(p int) bool
	Empty() *Heap
	Fill(a []uint) *Heap
	Filter(a []uint) []uint
	Insert(value uint) *Heap
	IsEmpty() bool
	IsLeaf(p int) bool
	IsNotEmpty() bool
	Left(p int) int
	Length() int
	Merge(h *Heap) *Heap
	Parent(p int) int
	Peek(p int) uint
	PeekFirst() uint
	PeekLast() uint
	Pop() uint
	Push(value uint) int
	Search(value uint) int
	ShiftDown(a int, b int) *Heap
	ShiftUp(b int, a int) *Heap
	Swap(i int, j int) *Heap
	Violates(value uint) error
}

// Heap declares the Max-Heap data structure.
type Heap []uint

// New instantiates a new Max-Heap. 
func New() *Heap {
	return &Heap{}
}

// Access reaches into Heap and attempts to read value at argument position, without checking whether the position is in Heap bounds.
func (heap *Heap) Access(p int) uint {
	return (*heap)[p]
}

// Append assigns a new unsigned integer to the end of the Heap, without bubbling the new value.
func (heap *Heap) Append(value uint) *Heap {

	*heap = append((*heap), value)

	return heap
}

// Arrange orders the Heap into a max Heap.
func (heap *Heap) Arrange(p int) int {
	for heap.Peek(p) > heap.Peek(heap.Parent(p)) {
		p = heap.Swap(p, heap.Parent(p)).Parent(p)
	}
	return p
}

// Bounds returns a boolean checking whether the argument position does not overflow or underflow the current Heap length.
func (heap *Heap) Bounds(p int) bool {
	return (p > -1) && (p < heap.Length())
}

// Empty removes all elements held by the Heap.
func (heap *Heap) Empty() *Heap {

	*heap = (*heap)[:0]

	return heap
}

// Fill moves elements from an argument slice into the Heap, resetting the Heap in the process.
func (heap *Heap) Fill(a []uint) *Heap {

	*heap = (*heap)[:0]

	*heap = append((*heap), heap.Filter(a)...)

	return heap
}

// Filter empties zeroes from argument slice.
func (heap *Heap) Filter(a []uint) []uint {

	b := make([]uint, 0)

	for _, n := range a {
		if n != 0 {
			b = append(b, n)
		}
	}
	return b
}

// Insert pushes a new unsigned integer into the Heap, but returns the Heap itself and not the position.
func (heap *Heap) Insert(value uint) *Heap {

	heap.Push(value)

	return heap
}

// IsEmpty checks the size of the Heap, determining whether it is empty or populated.
func (heap *Heap) IsEmpty() bool {
	return (heap.Length() == 0)
}

// IsLeaf checks the argument position is a leaf in the Heap.
func (heap *Heap) IsLeaf(p int) bool {
	return ((p >= (heap.Length() / 2)) && (p <= heap.Length()))
}

// IsNotEmpty checks the size of the Heap, checking whether it contains at least one element.
func (heap *Heap) IsNotEmpty() bool {
	return (heap.Length() > 0)
}

// Left calculates the left position in a binary search, relative to argument position.
func (heap *Heap) Left(p int) int {
	return ((p * 2) + 1)
}

// Length returns the number of elements stored in the current Heap.
func (heap *Heap) Length() int {
	return len((*heap))
}

// Merge concatenates accessed Heap with argument Heap and then orders the combined Heaps.
func (heap *Heap) Merge(h *Heap) *Heap {

	*heap = append((*heap), (*h)...)

	heap.Arrange(heap.Length() - 1)

	return heap
}

// Parent calculates the sum of argument position that would be its relative parent.
func (heap *Heap) Parent(p int) int {
	return (p / 2)
}

// Peek attempts to access the unsigned integer stored at the argument index and returns 0 when Heap fails the lookup.
func (heap *Heap) Peek(p int) uint {
	if heap.Bounds(p) {
		return heap.Access(p)
	}
	return 0
}

// PeekFirst attempts to access the unsigned integer stored at the beginning of the Heap; the heap-max.
func (heap *Heap) PeekFirst() uint {
	return heap.Peek(0)
}

// PeekLast attempts to access the unsigned integer stored at the end of the Heap; the heap-max-min.
func (heap *Heap) PeekLast() uint {
	return heap.Peek(heap.Length() - 1)
}

// Pop unsets the max of the Max-Heap and arranges the Heap before returning value. Returns zero if Heap is empty.
func (heap *Heap) Pop() uint {

	if heap.IsEmpty() {
		return 0
	}

	k := (*heap)[0]

	*heap = (*heap)[1:]

	return k
}

// Push adds a new unsigned integer into the Heap, then bubbles the value to the correct position before returning its index.
func (heap *Heap) Push(value uint) int {

	heap.Append(value)
	
	return heap.Arrange(heap.Length() - 1)
}

// Right calculates right left position in a binary search, relative to argument position.
func (heap *Heap) Right(p int) int {
	return ((p * 2) + 2)
}

// Search binary searches across the Heap to find an unsigned integer that matches the argument value in the Heap.
func (heap *Heap) Search(value uint) int {
	k := 0
	for heap.Bounds(k) {
		v := heap.Peek(k)
		if v == value {
			return k
		} else if v < value {
			k = heap.Left(k)
		} else {
			k = heap.Right(k)
		}
	}
	return -1
}

// ShiftDown moves position a to b.
func (heap *Heap) ShiftDown(a int, b int) *Heap {
	return heap.Swap(a, b)
}

// ShiftUp moves position b to a.
func (heap *Heap) ShiftUp(b int, a int) *Heap {
	return heap.Swap(b, a)
}

// Swap shuffles the references of i and j, putting i in j's position and j in i's.
func (heap *Heap) Swap(i int, j int) *Heap {
	if heap.Bounds(i) && heap.Bounds(j) {
		(*heap)[i], (*heap)[j] = heap.Peek(j), heap.Peek(i)
	}
	return heap
}

// Violates checks the argument value and generates an error when the value does not satisfy the accepted Heap values.
func (heap *Heap) Violates(value uint) error {
	if value == 0 {
		return fmt.Errorf("(*%p) cannot hold unsigned int of zero. Heap reserves value to determine absence in Heap set", &heap)
	}
	return nil
}

var _ heap = (*Heap)(nil)

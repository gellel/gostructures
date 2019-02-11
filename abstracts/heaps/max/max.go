package max

import "fmt"

type heap interface {
	Access(p int) uint
	Bounds(p int) bool
	Empty() *Heap
	Fill(a []uint) *Heap
	Filter(a []uint) []uint
	Left(p int) int
	Length() int
	Merge(h *Heap) *Heap
	Parent(p int) int
	Peek(p int) uint
	PeekFirst() uint
	PeekLast() uint
	Swap(i int, j int) *Heap
	Violates(value uint) error
}

// Heap declares the Max-Heap data structure.
type Heap []uint

// Access reaches into Heap and attempts to read value at argument position, without checking whether the position is in Heap bounds.
func (heap *Heap) Access(p int) uint {
	return (*heap)[p]
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

// Left calculates the left position in a binary search, relative to argument position.
func (heap *Heap) Left(p int) int {
	return p * 2
}

// Length returns the number of elements stored in the current Heap.
func (heap *Heap) Length() int {
	return len((*heap))
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

// Merge concatenates accessed Heap with argument Heap.
func (heap *Heap) Merge(h *Heap) *Heap {

	*heap = append((*heap), (*h)...)

	return heap
}

// Right calculates right left position in a binary search, relative to argument position.
func (heap *Heap) Right(p int) int {
	return ((p * 2) + 1)
}

// Swap shuffles the references of i and j, putting i in j's position and j in i's.
func (heap *Heap) Swap(i int, j int) *Heap {
	if heap.Bounds(i) && heap.Bounds(j) {
		(*heap)[i], (*heap)[j] = heap.Peek(i), heap.Peek(j)
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

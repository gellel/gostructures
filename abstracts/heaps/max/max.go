package max

import "fmt"

type f64 interface {
	Arrange(i int) *F64
	Contains(f float64) bool
	Empty() *F64
	Heapify(f []float64) *F64
	IndexOf(f float64) int
	IsEmpty() bool
	Insert(f float64) *F64
	IsLeaf(i int) bool
	IsPopulated() bool
	LeftOf(i int) int
	Len() int
	Merge(f *F64) *F64
	ParentOf(i int) int
	Peek() float64
	PeekAt(i int) float64
	Pop() float64
	RightOf(i int) int
	Search(f float64) float64
	SumOf() float64
	Swap(a int, b int) *F64
	ViolatesLength(i int) error
}

type F64 []float64


func (f64 *F64) Arrange(i int) *F64 {
	if f64.IsLeaf(i) {
		return f64
	}

	left := f64.LeftOf(i)

	right := f64.RightOf(i)

	if f64.PeekAt(i) < f64.PeekAt(left) || f64.PeekAt(i) > f64.PeekAt(right) {
		if f64.PeekAt(left) > f64.PeekAt(right) {
			f64.Swap(i, left).Arrange(left)
		} else {
			f64.Swap(i, right).Arrange(right)
		}
	}
	return f64
}

func (f64 *F64) Contains(n float64) bool {
	return false
}

func (f64 *F64) Empty() *F64 {

	*f64 = (*f64)[:0]

	return f64
}

func (f64 *F64) IndexOf(f float64) int {
	if f64.Len() > 0 {
		return 0
	}
	return -1
}

func (f64 *F64) IsEmpty() bool {
	return f64.Len() == 0
}

func (f64 *F64) Insert(f float64) *F64 {

	*f64 = *append((*f64), f)

	for n := f64[:f64.Len()]; n > f64.ParentOf(n) {
		
		f64.Swap(n, f64.ParentOf(n))

		n = f64.ParentOf(n)
	}
	return f64
}

func (f64 *F64) IsLeaf(i int) bool {
	return ((i >= (f64.Len() / 2)) && (i <= f64.Len()))
}

func (f64 *F64) IsPopulated() bool {
	return (f64.Len() > 0)
}

func (f64 *F64) Heapify(f []float64) *F64 {
	return f64
}

func (f64 *F64) LeftOf(i int) int {
	return (i * 2)
}

func (f64 *F64) Len() int {
	return len(*f64)
}

func (f64 *F64) Merge(f *F64) *F64 {

	*f64 = *append((*f64), (*f)...)

	return f64
}

func (f64 *F64) ParentOf(i int) int {
	return (i / 2)
}

func (f64 *F64) Peek() float64 {
	return f64.PeekAt(0)
}

func (f64 *F64) PeekAt(i int) float64 {
	return (*f64)[i]
}

func (f64 *F64) Pop() float64 {
	return 0.0
}

func (f64 *F64) RightOf(i int) int {
	return ((i * 2) + i)
}

func (f64 *F64) Search(n float64) float64 {
	return -1.0
}

func (f64 *F64) SumOf() float64 {
	return 0.0
}

func (f64 *F64) Swap(a int, b int) *F64 {
	
	err := f64.ViolatesLength(a)

	if err != nil {
		fmt.Println(err)
	}

	err = f64.ViolatesLength(b)

	if err != nil {
		fmt.Println(err)
	}

	(*f64)[a], (*f64)[b] = (*f64)[b], (*f64)[a]

	return f64
}

func (f64 *F64) ViolatesLength(i int) error {

	message := "(*%p) cannot reach argument position. argument index out of heap bounds. current range is 0-%d"

	if ((i < 0) || (i > f64.Len())) {
		return fmt.Errorf(message, &f64, f64.Len())
	}
	return nil
}

var _ f64 = (*F64)(nil)
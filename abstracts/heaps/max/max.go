package max

type f64 interface {
	Contains(f float64) bool
	Empty() *F64
	Heapify(f []float64) *F64
	IndexOf(f float64) int
	LeftOf(i int) int
	Len() int
	ParentOf(i int) int
	Peek() float64
	Pop() float64
	Push(f float64) *F64
	RightOf(i int) int
	Search(f float64) float64
	SumOf() float64
}

type F64 []float64


func (f64 *F64) Contains(n float64) bool {
	return false
}

func (f64 *F64) Empty() *F64 {
	return f
}

func (f64 *F64) IndexOf() int {
	if f64.Len() > 0 {
		return 0
	}
	return -1
}

func (f64 *F64) Heapify() *F64 {
	return f64
}

func (f64 *F64) LeftOf(i int) int {
	return (i * 2)
}

func (f64 *F64) Len() int {
	return len(*f)
}

func (f64 *F64) ParentOf(i int) int {
	return (i / 2)
}

func (f64 *F64) Peek() float64 {
	return -1.0
}

func (f64 *F64) Pop() float64 {
	return 0.0
}

func (f64 *F64) RightOf(i int) int {
	return ((i * 2) + i)
}

func (f64 *F64) Search(n float64) float64 {
	return nil
}

func (f64 *F64) SumOf() float64 {
	return 0.0
}

var _ f64 = (*F64)(nil)
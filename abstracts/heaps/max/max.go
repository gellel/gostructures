package max

type f64 interface {
	Contains(f float64) bool
	Empty() *F64
	Heapify(f []float64) *F64
	IndexOf(f float64) int
	LeftOf(i int) int
	Len() int
	Merge(f *F64) *F64
	ParentOf(i int) int
	Peek() float64
	Pop() float64
	Push(f float64) *F64
	RightOf(i int) int
	Search(f float64) float64
	SumOf() float64
	Swap(a int, b int) *F64
	ViolatesLen(i int) error
	ViolatesMax(f float64) error
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

func (f64 *F64) Merge(f *F64) *F64 {

	*f64 = append(*f64, f...)

	return f64
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

func (f64 *F64) Swap(a int, b int) float64 {
	
	f64.ViolatesLen(a)

	f64.ViolatesLen(b)

	*f64[a], *f64[b] = f64[b], f64[a]

	return f64
}

func (f64 *F64) ViolatesLen(i int) error {

	if a < 0 || b < 0 {
		return error
	}
	if a > f64.Len() || b > f64.Len() {
		return error
	}
	return nil
}

var _ f64 = (*F64)(nil)
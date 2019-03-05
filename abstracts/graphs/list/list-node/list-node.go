package node

type list interface{}

type List struct {
	Vertex int
	Next   *List
}

func New() *List {
	return &List{}
}

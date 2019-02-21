package container

type container interface {
	Peek() interface{}
	PeekAt(p int) interface{}
	PeekFirst() interface{}
	PeekLast() interface{}
	Poll() interface{}
	Push(value interface{}) int
}

type Container []interface{}

var _ container = (*Container)(nil)

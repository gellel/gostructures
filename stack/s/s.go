package s

type Stack interface {
	InStack() bool
	IsEmpty() bool
	Peek() interface{}
	Pop() interface{}
	Push() interface{}
}

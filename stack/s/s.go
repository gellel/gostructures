package s

type Stack interface {
	IsEmpty() bool
	InStack() bool
	Peek() interface{}
	Pop() interface{}
	Push() interface{}
}

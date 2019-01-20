package s

// Stack decleration.
type Stack interface {
	InStack() bool
	IsEmpty() bool
	Peek() interface{}
	Pop() interface{}
	Push(property interface{}) interface{}
}

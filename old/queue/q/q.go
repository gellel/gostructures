package q

// Queue decleration.
type Queue interface {
	Enqueue(i interface{}) (interface{}, int)
	Dequeue(i interface{}) (interface{}, int)
	InQueue(i int) bool
	Peek(i int) interface{}
	SizeOf() int
}

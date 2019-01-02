package q

type Q interface {
	Enqueue(i interface{}) (Q, int)
	Dequeue(i interface{}) (Q, int)
	InQueue(i int) bool
	Peek(i int) interface{}
	SizeOf() int
}

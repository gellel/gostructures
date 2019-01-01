package q

type Q interface {
	Enqueue(i interface{}) Q
	Dequeue(i interface{}) Q
	Peek(i int) interface{}
	Size() int
}

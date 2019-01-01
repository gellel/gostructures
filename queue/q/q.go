package q

type Q interface {
	Enqueue(i interface{}) (Q, int)
	Dequeue(i interface{}) (Q, int)
	Peek(i int) interface{}
	Size() int
}

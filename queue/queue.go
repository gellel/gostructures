package queue

import (
	"github.com/gellel/gostructures/queue/q"
)

// Queue of interfaces.
type Queue struct {
	q.Q
	slices []interface{}
}

func (s Queue) Enqueue(property interface{}) (Queue, interface{}) {
	a := append(s.slices, property)
	return s, a
}

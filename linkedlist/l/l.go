package l

import (
	"github.com/gellel/gostructures/linkedlist/node/n"
)

type LinkedList interface {
	Add(property interface{}) *n.Node
	Delete() int
	Prepend(property interface{}) *n.Node
	Reverse()
	Search() bool
	Traverse()
}

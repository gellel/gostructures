package l

import (
	"github.com/gellel/gostructures/linkedlist/node"
)

// LinkedList decleration.
type LinkedList interface {
	Add(property interface{}) *node.Node
	Delete() *node.Node
	Prepend(property interface{}) *node.Node
	Reverse()
	Search() bool
	SizeOf() int
	Walk()
}

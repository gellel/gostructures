package l

import (
	"github.com/gellel/gostructures/linked-list/node"
)

// LinkedList decleration.
type LinkedList interface {
	Add(property interface{}) *LinkedList
	Append(property interface{}) *LinkedList
	Delete() *node.Node
	Prepend(property interface{}) *LinkedList
	Reverse() *LinkedList
	Search(value interface{}) bool
	SizeOf() int
	Walk()
}

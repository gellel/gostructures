package node

import (
	"github.com/gellel/gostructures/linked-list/node/n"
)

// Node of LinkedList.
type Node struct {
	n.Node
	Next  *Node
	Value interface{}
}

// New instantiates a new LinkedList.Node.
func New(property interface{}) *Node {
	return &Node{Value: property}
}

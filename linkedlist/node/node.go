package node

import (
	"github.com/gellel/gostructures/linkedlist/node/n"
)

type Node struct {
	n.Node
	Next  *Node
	Value interface{}
}

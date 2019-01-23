package node

import node "github.com/gellel/gostructures/abstracts/trees/binary/binary-node"

type A interface {
	node.B
	BalanceFactor() int
	RotateLeft() *AVL
	RotateRight() *AVL
}

type AVL struct {
	node.Binary
}

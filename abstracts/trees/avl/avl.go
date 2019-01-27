package avl

import (
	node "github.com/gellel/gostructures/abstracts/trees/avl/avl-node"
)

type avl interface {
	Balance() int
	Find(value float64) *node.AVL
	Height() int
	Insert(value float64) *AVL
	Remove(value float64) *AVL
}

type AVL struct {
	Root *node.AVL
}

func New(value float64) *AVL {
	return &AVL{Root: node.New(value).UnsafelyAssignSide("ROOT")}
}

var _ avl = (*AVL)(nil)

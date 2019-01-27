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
	ToAVLSlice() []*node.AVL
	ToFloatSlice() []float64
	Walk() *AVL
}

type AVL struct {
	Root *node.AVL
}

func New(value float64) *AVL {
	return &AVL{Root: node.New(value).UnsafelyAssignSide("ROOT")}
}

func (avl *AVL) Balance() int {
	return avl.Root.Balance()
}

func (avl *AVL) Find(value float64) *node.AVL {
	return avl.Root.Find(value)
}

func (avl *AVL) Height() int {
	return avl.Root.Height()
}

func (avl *AVL) Insert(value float64) *AVL {
	avl.Root.Insert(value)
	return avl
}

func (avl *AVL) Remove(value float64) *AVL {
	avl.Root.Remove(value)
	return avl
}

func (avl *AVL) ToAVLSlice() []*node.AVL {
	return avl.Root.ToAVLSlice()
}

func (avl *AVL) ToFloatSlice() []float64 {
	return avl.Root.ToFloatSlice()
}

func (avl *AVL) Walk() *AVL {
	avl.Root.Walk()
	return avl
}

var _ avl = (*AVL)(nil)

package binarynode

import (
	treenode "github.com/gellel/gostructures/trees/tree-node"
)

type BinaryNode struct {
	*treenode.Node
	Parent *treenode.Node
	Left   *treenode.Node
	Right  *treenode.Node
	Value  interface{}
}

func New(value interface{}) *BinaryNode {
	return &BinaryNode{
		Parent: nil,
		Left:   nil,
		Right:  nil,
		Value:  value}
}

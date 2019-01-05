package binarynode

import (
	treenode "github.com/gellel/gostructures/trees/tree-node"
)

// BinaryNode extends treenode.Node.
// Expresses a Binary Search Tree Node.
type BinaryNode struct {
	*treenode.Node
}

// New instantiates a BinaryNode.
func New(value int) *BinaryNode {
	return &BinaryNode{
		treenode.New(value)}
}

func (node *BinaryNode) Insert(value int) *BinaryNode {
	return insert(value, node, node.Node)
}

func (node *BinaryNode) Maximum() *BinaryNode {
	return maximum(node)
}

func (node *BinaryNode) Minimum() *BinaryNode {
	return minimum(node)
}

func insert(value int, binaryNode *BinaryNode, node *treenode.Node) *BinaryNode {
	if value < node.Value {
		if node.HasLeft() {
			return insert(value, binaryNode, node.Left)
		}

		node.SetLeft(treenode.New(value))

	} else if value > node.Value {
		if node.HasRight() {
			return insert(value, binaryNode, node.Right)
		}

		node.SetRight(treenode.New(value))
	}
	return binaryNode
}

func maximum(binaryNode *BinaryNode) *BinaryNode {
	if binaryNode.HasEmptyRight() {
		return binaryNode
	}
	return maximum(binaryNode)
}

func minimum(binaryNode *BinaryNode) *BinaryNode {
	if binaryNode.HasEmptyLeft() {
		return binaryNode
	}
	return minimum(binaryNode.Left)
}

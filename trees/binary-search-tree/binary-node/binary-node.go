package binarynode

import (
	treenode "github.com/gellel/gostructures/trees/tree-node"
)

// BinaryNode extends treenode.Node.
// Expresses a Binary Search Tree Node.
type BinaryNode struct {
	*treenode.Node
	Parent *BinaryNode
	Left   *BinaryNode
	Right  *BinaryNode
	Value  int
}

// New.
func New(value int) *BinaryNode {
	return &BinaryNode{
		Parent: nil,
		Left:   nil,
		Right:  nil,
		Value:  value}
}

// Insert creates and assigns a Child BinaryNode based
// on the weight of the BinaryNode.Value.
// Core method of the Binary Search Tree.
func (binaryNode *BinaryNode) Insert(value int) *BinaryNode {

	// check that the new node should be left
	// balanced.
	if value < binaryNode.Value {
		// if we already have a BinaryNode.Left
		// move through the Left Search Tree.
		if binaryNode.HasLeft() {
			return binaryNode.Left.Insert(value)
		}
		return binaryNode.SetBinaryLeft(value)
	}
	// otherwise check that the new node
	// should be balanced right.
	if value > binaryNode.Value {
		// if we already have a BinaryNode.Left
		// move through the Right Search Tree.
		if binaryNode.HasRight() {
			return binaryNode.Right.Insert(value)
		}
		return binaryNode.SetBinaryRight(value)
	}
	return binaryNode
}

// SetBinaryLeft creates BinaryNode and satisfies assignment
// to inherited Struct (treenode.Node) and sets parent dependencies.
func (binaryNode *BinaryNode) SetBinaryLeft(value int) *BinaryNode {

	n := New(value)
	// use inherited method of treenode.Node.
	// requires passing back implemented member.
	// to tree.Node.SetLeft.
	binaryNode.SetLeft(n.Node)
	// return the BinaryNode.
	return n
}

// SetBinaryRight creates BinaryNode and satisfies assignment
// to inherited Struct (treenode.Node) and sets parent dependencies.
func (binaryNode *BinaryNode) SetBinaryRight(value int) *BinaryNode {

	n := New(value)
	// use inherited method of treenode.Node.
	// requires passing back implemented member.
	// to tree.Node.SetRight.
	binaryNode.SetRight(n.Node)
	// return the BinaryNode.
	return n
}

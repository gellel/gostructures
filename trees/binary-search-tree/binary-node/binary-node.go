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

// New instantiates a BinaryNode.
func New(value int) *BinaryNode {
	return &BinaryNode{
		Parent: nil,
		Left:   nil,
		Right:  nil,
		Value:  value}
}

// Contains evaluates whether the BinaryNode
// with the target value exists in the connected BinaryNodes.
func (binaryNode *BinaryNode) Contains(value int) bool {
	return binaryNode.Get(value) != nil
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

// Get attempts to find the BinaryNode
// with the target value.
func (binaryNode *BinaryNode) Get(value int) *BinaryNode {
	if binaryNode.Value == value {
		return binaryNode
	}
	// check through BinaryNode.Left.
	if value < binaryNode.Value && binaryNode.HasLeft() {
		return binaryNode.Left.Get(value)
	}
	// check through BinaryNode.Right.
	if value > binaryNode.Value && binaryNode.HasRight() {
		return binaryNode.Right.Get(value)
	}
	return nil
}

// Maximum finds the deepest BinaryNode.Right.
func (binaryNode *BinaryNode) Maximum() *BinaryNode {
	if binaryNode.HasEmptyRight() {
		return binaryNode
	}
	return binaryNode.Right.Maximum()
}

// Minimum finds the deepest BinaryNode.Left.
func (binaryNode *BinaryNode) Minimum() *BinaryNode {
	if binaryNode.HasEmptyLeft() {
		return binaryNode
	}
	return binaryNode.Left.Minimum()
}

func (binaryNode *BinaryNode) Remove() {}

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

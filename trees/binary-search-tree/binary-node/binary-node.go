package binarynode

import (
	"fmt"

	treenode "github.com/gellel/gostructures/trees/tree-node"
)

// BinaryNode extends treenode.Node.
// Expresses a Binary Search Tree Node.
type BinaryNode struct {
	treenode.Node
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
func (node *BinaryNode) Contains(value int) bool {
	return node.Get(value) != nil
}

func (node *BinaryNode) Inherits() {
	fmt.Println(node.Node)
}

// Insert creates and assigns a Child BinaryNode based
// on the weight of the BinaryNode.Value.
// Core method of the Binary Search Tree.
func (node *BinaryNode) Insert(value int) *BinaryNode {

	// check that the new node should be left
	// balanced.
	if value < node.Value {
		// if we already have a BinaryNode.Left
		// move through the Left Search Tree.
		if node.HasLeft() {
			return node.Left.Insert(value)
		}
		return node.SetBinaryLeft(value)
	}
	// otherwise check that the new node
	// should be balanced right.
	if value > node.Value {
		// if we already have a BinaryNode.Right
		// move through the Right Search Tree.
		if node.HasRight() {
			return node.Right.Insert(value)
		}
		return node.SetBinaryRight(value)
	}
	return node
}

// Get attempts to find the BinaryNode
// with the target value.
func (node *BinaryNode) Get(value int) *BinaryNode {
	if node.Value == value {
		return node
	}
	// check through BinaryNode.Left.
	if value < node.Value && node.HasLeft() {
		return node.Left.Get(value)
	}
	// check through BinaryNode.Right.
	if value > node.Value && node.HasRight() {
		return node.Right.Get(value)
	}
	return nil
}

// Maximum finds the deepest BinaryNode.Right.
func (node *BinaryNode) Maximum() *BinaryNode {
	if node.HasEmptyRight() {
		return node
	}
	return node.Right.Maximum()
}

// Minimum finds the deepest BinaryNode.Left.
func (node *BinaryNode) Minimum() *BinaryNode {
	if node.HasEmptyLeft() {
		return node
	}
	return node.Left.Minimum()
}

// Remove BinaryNode based on its value.
func (node *BinaryNode) Remove(value int) bool {

	// attempt to find target.
	n := node.Get(value)
	// cannot find required BinaryNode.
	if n == nil {
		return false
	}

	if n.HasEmptyLeft() && n.HasEmptyRight() && n.HasParent() {
		n.Parent.RemoveChild(n.Value)

	} else if n.HasLeft() && n.HasRight() {
		minimum := n.Right.Minimum()
		if minimum.Value != n.Value {
			node.Remove(minimum.Value)
			return true
		}

	} else if n.HasLeft() {
		n.Parent.Left = n.Left
		return true

	} else if n.HasRight() {
		n.Parent.Right = n.Right
		return true
	}

	return false
}

// SetBinaryLeft creates BinaryNode and satisfies assignment
// to inherited Struct (treenode.Node) and sets parent dependencies.
func (node *BinaryNode) SetBinaryLeft(value int) *BinaryNode {

	n := New(value)
	// use inherited method of treenode.Node.
	// requires passing back implemented member.
	// to tree.Node.SetLeft.
	node.SetLeft(&n.Node)
	// return the BinaryNode.
	return n
}

// SetBinaryRight creates BinaryNode and satisfies assignment
// to inherited Struct (treenode.Node) and sets parent dependencies.
func (node *BinaryNode) SetBinaryRight(value int) *BinaryNode {

	n := New(value)
	// use inherited method of treenode.Node.
	// requires passing back implemented member.
	// to tree.Node.SetRight.
	node.SetRight(&n.Node)
	// return the BinaryNode.
	return n
}

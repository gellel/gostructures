package binarysearchtree

import (
	treenode "github.com/gellel/gostructures/trees/tree-node"
)

// BinarySearchTree Abstract Data Structure.
type BinarySearchTree struct {
	Root *treenode.Node
}

// New instantiates *BinarySearchTree.
func New(value float64) *BinarySearchTree {
	return &BinarySearchTree{
		Root: treenode.New(value)}
}

// Add inserts *treenode.Node and returns *BinarySearchTree.
func (binarySearchTree *BinarySearchTree) Add(value float64) *BinarySearchTree {
	binarySearchTree.Insert(value)
	return binarySearchTree
}

// Contains evaluates whether *BinarySearchTree.Root
// contains a *treenode.Node with the corresponding value.
func (binarySearchTree *BinarySearchTree) Contains(value float64) bool {
	return binarySearchTree.Root.Contains(value)
}

// Distribution computes the distribution of *BinarySearchTree.Nodes.
func (binarySearchTree *BinarySearchTree) Distribution() float64 {
	return binarySearchTree.Root.Distribution()
}

// Find searches for *treenode.Node within *BinarySearchTree.Root.
func (binarySearchTree *BinarySearchTree) Find(value float64) *treenode.Node {
	return binarySearchTree.Root.Find(value)
}

// GetLargest finds *treenode.Node with largest *treenode.Node.Value
// within *BinarySearchTree.Root.
func (binarySearchTree *BinarySearchTree) GetLargest() *treenode.Node {
	return binarySearchTree.Root.GetLargest()
}

// GetLargestValue finds *treenode.Node with largest *treenode.Node.Value
// within *BinarySearchTree.Root and returns its assigned value.
func (binarySearchTree *BinarySearchTree) GetLargestValue() float64 {
	return binarySearchTree.Root.GetLargestValue()
}

// GetSmallest finds *treenode.Node with smallest *treenode.Node.Value
// within *BinarySearchTree.Root.
func (binarySearchTree *BinarySearchTree) GetSmallest() *treenode.Node {
	return binarySearchTree.Root.GetSmallest()
}

// GetSmallestValue finds *treenode.Node with smallest *treenode.Node.Value
// within *BinarySearchTree.Root and returns its assigned value.
func (binarySearchTree *BinarySearchTree) GetSmallestValue() float64 {
	return binarySearchTree.Root.GetSmallestValue()
}

// Height computes the sum of *BinarySeachTree.Root's children.
func (binarySearchTree *BinarySearchTree) Height() float64 {
	return binarySearchTree.Root.Height()
}

// Insert performs core Binary Search Tree balancing operation.
func (binarySearchTree *BinarySearchTree) Insert(value float64) *treenode.Node {
	return binarySearchTree.Root.Insert(value)
}

// Remove removes *treenode.Node from *BinarySearchTree.Root.
func (binarySearchTree *BinarySearchTree) Remove(value float64) bool {
	return binarySearchTree.Root.Remove(value)
}

// ToSlice constructs a Slice containing all *treenode.Node's
// contained within *BinarySearchTree.Root.
func (binarySearchTree *BinarySearchTree) ToSlice() []*treenode.Node {
	return binarySearchTree.Root.ToSlice()
}

// ToSliceFloat64 constructs a Slice containing all *treenode.Node.Value's
// contained within *BinarySearchTree.Root.
func (binarySearchTree *BinarySearchTree) ToSliceFloat64() []float64 {
	return binarySearchTree.Root.ToSliceFloat64()
}

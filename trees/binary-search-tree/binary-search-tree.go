package binarysearchtree

import (
	binarynode "github.com/gellel/gostructures/trees/binary-search-tree/binary-node"
)

// BinarySearchTree Abstract Data Structure.
type BinarySearchTree struct {
	Root *binarynode.Node
}

// New instantiates *BinarySearchTree.
func New(value float64) *BinarySearchTree {
	return &BinarySearchTree{
		Root: binarynode.New(value)}
}

// Add inserts *binarynode.Node and returns *BinarySearchTree.
func (binarySearchTree *BinarySearchTree) Add(value float64) *BinarySearchTree {
	binarySearchTree.Insert(value)
	return binarySearchTree
}

// Balance computes the distribution of *BinarySearchTree.Nodes.
func (binarySearchTree *BinarySearchTree) Balance() float64 {
	return binarySearchTree.Root.Balance()
}

// Contains evaluates whether *BinarySearchTree.Root
// contains a *binarynode.Node with the corresponding value.
func (binarySearchTree *BinarySearchTree) Contains(value float64) bool {
	return binarySearchTree.Root.Contains(value)
}

// Find searches for *binarynode.Node within *BinarySearchTree.Root.
func (binarySearchTree *BinarySearchTree) Find(value float64) *binarynode.Node {
	return binarySearchTree.Root.Find(value)
}

// GetLargest finds *binarynode.Node with largest *binarynode.Node.Value
// within *BinarySearchTree.Root.
func (binarySearchTree *BinarySearchTree) GetLargest() *binarynode.Node {
	return binarySearchTree.Root.GetLargest()
}

// GetLargestValue finds *binarynode.Node with largest *binarynode.Node.Value
// within *BinarySearchTree.Root and returns its assigned value.
func (binarySearchTree *BinarySearchTree) GetLargestValue() float64 {
	return binarySearchTree.Root.GetLargestValue()
}

// GetSmallest finds *binarynode.Node with smallest *binarynode.Node.Value
// within *BinarySearchTree.Root.
func (binarySearchTree *BinarySearchTree) GetSmallest() *binarynode.Node {
	return binarySearchTree.Root.GetSmallest()
}

// GetSmallestValue finds *binarynode.Node with smallest *binarynode.Node.Value
// within *BinarySearchTree.Root and returns its assigned value.
func (binarySearchTree *BinarySearchTree) GetSmallestValue() float64 {
	return binarySearchTree.Root.GetSmallestValue()
}

// Height computes the sum of *BinarySeachTree.Root's children.
func (binarySearchTree *BinarySearchTree) Height() float64 {
	return binarySearchTree.Root.Height()
}

// Insert performs core Binary Search Tree balancing operation.
func (binarySearchTree *BinarySearchTree) Insert(value float64) *binarynode.Node {
	return binarySearchTree.Root.Insert(value)
}

// ToSlice constructs a Slice containing all *binarynode.Node's
// contained within *BinarySearchTree.Root.
func (binarySearchTree *BinarySearchTree) ToSlice() []*binarynode.Node {
	return binarySearchTree.Root.ToSlice()
}

// ToSliceFloat64 constructs a Slice containing all *binarynode.Node.Value's
// contained within *BinarySearchTree.Root.
func (binarySearchTree *BinarySearchTree) ToSliceFloat64() []float64 {
	return binarySearchTree.Root.ToSliceFloat64()
}

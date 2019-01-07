package avltree

import (
	binarysearchtree "github.com/gellel/gostructures/trees/binary-search-tree"
	treenode "github.com/gellel/gostructures/trees/tree-node"
)

// AVLTree Adelson-Velsky and Landis
// Search Tree Abstract Data Structure.
type AVLTree struct {
	*binarysearchtree.BinarySearchTree
}

// New instantiates *AVLTree.
func New(value float64) *AVLTree {
	return &AVLTree{
		binarysearchtree.New(value)}
}

// RotateLeft performs a Right weighted balance
// for AVLTree. Extracts middle *treenode.Node
// from connected -2 balance. Sets middle
// as root and shifts original root to middle's left.
func (avlTree *AVLTree) RotateLeft(node *treenode.Node) {

	// de-reference and store current node (*treenode.Node).
	root := *node
	// destroy root *treenode.Node.Right
	// to prevent recursion overflow.
	root.Right = nil
	// update memory pointer for
	// node argument and assign
	// *treenode.Node.Right new value.
	*node = *node.Right
	// set memory address for copied *treenode.Node.
	node.Left = &root
}

func (avlTree *AVLTree) RotateLeftRight(root *treenode.Node) {

}

func (avlTree *AVLTree) RotateRight(root *treenode.Node) {

}

func (avlTree *AVLTree) RotateRightLeft(root *treenode.Node) {

}

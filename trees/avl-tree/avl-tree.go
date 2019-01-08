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
// from *treenode.Node -2 distribution. Sets middle
// as root and shifts original root to middle's left.
func (avlTree *AVLTree) RotateLeft(node *treenode.Node) {
	// de-reference and store current node (*treenode.Node).
	root := *node
	// destroy root *treenode.Node.Right
	// to prevent recursion overflow.
	root.Right = nil
	// update memory pointer for
	// node argument and assign
	// *treenode.Node.Right as new value.
	*node = *node.Right
	// set memory address for copied *treenode.Node.
	node.Left = &root
}

// RotateLeftRight performs a *treenode.Node
// reassignment, setting *treenode.Node.Left as
// the second left-child of root. Shifts
// *treenode.Node.Left.Right as first left-child of
// accessed root node .
func (avlTree *AVLTree) RotateLeftRight(node *treenode.Node) {
	// de-reference and store current node's first-child.
	left := *node.Left
	// de-reference and store first-child's right node.
	leftRight := *left.Right
	// destroy *treenode.Node.Left's reference
	// to it's right-child (*treenode.Node.Left.Right)
	// to prevent recursion overflow.
	left.Right = nil
	// set root node's first left-child
	// to the stored node that existed at
	// *treenode.Node.Left.Right.
	node.Left = &leftRight
	// shift origin root node's first left-child
	// to become new *treenode.Node.Left's first
	// left-child.
	leftRight.Left = &left
	// rebalance the overweight root node.
	avlTree.RotateRight(node)
}

// RotateRight performs a Left weighted balance
// for AVLTree. Extracts middle *treenode.Node
// from *treenode.Node 2 distribution. Sets middle
// as root and shifts original root to middle's right.
func (avlTree *AVLTree) RotateRight(node *treenode.Node) {
	// de-reference and store current node (*treenode.Node).
	root := *node
	// destroy root *treenode.Node.Left
	// to prevent recursion overflow.
	root.Left = nil
	// update memory pointer for
	// node argument and assign
	// *treenode.Node.Left as new value.
	*node = *node.Left
	// set memory address for copied *treenode.Node.
	node.Right = &root

}

func (avlTree *AVLTree) RotateRightLeft(root *treenode.Node) {

}

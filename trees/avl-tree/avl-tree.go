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

func (avlTree *AVLTree) RotateLeft(root *treenode.Node) {

}

func (avlTree *AVLTree) RotateLeftRight(root *treenode.Node) {

}

func (avlTree *AVLTree) RotateRight(root *treenode.Node) {

}

func (avlTree *AVLTree) RotateRightLeft(root *treenode.Node) {

}

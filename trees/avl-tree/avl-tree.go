package avltree

import (
	binarysearchtree "github.com/gellel/gostructures/trees/binary-search-tree"
	binarynode "github.com/gellel/gostructures/trees/binary-search-tree/binary-node"
)

type AVLTree struct {
	*binarysearchtree.BinarySearchTree
}

func New(value float64) *AVLTree {
	return &AVLTree{
		binarysearchtree.New(value)}
}

func (avlTree *AVLTree) Balance(node *binarynode.Node) {
	// binarynode.Node has too many left
	// assigned binarynode.Node's.
	if node.Balance() > 1 {
		// if binarynode.Node.Right is overweight
		// its children need balance.
		if node.Left.Balance() > 0 {
			avlTree.RotateLeftLeft(node)
			// if binarynode.Node Left is overweight
			// its children need balance.
		} else if node.Left.Balance() < 0 {
			avlTree.RotateLeftRight(node)
		}
		// binarynode.Node has too many right
		// assigned binarynode.Node's.
	} else if node.Balance() < -1 {
		if node.Right.Balance() < 0 {
			avlTree.RotateRightRight(node)
		} else if node.Right.Balance() > 0 {
			avlTree.RotateRightLeft(node)
		}
	}
}

func (avlTree *AVLTree) Insert(value float64) {
	avlTree.BinarySearchTree.Insert(value)

	current := avlTree.BinarySearchTree.Find(value)

	for current != nil {
		avlTree.Balance(current)
		current = current.Parent
	}
}

package avltree

import (
	"fmt"
	"log"

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

// Balance performs core AVL Tree operation.
func (avlTree *AVLTree) Balance(node *treenode.Node) {

	//fmt.Println(node.Value, node.Distribution(), node.HasLeft(), node.HasRight(), node.HasParent())

	if node.Distribution() > 1 {
		if node.Right.Distribution() > 1 {
			avlTree.RotateLeftRight(node.Left)
		} else {
			avlTree.RotateLeft(node)
		}
	} else if node.Distribution() < -1 {
		if node.Left.Distribution() < -1 {
			avlTree.RotateRightLeft(node)
		} else {
			avlTree.RotateRight(node)
		}
	}
}

// Insert performs core Binary Search Tree insert operation,
// but attempts to balance *AVLTree using AVL method.
func (avlTree *AVLTree) Insert(value float64) *AVLTree {

	node := avlTree.BinarySearchTree.Insert(value)

	if node.HasLeft() && node.Left.HasValue(value) {
		node = node.Left
	} else if node.HasRight() && node.Right.HasValue(value) {
		node = node.Right
	}

	c := 0
	for node.HasParent() {

		fmt.Println(node.Left, node.Value, node.Right)

		avlTree.Balance(node)

		fmt.Println(node.Left, node.Value, node.Right)
		c++

		if c > 10 {
			log.Panicf("")
		}
		node = node.Parent
	}
	return avlTree
}

// Remove removes a *treenode.Node based on value.
// Rebalances AVLTree if successful.
func (avlTree *AVLTree) Remove(value float64) *AVLTree {
	if avlTree.BinarySearchTree.Remove(value) && avlTree.Root != nil {
		avlTree.Balance(avlTree.Root)
	}
	return avlTree
}

// RotateLeft performs a *treenode.Node
// child node re-alignment. Sets *treenode.Node.Right.
// as new *treenode.Node. Sets original *treenode.Node as
// new *treenode.Node.Left.
func (avlTree *AVLTree) RotateLeft(node *treenode.Node) {

	right := node.Right

	node.Right = nil

	root := *node

	*node = *right

	node.Left = &root

	// from: [root  1]-->[root.right  2]-->[root.right.right  3]
	// to:   [root.left  1]<--[root  2]-->[root.right   3]
}

// RotateLeftRight performs a *treenode.Node
// child node re-alignment. Sets *treenode.Node.Left.Right as
// the child node of *treenode.Node.Left.
// After all *treenode.Node children are left-aligned
// the root node is rebalanced using RotateLeft.
func (avlTree *AVLTree) RotateLeftRight(node *treenode.Node) {
	// de-reference and store current root.Left.
	left := *node.Left
	// de-reference and store current root.Left.Right.
	leftRight := *left.Right
	// destroy root.Left.Right to prevent
	// recursion overflow due to circular reference.
	left.Right = nil
	// shift root.Left.Right to become root.Left.
	node.Left = &leftRight
	// shift root.Left.Right to become root.Left.Left.
	leftRight.Left = &left

	// from: [root  3]-->[root.left  1]-->[root.left.right  2]
	// to:   [root  3]-->[root.left  2]-->[root.left.left   1]

	// rebalance the overweight root node.
	avlTree.RotateRight(node)
}

// RotateRight performs a *treenode.Node
// child node re-alignment. Sets *treenode.Node.Left.
// as new *treenode.Node. Sets original *treenode.Node as
// new *treenode.Node.Right.
func (avlTree *AVLTree) RotateRight(node *treenode.Node) {

	left := node.Left

	node.Left = nil

	root := *node

	*node = *left

	node.Right = &root

	// from: [root  3]-->[root.left  2]-->[root.left.left  1]
	// to:   [root.left  1]<--[root  2]-->[root.right   3]
}

// RotateRightLeft performs a *treenode.Node
// child node re-alignment. Sets *treenode.Node.Right.Left as
// the child node of *treenode.Node.Right.
// After all *treenode.Node children are right-aligned
// the root node is rebalanced using RotateLeft.
func (avlTree *AVLTree) RotateRightLeft(node *treenode.Node) {
	// de-reference and store current root.Right.
	right := *node.Right
	// de-reference and store current root.Right.Left.
	rightLeft := *right.Left
	// destroy root.Right.Left to prevent
	// recursion overflow due to circular reference.
	right.Left = nil
	// shift root.Right.Left to become root.Right.
	node.Right = &rightLeft
	// shift root.Right.Left to become root.Right.Right.
	rightLeft.Right = &right

	// from: [root  1]-->[root.right  3]-->[root.right.left   2]
	// to:   [root  1]-->[root.right  2]-->[root.right.right  3]

	// rebalance right-weighted root node.
	avlTree.RotateLeft(node)
}

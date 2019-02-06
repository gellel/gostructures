// Package redblack exports a RedBlack-Search-Tree.
// A RedBlack-Search-Tree is abstract data structure that
// manages the insertion of floating point values so that they
// can be accessed and arranged in guaranteed O(log(n)) time.
// Each RedBlack-Search-Tree holds a collection of RedBlack-Tree-Nodes
// that manages the insertion and ordering of new nodes relative
// to the accessed RedBlack-Tree-Node.
// In the context of the RedBlack-Search-Tree, this action takes place
// relative to the root node. Package also exports a
// function "New" that simplifies instantiating
// RedBlack-Search-Tree pointers.
package redblack

import (
	node "github.com/gellel/gostructures/abstracts/trees/red-black/red-black-node"
)

// RedBlack declares the implementation for a RedBlack-Search-Tree.
// Methods in the RedBlack-Search-Tree leverage the assigned root
// RedBlack-Tree-Node.
type rb interface {
	Find(value float64) *node.RedBlack
	Insert(value float64) *RedBlack
	Remove(value float64) *RedBlack
	ToRedBlackSlice() []*node.RedBlack
	ToFloatSlice() []float64
	Walk() *RedBlack
}

// RedBlack delcares the struct for a RedBlack-Search-Tree. Each RedBlack-Search-Tree
// contains a root RedBlack-Tree-Node that acts as the insertion point for all
// subsequent floating point values. Methods provided by RedBlack-Search-Tree leverage
// the methods of the assigned RedBlack-Tree-Node at the root position.
type RedBlack struct {
	Root *node.RedBlack // Root is the heighest RedBlack-Tree-Node stored within the RedBlack-Search-Tree. Each RedBlack-Seach-Tree must contain a root node to perform ordered operations.
}

// New instantiates a new RedBlack-Search-Tree pointer.
func New(value float64) *RedBlack {
	return &RedBlack{Root: node.New(value).AssignSide(node.ROOT)}
}

// Find checks whether the RedBlack-Search-Tree contains the argument value.
func (redBlack *RedBlack) Find(value float64) *node.RedBlack {
	return redBlack.Root.Find(value)
}

// Insert creates and assigns a new RedBlack-Tree-Node to the accessed RedBlack-Search-Tree.
// When an argument value is smaller than the accessed RedBlack-Tree-Node value, the new instance
// is stored on the accessed RedBlack-Tree-Node's left position. Alternatively, when
// a larger value is provided, it is stored than the right position. Equal sums are
// discarded and no new RedBlack-Tree-Node is created. When an insertion is performed,
// accessed RedBlack-Tree-Node checks that the tree does not violate red black conditions.
func (redBlack *RedBlack) Insert(value float64) *RedBlack {
	redBlack.Root.Insert(value)
	return redBlack
}

// Remove deletes a RedBlack-Tree-Node contained within the RedBlack-Search-Tree.
func (redBlack *RedBlack) Remove(value float64) *RedBlack {
	redBlack.Root.Remove(value)
	return redBlack
}

// ToRedBlackSlice creates an ordered slice of the assigned RedBlack-Tree-Node's child-nodes.
func (redBlack *RedBlack) ToRedBlackSlice() []*node.RedBlack {
	return redBlack.Root.ToRedBlackSlice()
}

// ToFloatSlice creates an ordered slice of the assigned RedBlack-Tree-Node's child-nodes values.
func (redBlack *RedBlack) ToFloatSlice() []float64 {
	return redBlack.Root.ToFloatSlice()
}

// Walk accesses all assigned child-nodes contained within the RedBlack-Search-Tree.
func (redBlack *RedBlack) Walk() *RedBlack {
	redBlack.Root.Walk()
	return redBlack
}

var _ rb = (*RedBlack)(nil)

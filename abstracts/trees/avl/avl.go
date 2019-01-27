// Package binary exports a AVL-Search-Tree.
// A AVL-Search-Tree is abstract data structure that
// manages the insertion of floating point values so that they
// can be accessed and arranged in guaranteed O(log(n)) time.
// Each AVL-Search-Tree holds a collection of AVL-Tree-Nodes
// that manages the insertion and ordering of new nodes relative
// to the accessed AVL-Tree-Node.
// In the context of the AVL-Search-Tree, this action takes place
// relative to the root node. Package also exports a
// function "New" that simplifies instantiating
// AVL-Search-Tree pointers.
package avl

import (
	node "github.com/gellel/gostructures/abstracts/trees/avl/avl-node"
)

// AVL declares the implementation for a AVL-Search-Tree.
// Methods in the AVL-Search-Tree leverage the assigned root
// AVL-Tree-Node.
type avl interface {
	Balance() int
	Find(value float64) *node.AVL
	Height() int
	Insert(value float64) *AVL
	Remove(value float64) *AVL
	ToAVLSlice() []*node.AVL
	ToFloatSlice() []float64
	Walk() *AVL
}

// AVL delcares the struct for a AVL-Search-Tree. Each AVL-Search-Tree
// contains a root AVL-Tree-Node that acts as the insertion point for all
// subsequent floating point values. Methods provided by AVL-Search-Tree leverage
// the methods of the assigned AVL-Tree-Node at the root position.
type AVL struct {
	Root *node.AVL // Root is the heighest AVL-Tree-Node stored within the AVL-Search-Tree. Each AVL-Seach-Tree must contain a root node to perform ordered operations.
}

// New instantiates a new AVL-Search-Tree pointer.
func New(value float64) *AVL {
	return &AVL{Root: node.New(value).UnsafelyAssignSide("ROOT")}
}

// Balance computes the weighting of the AVL-Search-Tree.
func (avl *AVL) Balance() int {
	return avl.Root.Balance()
}

// Find checks whether the AVL-Search-Tree contains the argument value.
func (avl *AVL) Find(value float64) *node.AVL {
	return avl.Root.Find(value)
}

// Height computes the current depth of AVL-Search-Tree.
func (avl *AVL) Height() int {
	return avl.Root.Height()
}

// Insert creates and assigns a new AVL-Tree-Node to the accessed AVL-Search-Tree.
// When an argument value is smaller than the accessed AVL-Tree-Node value, the new instance
// is stored on the accessed AVL-Tree-Node's left position. Alternatively, when
// a larger value is provided, it is stored than the right position. Equal sums are
// discarded and no new AVL-Tree-Node is created. When an insertion is performed,
// accessed AVL-Tree-Node checks that the it is evenly distributed. When bias left,
// the child nodes at the left position are shuffled one over so that they are weighted right
// evenly. The alternative case moves nodes from the right to the left.
func (avl *AVL) Insert(value float64) *AVL {
	avl.Root.Insert(value)
	return avl
}

// Remove deletes a AVL-Tree-Node contained within the AVL-Search-Tree.
func (avl *AVL) Remove(value float64) *AVL {
	avl.Root.Remove(value)
	return avl
}

// ToAVLSlice creates an ordered slice of the assigned AVL-Tree-Node's child-nodes.
func (avl *AVL) ToAVLSlice() []*node.AVL {
	return avl.Root.ToAVLSlice()
}

// ToFloatSlice creates an ordered slice of the assigned AVL-Tree-Node's child-nodes values.
func (avl *AVL) ToFloatSlice() []float64 {
	return avl.Root.ToFloatSlice()
}

// Walk accesses all assigned child-nodes contained within the AVL-Search-Tree.
func (avl *AVL) Walk() *AVL {
	avl.Root.Walk()
	return avl
}

var _ avl = (*AVL)(nil)

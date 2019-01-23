// Package binary exports a Binary-Search-Tree.
// A Binary-Search-Tree is abstract data structure that
// manages the insertion of floating point values so that they
// can be accessed and arranged O(log(n)) time.
// Each Binary-Search-Tree holds a collection of Binary-Tree-Nodes
// that manages the insertion and ordering of new nodes relative
// to the accessed Binary-Tree-Node.
// In the context of the Binary-Search-Tree, this action takes place
// relative to the root node. Package also exports a
// function "New" that simplifies instantiating
// Binary-Search-Tree pointers.
package binary

import node "github.com/gellel/gostructures/abstracts/trees/binary/binary-node"

// Binary declares the implementation for a Binary-Search-Tree.
// Methods in the Binary-Search-Tree leverage the assigned root
// Binary-Tree-Node.
type binary interface {
	Find(value float64) *node.Binary
	Height() int
	Insert(value float64) *Binary
	Remove(value float64) *Binary
	ToBinarySlice() []*node.Binary
	ToFloatSlice() []float64
	Walk()
}

// Binary delcares the struct for a Binary-Search-Tree. Each Binary-Search-Tree
// contains a root Binary-Tree-Node that acts as the insertion point for all
// subsequent floating point values. Methods provided by Binary-Search-Tree leverage
// the methods of the assigned Binary-Tree-Node at the root position.
type Binary struct {
	// Root is the heighest Binary-Tree-Node stored within the Binary-Search-Tree.
	// Each Binary-Seach-Tree must contain a root node to perform ordered operations.
	Root *node.Binary
}

// New instantiates a new Binary-Search-Tree pointer.
func New(value float64) *Binary {
	return &Binary{Root: node.New(value)}
}

// Find checks whether the Binary-Search-Tree contains the argument value.
func (binary *Binary) Find(value float64) *node.Binary {
	return binary.Root.Find(value)
}

// Height computes the current depth of Binary-Search-Tree.
func (binary *Binary) Height() int {
	return binary.Root.Height()
}

// Insert creates and assigns a new Binary-Tree-Node to the Binary-Search-Tree.
// When an argument value is smaller than the root Binary-Tree-Node value, the new instance
// is stored on the left position. Alternatively, when a larger value is provided,
// it is stored than the right position. Equal sums are
// discarded and no new Binary-Tree-Node is created. When the root node contains
// a left or right node where the value would normally sit, the aforementioned
// method is performed for n-number of child nodes until a suitable position is found.
func (binary *Binary) Insert(value float64) *Binary {
	binary.Root.Insert(value)
	return binary
}

// Remove deletes a Binary-Tree-Node contained within the Binary-Search-Tree.
func (binary *Binary) Remove(value float64) *Binary {
	binary.Root.Remove(value)
	return binary
}

// ToBinarySlice creates an ordered slice of the assigned Binary-Tree-Node's child-nodes.
func (binary *Binary) ToBinarySlice() []*node.Binary {
	return binary.Root.ToBinarySlice()
}

// ToFloatSlice creates an ordered slice of the assigned Binary-Tree-Node's child-nodes values.
func (binary *Binary) ToFloatSlice() []float64 {
	return binary.Root.ToFloatSlice()
}

// Walk accesses all assigned child-nodes contained within the Binary-Search-Tree.
func (binary *Binary) Walk() {
	binary.Root.Walk()
}

var _ binary = (*Binary)(nil)

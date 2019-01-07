package treenode

import (
	"math"
)

// A Node contained within a Search Tree.
type Node struct {
	Parent *Node   // *Node hosting *Node.
	Left   *Node   // *Node hosted by *Node.
	Right  *Node   // *Node hosted by *Node.
	Value  float64 // Sum of *Node.
}

// New instantiates *Node.
func New(value float64) *Node {
	return &Node{
		Parent: nil,
		Left:   nil,
		Right:  nil,
		Value:  value}
}

// Contains evaluates whether a *Node with
// corresponding *Node.Value.
func (node *Node) Contains(value float64) bool {
	return node.Find(value) != nil
}

// Distribution collects the sum of *Node's children
// and computes the weighting of *Node.Left
// and *Node.Right. Negative value is
// right weighted. Positive value is left weighted.
func (node *Node) Distribution() float64 {
	return (node.HeightLeft() - node.HeightRight())
}

// Find searches *Node children for
// *Node.Value.
func (node *Node) Find(value float64) *Node {
	if node.HasValue(value) {
		return node
	}
	if node.HasLeft() {
		return node.Left.Find(value)
	}
	if node.HasRight() {
		return node.Right.Find(value)
	}
	return nil
}

// GetLargest moves through *Node.Right
// until *Node.Right has no assigned *Node.Right.
func (node *Node) GetLargest() *Node {
	if node.HasRight() {
		return node.Right.GetLargest()
	}
	return node
}

// GetLargestValue gets *Node.Value from *Node.GetLargest.
func (node *Node) GetLargestValue() float64 {
	return node.GetLargest().Value
}

// GetSmallest moves through *Node.Left
// until *Node.Left has no assigned *Node.Left.
func (node *Node) GetSmallest() *Node {
	if node.HasLeft() {
		return node.Left.GetSmallest()
	}
	return node
}

// GetSmallestValue gets *Node.Value from *Node.GetSmallest.
func (node *Node) GetSmallestValue() float64 {
	return node.GetSmallest().Value
}

// HasEmptyLeft checks *Node.Left for nil.
func (node *Node) HasEmptyLeft() bool {
	return node.Left == nil
}

// HasEmptyParent checks *Node.Parent for nil.
func (node *Node) HasEmptyParent() bool {
	return node.Parent == nil
}

// HasEmptyRight checks *Node.Right for nil.
func (node *Node) HasEmptyRight() bool {
	return node.Right == nil
}

// HasLeft checks *Node.Left is *Node.
func (node *Node) HasLeft() bool {
	return !node.HasEmptyLeft()
}

// HasParent checks *Node.Parent is *Node.
func (node *Node) HasParent() bool {
	return !node.HasEmptyParent()
}

// HasRight checks *Node.Right is *Node.
func (node *Node) HasRight() bool {
	return !node.HasEmptyRight()
}

// HasValue evaluates whether passed
// float64 value is *Node.Value.
func (node *Node) HasValue(value float64) bool {
	return node.Value == value
}

// Height computes sum of *Node children
// selecting the max of *Node.Left and *Node.Right.
func (node *Node) Height() float64 {
	return math.Max(node.HeightLeft(), node.HeightRight())
}

// HeightLeft computes height of *Node.Left.
func (node *Node) HeightLeft() float64 {
	if node.HasEmptyLeft() {
		return 0.0
	}
	return (node.Left.Height() + 1.0)
}

// HeightRight computes height of *Node.Right.
func (node *Node) HeightRight() float64 {
	if node.HasEmptyRight() {
		return 0.0
	}
	return (node.Right.Height() + 1.0)
}

// Insert performs core Binary Search Tree
// assignment operation. Smaller values
// are stored on *Node.Left. Larger values
// are stored on *Node.Right.
func (node *Node) Insert(value float64) *Node {
	if value < node.Value {
		if node.HasLeft() {
			return node.Left.Insert(value)
		}
		return node.SetLeft(value)
	}
	if value > node.Value {
		if node.HasRight() {
			return node.Right.Insert(value)
		}
		return node.SetRight(value)
	}
	return node
}

// OverweightLeft determines if *Node
// contains more left children.
func (node *Node) OverweightLeft() bool {
	return node.Distribution() > 0
}

// OverweightRight determines if *Node
// contains more right children.
func (node *Node) OverweightRight() bool {
	return node.Distribution() < 0
}

// Remove removes a *Node connection.
func (node *Node) Remove(value float64) bool {
	if node.HasValue(value) {
		if node.HasParent() {
			return node.Parent.RemoveByValue(value)
		}
		return node.RemoveByValue(value)
	} else if value < node.Value && node.HasLeft() {
		return node.Left.Remove(value)
	} else if value > node.Value && node.HasRight() {
		return node.Right.Remove(value)
	}
	return false
}

// RemoveByValue removes a *Node connection by *Node.Value.
func (node *Node) RemoveByValue(value float64) bool {
	if node.HasLeft() && node.Left.HasValue(value) {
		node.RemoveLeft()
		return true
	}
	if node.HasRight() && node.Right.HasValue(value) {
		node.RemoveRight()
		return true
	}
	return false
}

// RemoveLeft unassigns *Node.Left.
func (node *Node) RemoveLeft() *Node {
	node.Left = nil
	return node
}

// RemoveRight unassigns *Node.Right.
func (node *Node) RemoveRight() *Node {
	node.Right = nil
	return node
}

// Same compares whether the provided *Node
// is the same reference.
func (node *Node) Same(n *Node) bool {
	return node == n
}

// SetLeft assigns *Node.Left a new *Node.
func (node *Node) SetLeft(value float64) *Node {
	node.Left = New(value)
	node.Left.Parent = node
	return node
}

// SetLeftNode graphs an existing *Node
// to the accessed *Node.
func (node *Node) SetLeftNode(n *Node) *Node {
	node.Left = n
	return node
}

// SetRight assigns *Node.Right a new *Node.
func (node *Node) SetRight(value float64) *Node {
	node.Right = New(value)
	node.Right.Parent = node
	return node
}

// SetRightNode graphs an existing *Node
// to the accessed *Node.
func (node *Node) SetRightNode(n *Node) *Node {
	node.Right = n
	return node
}

// SetValue mutates *Node.Value and reassigns *Node.
func (node *Node) SetValue(value float64) *Node {
	if node.HasValue(value) {
		return node
	}
	node.Value = value
	if node.HasParent() {
		node.Parent = node.Parent.GetLargest()
	}
	return node
}

// ToSlice arranges *Node children into
// and ordered Slice of *Node.
func (node *Node) ToSlice() []*Node {
	nodes := make([]*Node, 0)
	if node.HasLeft() {
		nodes = append(nodes, node.Left.ToSlice()...)
	}
	nodes = append(nodes, node)
	if node.HasRight() {
		nodes = append(nodes, node.Right.ToSlice()...)
	}
	return nodes
}

// ToSliceFloat64 arranges *Node children into
// and ordered Slice of *Node.Value.
func (node *Node) ToSliceFloat64() []float64 {
	float64s := make([]float64, 0)
	if node.HasLeft() {
		float64s = append(float64s, node.Left.ToSliceFloat64()...)
	}
	float64s = append(float64s, node.Value)
	if node.HasRight() {
		float64s = append(float64s, node.Right.ToSliceFloat64()...)
	}
	return float64s
}

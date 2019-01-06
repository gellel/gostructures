package binarynode

import "math"

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

// Balance collects the sum of *Node's children
// and computes the distribution of *Node.Left
// and *Node.Right.
func (node *Node) Balance() float64 {
	return (node.HeightLeft() - node.HeightRight())
}

// Contains evaluates whether a *Node with
// corresponding *Node.Value.
func (node *Node) Contains(value float64) bool {
	return node.Find(value) != nil
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

// GetLargest gets *Node.Value from *Node.GetLargest.
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

// GetSmallest gets *Node.Value from *Node.GetSmallest.
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

// SetLeft assigns *Node.Left a new *Node.
func (node *Node) SetLeft(value float64) *Node {
	node.Left = New(value)
	node.Left.Parent = node
	return node
}

// SetRight assigns *Node.Right a new *Node.
func (node *Node) SetRight(value float64) *Node {
	node.Right = New(value)
	node.Right.Parent = node
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

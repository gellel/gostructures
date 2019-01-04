package treenode

// Node structure for Search Tree.
type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	Value  interface{}
}

// New instantiates Node.
func New(value interface{}) *Node {
	return &Node{
		Parent: nil,
		Left:   nil,
		Right:  nil,
		Value:  value}
}

// HasEmptyLeft checks if Node.Left is nil.
func (node *Node) HasEmptyLeft() bool {
	return node.Left == nil
}

// HasEmptyRight checks if Node.Left is nil.
func (node *Node) HasEmptyRight() bool {
	return node.Right == nil
}

// HasLeft checks if Node has Node.Left.
func (node *Node) HasLeft() bool {
	return node.Left != nil
}

// HasParent checks if Node has Node.Parent.
func (node *Node) HasParent() bool {
	return node.Parent != nil
}

// HasRight checks if Node has Node.Right.
func (node *Node) HasRight() bool {
	return node.Right != nil
}

// HasValue checks if provided value is Node.Value.
func (node *Node) HasValue(value interface{}) bool {
	return node.Value == value
}

// HasLeftValue checks if provided value
// is stored in Node.Left.Value.
func (node *Node) HasLeftValue(value interface{}) bool {
	return node.HasLeft() && node.Left.HasValue(value)
}

// HasRightValue checks if provided value
// is stored in Node.Right.Value.
func (node *Node) HasRightValue(value interface{}) bool {
	return node.HasRight() && node.Right.HasValue(value)
}

// Height returns max of connected Nodes.
func (node *Node) Height() int {
	// get sum of Node.Left connected Nodes.
	l := node.HeightLeft()
	// get sum of Node.Right connected Nodes.
	r := node.HeightRight()
	// find largest.
	if l > r {
		return l
	}
	return r
}

// HeightLeft returns sum of connected
// Nodes in Node.Left.
func (node *Node) HeightLeft() int {
	if node.HasEmptyLeft() {
		// if there is no Node.Left
		// height is therefore zero.
		return 0
	}
	// height is sum of Node.Left plus current Node.
	return node.Left.Height() + 1
}

// HeightRight returns sum of connected
// Nodes in Node.Right.
func (node *Node) HeightRight() int {
	if node.HasEmptyRight() {
		// if there is no Node.Right
		// height is therefore zero.
		return 0
	}
	// height is sum of Node.Right plus current Node.
	return node.Right.Height() + 1
}

// RemoveChild removes Node based on contained value.
// Removes Node.Left before Node.Right.
func (node *Node) RemoveChild(value interface{}) bool {
	if node.HasLeftValue(value) {
		node.RemoveLeft()
		return true
	}
	if node.HasRightValue(value) {
		node.RemoveRight()
		return true
	}
	return false
}

// RemoveLeft removes Node.Left.
func (node *Node) RemoveLeft() *Node {
	node.Left = nil
	// return accessed Node.
	return node
}

// RemoveRight unsets Node.Right.
func (node *Node) RemoveRight() *Node {
	node.Right = nil
	// return accessed Node.
	return node
}

// SetLeft reassigns as Node to the reference Node.
// Current Node.Left is collected by garbage collector if orphaned.
func (node *Node) SetLeft(n *Node) *Node {
	if node.HasLeft() {
		// remove Node.Left.Parent as
		// the current Node is to become the
		// provided Node's parent.
		node.Left.Parent = nil
	}
	// set Node.Left to provided Node.
	node.Left = n
	// update Node.Left.Parent to Node.
	node.Left.Parent = node
	// return accessed Node.
	return node
}

// SetRight reassigns as Node to the reference Node.
// Current Node.Right is collected by garbage collector if orphaned.
func (node *Node) SetRight(n *Node) *Node {
	if node.HasRight() {
		// remove Node.Right.Parent as
		// the current Node is to become the
		// provided Node's parent.
		node.Right.Parent = nil
	}
	// set Node.Right to provided Node.
	node.Right = n
	// update Node.Right.Parent to Node.
	node.Right.Parent = node
	// returned accessed Node.
	return node
}

// SetValue changes the assigned value to Node.Value.
func (node *Node) SetValue(value interface{}) *Node {
	node.Value = value
	// returned accessed Node.
	return node
}

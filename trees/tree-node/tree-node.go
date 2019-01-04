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

// SetLeft reassigns as Node to the reference Node.
// Current Node.Left is collected by garbage collector if orphaned.
func (node *Node) SetLeft(n *Node) *Node {
	if !(node.Left == nil) {
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
	if !(node.Right == nil) {
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
	return node
}

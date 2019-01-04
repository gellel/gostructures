package treenode

type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	Value  interface{}
}

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
	// update Node.Left.Parent to Node.
	return node.Set(node.Left, n).Set(node.Left.Parent, node)
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
	// update Node.Right.Parent to Node.
	return node.Set(node.Right, n).Set(node.Right.Parent, node)
}

func (node *Node) Set(side *Node, n *Node) *Node {
	side = n
	return side
}

func (node *Node) SetValue(value interface{}) *Node {
	node.Value = value
	return node
}

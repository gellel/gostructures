// Package node contains Single Linked-List node.
package node

// Node is a single LinkedList Node.
type Node struct {
    Next *Node
	Value interface{}
}

// New instantiates a new Single Linked-List Node.
func New(property interface{}) *Node {
    return &Node{Value: property}
}
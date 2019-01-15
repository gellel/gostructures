// Package node exports Double Linked-List Node. A Double Linked-List node
// is a connection within a Double Linked-List. Double Linked-List node
// can access previous and next siblings. Package also exports a node
// pointer instantiation function.
package node

// Node defines a connection within the Double Linked-List. Unlike
// a node within a Single Linked-List the Double Linked-List can
// access its previous connection. Requires more memory.
type Node struct {
	Previous *Node       // Previous connection within Double Linked-List node.
	Next     *Node       // Next connection within Double Linked-List node.
	Value    interface{} // Property of any type assigned to the Double Linked-List node.
}

// New instantiates a new Double Linked-List Node. Value
// can be of any supported of the supported Go interfaces.
func New(property interface{}) *Node {
	return &Node{Value: property}
}

// Package double exports a Double Linked-List structure. Double Linked-List
// provides the insertion, searching and deletion methods while offering
// some extensions to the base Double Linked-List interface. Package also
// exports a Double Linked-List pointer instantiation function.
package double

import node "github.com/gellel/gostructures/abstracts/linked-lists/double/double-node"

// LinkedList defines the Double Linked-List data structure.
// Double Linked-List holds a sequence of Double Linked-List nodes.
// Each connection can be traversed from one node to another provided
// that they have a non Nil previous or next. Double Linked-List can
// contain any mixture of data.
type LinkedList struct {
	Head *node.Node // First Double Linked-List node.
	Tail *node.Node // Last Double Linked-List node.
}

func (double *LinkedList) Append() *LinkedList {
	return double
}

func (double *LinkedList) InsertAfter(node *node.Node, property interface{}) *LinkedList {

}

func (double *LinkedList) InsertBefore(node *node.Node, property interface{}) *LinkedList {

}

func (double *LinkedList) Find() *node.Node {

}

func (double *LinkedList) Prepend() *LinkedList {
	return double
}

// Package double exports a Double Linked-List structure. Double Linked-List
// provides the insertion, searching and deletion methods while offering
// some extensions to the base Double Linked-List interface. Package also
// exports a Double Linked-List pointer instantiation function.
package double

// DoubleLinkedList defines the Double Linked-List data structure.
// Double Linked-List holds a sequence of Double Linked-List nodes.
// Each connection can be traversed from one node to another provided
// that they have a non Nil previous or next. Double Linked-List can
// contain any mixture of data.
type DoubleLinkedList struct {
}

func (doubleLinkedList *DoubleLinkedList) Add() *DoubleLinkedList {
	return doubleLinkedList
}

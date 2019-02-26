// Package single exports a Single Linked-List structure. Single Linked-List
// provides the insertion, searching and deletion methods while offering
// some extensions to the standard Single Linked-List. Package also
// exports a Single Linked-List pointer instantiation function.
package single

import (
	"fmt"

	node "github.com/gellel/gostructures/abstracts/linked-lists/single/single-node"
)

// LinkedList declares the interface for a singly Linked-List.
type linkedList interface {
	Append(value interface{}) *LinkedList
	Find(value interface{}) *node.Single
	FindAll(value interface{}) []*node.Single
	HasHead() bool
	HasTail() bool
	InsertAfter(s *node.Single, value interface{}) *LinkedList
	InsertBefore(s *node.Single, value interface{}) *LinkedList
	IsEmpty() bool
	IsPopulated() bool
	Prepend(value interface{}) *LinkedList
	Remove(value interface{}) bool
	Set(s *node.Single) *LinkedList
	SetHead(s *node.Single) *LinkedList
	SetTail(s *node.Single) *LinkedList
	Shift() *LinkedList
	ShiftGet() interface{}
	Size() int
	Walk()
}

// LinkedList declares the struct for a Single Linked-List.
// Single Linked-List holds a sequence of Single Linked-List-Nodes.
// Each connection can be traversed from one node to another provided
// that they have an adjacent Single Linked-List-Node.
// Single Linked-List can contain any mixture of data.
type LinkedList struct {
	Head *node.Single
	Tail *node.Single
}

// New instantiates a new singly Linked-List.
func New() *LinkedList {
	return &LinkedList{}
}

// Append adds a new single Linked-List-Node to the end of accessed Linked-List.
func (single *LinkedList) Append(value interface{}) *LinkedList {
	if single.IsEmpty() {
		return single.Set(node.New(value))
	}
	return single.SetTail(single.Tail.AddNext(value, true))
}

// Find searches for values stored in the Linked-List.
func (single *LinkedList) Find(value interface{}) *node.Single {
	n := single.Head
	for n != nil {
		if n.HasValue(value) {
			return n
		}
		n = n.Next
	}
	return nil
}

// FindAll searches for all values stored in the Linked-List.
func (single *LinkedList) FindAll(value interface{}) []*node.Single {
	s := make([]*node.Single, 0)
	n := single.Head
	for n != nil {
		if n.HasValue(value) {
			s = append(s, n)
		}
		n = n.Next
	}
	return s
}

// HasHead checks whether the Linked-List has a Head.
func (single *LinkedList) HasHead() bool {
	return single.Head != nil
}

// HasTail checks whether the Linked-List has a Tail.
func (single *LinkedList) HasTail() bool {
	return single.Tail != nil
}

// InsertAfter adds a new Single-Linked-List-Node after the provided Single-Linked-List-Node.
func (single *LinkedList) InsertAfter(s *node.Single, value interface{}) *LinkedList {
	if single.IsEmpty() {
		single.SetHead(s.AddNext(value, false)).SetTail(s.Next)
	} else if s.HasNext() {
		s.AssignNext(node.New(value).AssignNext(s.Next, false), false)
	} else {
		single.Append(value)
	}
	return single
}

// InsertBefore adds a new Single-Linked-List-Node before the provided Single-Linked-List-Node.
func (single *LinkedList) InsertBefore(s *node.Single, value interface{}) *LinkedList {
	if single.IsEmpty() {
		single.SetHead(s.AddNext(value, false)).SetTail(s.Next)
	} else if s.HasNext() {
		n := node.New(s.Value).AssignNext(s.Next, false)
		s.AssignValue(value).AssignNext(n, false)
	} else {
		single.Prepend(value)
	}
	return single
}

// IsEmpty checks whether the Linked-List is empty.
func (single *LinkedList) IsEmpty() bool {
	return single.Head == nil && single.Tail == nil
}

// IsPopulated checks whether the Linked-List contains both a Head and Tail Single-Linked-List-Node.
func (single *LinkedList) IsPopulated() bool {
	return single.Head != nil && single.Tail != nil
}

// Prepend adds a Single-Linked-List-Node to beginning of the Linked-List.
func (single *LinkedList) Prepend(value interface{}) *LinkedList {
	if single.IsEmpty() {
		return single.Set(node.New(value))
	}
	return single.SetHead(node.New(value).AssignNext(single.Head, false))
}

// Remove deletes a Single-Linked-List-Node from the Linked-List by value.
func (single *LinkedList) Remove(value interface{}) bool {
	if single.IsEmpty() {
		return false
	}
	for single.HasHead() && single.Head.HasValue(value) {
		single.SetHead(single.Head.Next)
	}
	n := single.Head
	for n != nil {
		if n.HasNext() && n.HasValue(value) {
			n.Next = n.Next.Next
		}
	}
	return false
}

// Set assigns the provided Single-Linked-List-Node to both the Head and Tail of the Linked-List.
func (single *LinkedList) Set(s *node.Single) *LinkedList {
	single.Head = s
	single.Tail = single.Head
	return single
}

// SetHead sets the Linked-List Head.
func (single *LinkedList) SetHead(s *node.Single) *LinkedList {
	single.Head = s
	return single
}

// SetTail sets the Linked-List Tail.
func (single *LinkedList) SetTail(s *node.Single) *LinkedList {
	single.Tail = s
	return single
}

// Shift sets the Linked-List Head to the adjacent Linked-List-Node.
func (single *LinkedList) Shift() *LinkedList {
	if single.HasHead() && single.Head.HasNext() {
		return single.SetHead(single.Head.Next)
	}
	single.Head = nil
	single.Tail = nil
	return single
}

// ShiftGet sets the Linked-List head to the adjacent Linked-List-Node and returns the previous Linked-List Head's value.
func (single *LinkedList) ShiftGet() interface{} {
	if single.HasHead() {
		value := single.Head.Value
		if single.Head.HasNext() {
			single.SetHead(single.Head.Next)
		} else {
			single.Head = nil
			single.Tail = nil
		}
		return value
	}
	return nil
}

// Size computes the number of Single-Linked-List-Nodes in the Linked-List.
func (single *LinkedList) Size() int {
	n := single.Head
	s := 0
	for n != nil {
		n = n.Next
		s = s + 1
	}
	return s
}

// Walk prints all the stored values in the Linked-List using fmt.Println.
func (single *LinkedList) Walk() {
	n := single.Head
	for n != nil {
		fmt.Println(n)
		n = n.Next
	}
}

var _ linkedList = (*LinkedList)(nil)

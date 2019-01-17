// Package double exports a Double-Linked-List structure. Double-Linked-List
// provides the insertion, searching and deletion methods while offering
// some extensions to the base Double-Linked-List interface. Package also
// exports a Double-Linked-List pointer instantiation function.
package double

import (
	"fmt"

	node "github.com/gellel/gostructures/abstracts/linked-lists/double/double-node"
)

// LinkedList declares the interface for a Double-Linked-List.
type linkedList interface {
	Append(value interface{}) *LinkedList
	Find(value interface{}) *node.Double
	FindAll(value interface{}) []*node.Double
	HasHead() bool
	HasTail() bool
	InsertAfter(d *node.Double, value interface{}) *LinkedList
	InsertBefore(d *node.Double, value interface{}) *LinkedList
	IsEmpty() bool
	IsPopulated() bool
	Prepend(value interface{}) *LinkedList
	Remove(value interface{}) bool
	Set(d *node.Double) *LinkedList
	SetHead(d *node.Double) *LinkedList
	SetTail(d *node.Double) *LinkedList
	Size() int
	Walk()
	WalkReverse()
}

// LinkedList defines the Double-Linked-List data structure.
// Double-Linked-List holds a sequence of Double-Linked-List nodes.
// Each connection can be traversed from one node to another provided
// that they have a non Nil previous or next. Double-Linked-List can
// contain any mixture of data.
type LinkedList struct {
	Head *node.Double // First Double-Linked-List node.
	Tail *node.Double // Last Double-Linked-List node.
}

// New instantiates a new Double-Linked-List.
func New() *LinkedList {
	return &LinkedList{}
}

// Append adds a new Double-Linked-List-Node to the Tail of the Double-Linked-List.
func (double *LinkedList) Append(value interface{}) *LinkedList {
	if double.IsEmpty() {
		return double.Set(node.New(value))
	}
	return double.SetTail(double.Tail.AddNext(value, true).AssignPrevious(double.Tail, false))
}

// Find searches for a values stored in a Double-Linked-List-Node in the Double-Linked-List.
func (double *LinkedList) Find(value interface{}) *node.Double {
	n := double.Head
	for n != nil {
		if n.HasValue(value) {
			return n
		}
		n = n.Next
	}
	return nil
}

// FindAll searches for all values that matches the argument value in the Double-Linked-List.
func (double *LinkedList) FindAll(value interface{}) []*node.Double {
	s := make([]*node.Double, 0)
	n := double.Head
	for n != nil {
		if n.HasValue(value) {
			s = append(s, n)
		}
		n = n.Next
	}
	return s
}

// HasHead checks whether the Double-Linked-List has non-nil value stored in the Linked-List Head.
func (double *LinkedList) HasHead() bool {
	return double.Head != nil
}

// HasTail checks whether the Double-Linked-List has non-nil value stored in the Linked-List Tail.
func (double *LinkedList) HasTail() bool {
	return double.Tail != nil
}

// InsertAfter adds a new Double-Linked-List-Node after the argument Double-Linked-List-Node within the Double-Linked-List.
func (double *LinkedList) InsertAfter(d *node.Double, value interface{}) *LinkedList {
	if double.IsEmpty() {
		double.SetHead(d.AddNext(value, false)).SetTail(d.Next)
	} else if d.HasNext() {
		d.AssignNext(node.New(value).AssignNext(d.Next, false), false)
	} else {
		double.Append(value)
	}
	return double
}

// InsertBefore adds a new Double-Linked-List-Node before the argument Double-Linked-List-Node within the Double-Linked-List.
func (double *LinkedList) InsertBefore(d *node.Double, value interface{}) *LinkedList {
	if double.IsEmpty() {
		double.SetHead(d.AddPrevious(value, true).AssignNext(d, false))
	} else if d.HasPrevious() {
		d.Previous.AddNext(value, true).AssignNext(d, false) // access d.Previous node. create & assign new node. append d to new node as next.
	} else {
		double.Prepend(value)
	}
	return double
}

// IsEmpty checks whether the Double-Linked-List contains a nil value in both the Linked-List Head and Linked-List Tail.
func (double *LinkedList) IsEmpty() bool {
	return double.Head == nil && double.Tail == nil
}

// IsPopulated checks whether the Double-Linked-List contains a Double-Linked-List-Node in the Linked-List Head and Linked-List Tail in the Double-Linked-List.
func (double *LinkedList) IsPopulated() bool {
	return double.Head != nil && double.Tail != nil
}

// Prepend adds a new Double-Linked-List-Node at the beginning of the Double-Linked-List.
func (double *LinkedList) Prepend(value interface{}) *LinkedList {
	if double.IsEmpty() {
		return double.Set(node.New(value))
	}
	return double.SetTail(double.Tail.AddNext(value, true))
}

func (double *LinkedList) Remove(value interface{}) bool {
	if double.IsEmpty() {
		return false
	}
	for double.HasHead() && double.Head.HasValue(value) {
		double.SetHead(double.Head.RemovePrevious().Next)
	}
	n := double.Head
	for n != nil {
		if n.HasValue(value) {
			if n.HasNext() {
				n.Previous.Next = n.Next
			}
		}
		n = n.Next
	}
	return false
}

// Set adds the argument Double-Linked-List-Node to both the Double-Linked-List Head and Double-Linked-List Tail.
func (double *LinkedList) Set(d *node.Double) *LinkedList {
	return double.SetHead(d).SetTail(double.Head)
}

// SetHead assigns the argument Double-Linked-List-Node to the Double-Linked-List Head.
func (double *LinkedList) SetHead(d *node.Double) *LinkedList {
	double.Head = d
	return double
}

// SetTail assigns the argument Double-Linked-List-Node to the Double-Linked-List Tail.
func (double *LinkedList) SetTail(d *node.Double) *LinkedList {
	double.Tail = d
	return double
}

// Size computes the number of Double-Linked-List-Node's in the Double-Linked-List.
func (double *LinkedList) Size() int {
	n := double.Head
	s := 0
	for n != nil {
		n = n.Next
		s = s + 1
	}
	return s
}

// Walk iterates from the Double-Linked-List Head through to the Double-Linked-List Tail.
func (double *LinkedList) Walk() {
	n := double.Head
	for n != nil {
		fmt.Println(n)
		n = n.Next
	}
}

// WalkReverse iterates from the Double-Linked-List Tail through to the Double-Linked-List Head.
func (double *LinkedList) WalkReverse() {
	n := double.Tail
	for n != nil {
		fmt.Println(n)
		n = n.Previous
	}
}

var _ linkedList = (*LinkedList)(nil)

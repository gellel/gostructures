package linkedlist

import (
	"fmt"

	"github.com/gellel/gostructures/linkedlist/l"
	"github.com/gellel/gostructures/linkedlist/node"
)

// LinkedList of interfaces.
type LinkedList struct {
	l.LinkedList
	Head *node.Node
	Tail *node.Node
}

// New instantiates a new LinkedList pointer.
func New() *LinkedList {
	return &LinkedList{}
}

// Add enqueues a new LinkedList.Node.
// Value on add is placed at the tail of the LinkedList.
func (linkedList *LinkedList) Add(property interface{}) *LinkedList {

	n := node.New(property)
	// check LinkedList.Head && LinkedList.Tail is empty.
	if linkedList.isEmpty() {
		return linkedList.addHead(n)
	}
	// add to tail
	return linkedList.addTail(n)
}

// Append adds new LinkedList.Node to the LinkedList.Tail.
func (linkedList *LinkedList) Append(property interface{}) *LinkedList {

	n := node.New(property)

	// handle empty LinkedList.Head exception.
	if linkedList.isEmpty() {
		return linkedList.addHead(n)
	}

	linkedList.Tail.Next = n

	linkedList.Tail = n

	return linkedList
}

// Delete removes one or many LinkedList.Nodes.
func (linkedList *LinkedList) Delete(value interface{}) *node.Node {

	// set temp variable for storing LinkedList.Node that is
	// going to be deleted.
	var d *node.Node

	if linkedList.Head == nil {
		return nil
	}

	// update LinkedList.Head until a LinkedList.Node.Value is
	// found that does not contain the target value.
	for linkedList.Head != nil && linkedList.Head.Value == value {

		d = linkedList.Head

		linkedList.Head = linkedList.Head.Next
	}

	// store current LinkedList.Head
	// to check against next set of targets.
	n := linkedList.Head

	if !(n == nil) {
		// look ahead and iterate over LinkedList.Node(s)
		for !(n.Next == nil) {
			if n.Next.Value == value {
				d = n.Next
				// set LinkedList.Node to its LinkedList.Node.Next
				// results in n(a).Next -> n(b).Value(target) -> n(c)
				// becoming n(a).Next -> n(c)
				n.Next = n.Next.Next
			} else {
				// update n to prevent infininte loop.
				n = n.Next
			}
		}
	}

	// check if LinkedList.Tail is target and delete.
	// n should be the previous LinkedList.Node.
	if linkedList.Tail.Value == value {
		linkedList.Tail = n
	}

	return d
}

// Prepend enqueues a new LinkedList.Node ahead of the LinkedList.Head.
// LinkedList.Head becomes the the Node.Next of the new LinkedList.Node.
func (linkedList *LinkedList) Prepend(property interface{}) *LinkedList {

	n := node.New(property)

	if linkedList.isEmpty() {
		return linkedList.addHead(n)
	}

	n.Next = linkedList.Head

	linkedList.Head = n

	return linkedList
}

// Reverse reorders the stored LinkedList.Nodes.
func (linkedList *LinkedList) Reverse() *LinkedList {

	current := linkedList.Head

	var previous *node.Node

	var next *node.Node

	for !(current == nil) {
		// store current LinkedList.Node's next LinkedList.Node.
		// this enables the loop to continue for remaining LinkedList.Nodes.
		next = current.Next
		// set the currented LinkedList.Node to reference the previous LinkedList.Node
		// before: previous -> current -> next
		// after: current -> previous -> ?
		current.Next = previous
		// store previous (current) LinkedList.Node.
		previous = current
		// shift LinkedList.Node to next node.
		current = next
	}

	// set LinkedList.Tail to first LinkedList.Node (Head)
	linkedList.Tail = linkedList.Head

	linkedList.Head = previous

	return linkedList
}

// Search checks that a provided value is within
// the LinkedList.
func (linkedList *LinkedList) Search(value interface{}) bool {

	n := linkedList.Head

	for n != nil && n.Value != nil {

		if n.Value == value {
			return true
		}

		n = n.Next
	}
	return false
}

// SizeOf returns an integer representing number
// of connected LinkedList.Nodes.
func (linkedList *LinkedList) SizeOf() int {
	s := 0

	n := linkedList.Head

	for !(n == nil) {
		s = s + 1

		n = n.Next
	}
	return s
}

// Walk moves through N items and logs its contents.
func (linkedList *LinkedList) Walk() {

	n := linkedList.Head

	for !(n == nil) {

		fmt.Println(n.Value)

		n = n.Next
	}
}

func (linkedList *LinkedList) addHead(n *node.Node) *LinkedList {

	linkedList.Head = n

	linkedList.Tail = linkedList.Head

	return linkedList
}

func (linkedList *LinkedList) addTail(n *node.Node) *LinkedList {

	linkedList.Tail.Next = n

	linkedList.Tail = n

	return linkedList
}

func (linkedList *LinkedList) hasHead() bool {
	return linkedList.Head != nil
}

func (linkedList *LinkedList) hasTail() bool {
	return linkedList.Tail != nil
}

func (linkedList *LinkedList) isEmpty() bool {
	return !(linkedList.notEmpty())
}

func (linkedList *LinkedList) notEmpty() bool {
	return (linkedList.hasHead() && linkedList.hasTail())
}

package linkedlist

import (
	"fmt"

	"github.com/gellel/gostructures/linkedlist/node"
)

type LinkedList struct {
	Head *node.Node
	Tail *node.Node
}

func New() *LinkedList {
	return &LinkedList{}
}

// Add enqueues a new LinkedList.Node.
// Value on add is placed at the tail of the LinkedList.
func (linkedList *LinkedList) Add(property interface{}) *LinkedList {

	n := node.New(property)

	if linkedList.isEmpty() {
		return linkedList.addHead(n)
	}
	return linkedList.addTail(n)
}

// Delete removes a LinkedList.Node but reattaches the orphaned LinkedList.Node.
func (linkedList *LinkedList) Delete(value interface{}) bool {

	if linkedList.Head == nil {
		return false
	}

	n := linkedList.Head

	if n.Value == value {

		// reset LinkedList.Head and LinkedList.Tail
		if linkedList.Head == linkedList.Tail {

			linkedList.Head = nil

			linkedList.Tail = linkedList.Head

		} else {
			// set LinkedList.Head to LinkedList.Head.Next
			linkedList.Head = n.Next
		}

		return true
	}

	// move through O(n) trying to find target.
	for n.Next != nil && n.Next.Value != value {
		n = n.Next
	}

	if n.Next != nil {

		if n.Next == linkedList.Tail {

			linkedList.Tail = n
		}

		n.Next = n.Next.Next

		return true
	}

	return false
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

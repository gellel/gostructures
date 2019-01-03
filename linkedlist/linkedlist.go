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

func (linkedList *LinkedList) Add(property interface{}) *LinkedList {

	n := node.New(property)

	if linkedList.isEmpty() {
		return linkedList.addFirst(n)
	}
	return linkedList.add(n)
}

func (linkedList *LinkedList) Walk() {

	n := linkedList.Head

	for !(n == nil) {

		fmt.Println(n.Value)

		n = n.Next
	}
}

func (linkedList *LinkedList) addFirst(n *node.Node) *LinkedList {

	linkedList.Head = n

	linkedList.Tail = linkedList.Head

	return linkedList
}

func (linkedList *LinkedList) add(n *node.Node) *LinkedList {

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

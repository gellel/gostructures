package linkedlist

import (
	"github.com/gellel/gostructures/linkedlist/node"
)

type LinkedList struct {
	Node *node.Node
}

func New() *LinkedList {
	return &LinkedList{}
}

func (LinkedList *LinkedList) Add(property interface{}) *LinkedList {

	n := node.New(property)

	if LinkedList.hasFirst() {
		return LinkedList.addFirst(n)
	}
	return LinkedList.add(n)
}

func (linkedList *LinkedList) addFirst(n *node.Node) *LinkedList {

	linkedList.Node = n

	return linkedList
}

func (linkedList *LinkedList) add(n *node.Node) *LinkedList {
	c := linkedList.Node

	for !(c.Next == nil) {
		c = c.Next
	}
}

func (linkedList *LinkedList) hasFirst() bool {
	return !(linkedList.Node == nil)
}

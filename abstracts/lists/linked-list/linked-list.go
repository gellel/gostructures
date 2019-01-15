package linkedlist

import (
    "github.com/gellel/gostructures/abstracts/lists/linked-list/list-node"
)

type LinkedList struct {
    Head *node.Node
    Tail *node.Node
}

// New instantiates a new LinkedList pointer.
func New() *LinkedList {
    return &LinkedList{}
}

// AddHead sets the initial *LinkedList.Head.
func (linkedList *LinkedList) AddHead(node *node.Node) *LinkedList {
    // set *LinkedList.Head to provided
    // *LinkedList.Node. removes all current
    // connections if called manually.
    linkedList.Head = node
    // assume that the method was manually
    // invoked and update the *LinkedList.Tail
    // to become the current *LinkedList.Head.
    linkedList.Tail = linkedList.Head

    return linkedList
}

// Append inserts a new *LinkedList.Node at
// the end (Tail) of the LinkedList.
// When the *LinkedList is empty, this new *LinkedList.Node
// becomes both the *LinkedList.Head and *LinkedList.Tail.
func (linkedList *LinkedList) Append(property interface{}) *LinkedList {
    if linkedList.IsEmpty() {
        return linkedList.AddHead(node.New(property))
    }
    return linkedList.AppendTail(node.New(property))
}

// AppendTail updates *LinkedList.Tail.Next to
// contain an adjacent reference to the provided *LinkedList.Node.
// The provided *LinkedList.Node then becomes *LinkedList.Tail.
func (linkedList *LinkedList) AppendTail(node *node.Node) *LinkedList {
    // update the current *LinkedList.Tail to
    // contain a adjacent reference to the provided 
    // *LinkedList.Node.
    linkedList.Tail.Next = node
    // set the current *LinkedList.Tail
    // to become the provided *LinkedList.Node.
    linkedList.Tail = node

    return linkedList
}

// Contain searches across O(n) entries attempting
// to find a *LinkedList.Node that contains a
// corresponding value.
func (linkedList *LinkedList) Contains(property interface{}) bool {
    return linkedList.Find(property) != nil
}

// Delete removes a connection between from one
// *LinkedList.Node to an adjacent *LinkedList.Node.
// Operation may require O(n) iterations before the
// target *LinkedList.Node is found. Best case is
// Omega(1) if value is contained in either *LinkedList.Head
// or *LinkedList.Tail.
func (linkedList *LinkedList) Delete(property interface{}) *LinkedList {
    return linkedList
}

// Find searches for an adjacent *LinkedList.Node that
// contains a value of the provided interface type.
// Operation may require O(n) iterations before the
// target *LinkedList.Node is found. Best case is
// Omega(1) if value is contained in either *LinkedList.Head
// or *LinkedList.Tail.
func (linkedList *LinkedList) Find(property interface{}) *node.Node {
    // attempt to perform Omega(1) runtime
    // by checking whether *LinkedList.Head
    // or *LinkedList.Tail contains the required value.
    if linkedList.HasHead() && linked.Head.Value == property {
        return linkedList.Head
    } else if linkedList.HasTail() && linked.Tail.Value == property {
        return linkedList.Tail
    }
    // otherwise collect first *LinkedList.Node
    node := LinkedList.Head
    // iterate while current *LinkedList.Node
    // contains an adjacent *LinkedList.Node.
    for node.Next && node.Next != linkedList.Tail {
        // check whether current adjacent
        // *LinkedList.Node.Value is the
        // provided value.
        if node.Next.Value == property {
            return node.Next
        }
        // update current node and
        // continue moving through the
        // connected sequence of LinkedList.Node.
        node = node.Next
    }
    return nil
}

// HasEmptyHead checks whether *LinkedList.Head
// contains a Nil reference.
func (linkedList *LinkedList) HasEmptyHead() bool {
    return linkedList.Head == nil
}

// HasHead inverts the result of *LinkedList.HasEmptyHead
// and returns true should a non-Nil reference be
// assigned to the *LinkedList.Head property.
func (linkedList *LinkedList) HasHead() bool {
    return !linkedList.HasEmptyHead()
}

// HasEmptyTail checks whether *LinkedList.Tail
// contains a Nil reference.
func (linkedList *LinkedList) HasEmptyTail() bool {
    return linkedList.Tail == nil
}

// HasTail inverts the result of *LinkedList.HasEmptyTail
// and returns true should a non-Nil reference be
// assigned to the *LinkedList.Tail property.
func (linkedList *LinkedList) HasTail() bool {
    return !linkedList.HasEmptyTail()
}

// IsEmpty checks whether both *LinkedList.Head
// and *LinkedList.Tail contains a Nil reference.
// A newly instantiated *LinkedList will yield true.
func (linkedList *LinkedList) IsEmpty() bool {
    return linkedList.HasEmptyHead() && linkedList.HasEmptyTail()
}

// IsPopulated inverts the result *LinkedList.IsEmpty
// and returns true should both *LinkedList.Head
// and *LinkedList.Tail contain a non-Nil property.
// A *LinkedList of at least one *LinkedList.Node
// will yield true.
func (linkedList *LinkedList) IsPopulated() bool {
    return !linkedList.IsEmpty()
}

// Prepend inserts a new *LinkedList.Node at the
// beginning of the *LinkedList. When populated, current
// *LinkedList.Head is updated to become new *LinkedList.Node. 
// The previous *LinkedList.Head then becomes the 
// provided LinkedList.Node.Next reference.
func (linkedList *LinkedList) Prepend(property interface{}) *LinkedList {
    if linkedList.IsEmpty() {
        return linkedList.AddHead(node.New(property))
    }
    return linkedList.PrependHead(node.New(property))
}

// PrependHead performs the LinkedList insertion operation.
// Method focuses on updating the *LinkedList.Head
// to reference the provided *LinkedList.Node.
// Manual invocation requires operator to manage
// *LinkedList.Tail state.
func (linkedList *LinkedList) PrependHead(node *node.Node) *LinkedList {
    // set *LinkedList.Node to contain an
    // adjacent reference to current
    // *LinkedList.Head.
    node.Next = linkedList.Head
    // set provided *LinkedList.Node
    // as the new *LinkedList.Head.
    linkedList.Head = node
    
    return linkedList
}


var _ definition.LinkedList = (*LinkedList)(nil)
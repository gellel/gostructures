package linkedlist

type LinkedList struct {
    Head *Node
    Tail *Node
}

// New instantiates a new LinkedList pointer.
func New() *LinkedList {
    return &LinkedList{}
}

// Append inserts a new *LinkedList.Node at
// the beginning (Head) of the LinkedList.
// When the *LinkedList is empty, this new *LinkedList.Node
// becomes both the *LinkedList.Tail and *LinkedList.Head.
func (linkedList *LinkedList) Append(property interface{}) *LinkedList {
    return linkedlist
}

// Contain searches across O(n) entries attempting
// to find a *LinkedList.Node that contains a
// corresponding value.
func (linkedList *LinkedList) Contains(property interface{}) bool {
    return false
}

// Delete removes a connection between from one
// *LinkedList.Node to an adjacent *LinkedList.Node.
// Operation may require O(n) iterations before the
// target *LinkedList.Node is found. Best case is
// Omega(1) if value is contained in either *LinkedList.Head
// or *LinkedList.Tail.
func (linkedlist *LinkedList) Delete(property interface{}) *LinkedList {
    return linkedList
}

// Prepend inserts a new *LinkedList.Node at the
// beginning of the *LinkedList.
// When populated, current *LinkedList.Head becomes 
// the new *LinkedList.Tail and LinkedList.Head 
// becomes the new LinkedList.Node
// with a Next reference to LinkedList.Head.
func (linkedlist *LinkedList) Prepend(property interface{}) *LinkedList {
    return linkedList
}

// Find searches for an adjacent *LinkedList.Node that
// contains a value of the provided interface type.
// Operation may require O(n) iterations before the
// target *LinkedList.Node is found. Best case is
// Omega(1) if value is contained in either *LinkedList.Head
// or *LinkedList.Tail.
func (linkedList *LinkedList) Find(property interface{}) *node.Node {
    return nil
}

// HasEmptyHead checks whether *LinkedList.Head
// contains a Nil reference.
func (linkedlist *LinkedList) HasEmptyHead() bool {
    return linkedList.Head == nil
}

// HasHead inverts the result of *LinkedList.HasEmptyHead
// and returns true should a non-Nil reference be
// assigned to the *LinkedList.Head property.
func (linkedList *LinkedList) HasHead() bool {
    return !linkedlist.HasEmptyHead()
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
    return !linkedlist.IsEmpty()
}

var _ definition.LinkedList = (*LinkedList)(nil)
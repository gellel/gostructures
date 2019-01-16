package single

import (
	"fmt"

	node "github.com/gellel/gostructures/abstracts/linked-lists/single/single-node"
)

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
	Size() int
	Walk()
}

type LinkedList struct {
	Head *node.Single
	Tail *node.Single
}

func New() *LinkedList {
	return &LinkedList{}
}

func (single *LinkedList) Append(value interface{}) *LinkedList {
	if single.IsEmpty() {
		return single.Set(node.New(value))
	}
	return single.SetTail(single.Tail.AddNext(value, true))
}

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

func (single *LinkedList) HasHead() bool {
	return single.Head != nil
}

func (single *LinkedList) HasTail() bool {
	return single.Tail != nil
}

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

func (single *LinkedList) IsEmpty() bool {
	return single.Head == nil && single.Tail == nil
}

func (single *LinkedList) IsPopulated() bool {
	return single.Head != nil && single.Tail != nil
}

func (single *LinkedList) Prepend(value interface{}) *LinkedList {
	if single.IsEmpty() {
		return single.Set(node.New(value))
	}
	return single.SetHead(node.New(value).AssignNext(single.Head, false))
}

func (single *LinkedList) Remove(value interface{}) bool {
	return false
}

func (single *LinkedList) Set(s *node.Single) *LinkedList {
	single.Head = s
	single.Tail = single.Head
	return single
}

func (single *LinkedList) SetHead(s *node.Single) *LinkedList {
	single.Head = s
	return single
}

func (single *LinkedList) SetTail(s *node.Single) *LinkedList {
	single.Tail = s
	return single
}

func (single *LinkedList) Size() int {
	n := single.Head
	s := 0
	for n != nil {
		n = n.Next
		s = s + 1
	}
	return s
}

func (single *LinkedList) Walk() {
	n := single.Head
	for n != nil {
		fmt.Println(n)
		n = n.Next
	}
}

var _ linkedList = (*LinkedList)(nil)

/*

// LinkedList structure.
type LinkedList struct {
	Head *node.Node
	Tail *node.Node
}

// New instantiates a new LinkedList pointer.
func New() *LinkedList {
	return &LinkedList{}
}

// AddHead sets *LinkedList.Head.
func (linkedList *LinkedList) AddHead(node *node.Node) *LinkedList {
	// set *LinkedList.Head to provided *LinkedList.Node.
	linkedList.Head = node
	// assumes *LinkedList.Tail is empty. assign *LinkedList.Head
	// to be the current *LinkedList.Head.
	linkedList.Tail = linkedList.Head
	// return the accessed LinkedList pointer.
	return linkedList
}

// Append inserts a new *LinkedList.Node at
// the end (Tail) of the LinkedList.
// When the *LinkedList is empty, this new *LinkedList.Node
// becomes both the *LinkedList.Head and *LinkedList.Tail.
func (linkedList *LinkedList) Append(property interface{}) *LinkedList {
	// perform first insertion when
	// there is no connections.
	if linkedList.IsEmpty() {
		// return the accessed LinkedList pointer.
		return linkedList.AddHead(node.New(property))
	}
	// return the accessed LinkedList pointer.
	return linkedList.AppendTail(node.New(property))
}

// AppendTail updates *LinkedList.Tail.Next to contain an adjacent
// reference to the provided *LinkedList.Node.
// The provided *LinkedList.Node then becomes *LinkedList.Tail.
func (linkedList *LinkedList) AppendTail(node *node.Node) *LinkedList {
	// update the current *LinkedList.Tail to
	// contain a adjacent reference to the provided *LinkedList.Node.
	linkedList.Tail.Next = node
	// set the current *LinkedList.Tail
	// to become the provided *LinkedList.Node.
	linkedList.Tail = node
	// return the accessed LinkedList pointer.
	return linkedList
}

// Contains searches across O(n) entries attempting
// to find a *LinkedList.Node that holds a corresponding value.
func (linkedList *LinkedList) Contains(property interface{}) bool {
	return linkedList.Find(property) != nil
}

// Delete removes a connection between from one LinkedList.Node
// to an adjacent *LinkedList.Node. Operation requires
// O(n) iterations to remove provided property.
func (linkedList *LinkedList) Delete(property interface{}) *LinkedList {

	// exit early if there is empty.
	if linkedList.HasEmptyHead() {
		return linkedList
	}
	node := linkedList.Head
	// update *LinkedList.Head until assigned *LinkedList.Node.Value
	// does not contain the target value.
	for linkedList.HasHead() && linkedList.Head.Value == property {
		node = linkedList.Head
		linkedList.Head = linkedList.Head.Next
	}
	// store current *LinkedList.Head to check against next set of *LinkedList.Nodes.
	node = linkedList.Head
	if node != nil {
		// look ahead and iterate over adjacent *LinkedList.Nodes
		for node.Next != nil {
			if node.Next.Value == property {
				// set LinkedList.Node to its LinkedList.Node.Next
				// results in n(a).Next -> n(b).Value(target) -> n(c)
				// becoming n(a).Next -> n(c)
				node.Next = node.Next.Next
			} else {
				// update n to prevent infininte loop.
				node = node.Next
			}
		}
	}

	// check if *LinkedList.Tail is target and delete.
	// node should be the previous *LinkedList.Node.
	if linkedList.Tail.Value == property {
		linkedList.Tail = node
	}

	return linkedList
}

// Find searches for an adjacent *LinkedList.Node that
// contains a value of the provided interface type.
// Operation may require O(n) iterations before the
// target *LinkedList.Node is found. Best case is
// Omega(1) if value is contained in either *LinkedList.Head
// or *LinkedList.Tail.
func (linkedList *LinkedList) Find(property interface{}) *node.Node {
	// attempt to perform Omega(1) runtime by checking
	// whether *LinkedList.Head contains the required value.
	if linkedList.HasHead() && linkedList.Head.Value == property {
		return linkedList.Head
	}
	// attempt to perform Omega(1) runtime by checking
	// whether *LinkedList.Tail contains the required value.
	if linkedList.HasTail() && linkedList.Tail.Value == property {
		return linkedList.Tail
	}
	// otherwise perform O(n).
	// start by collecting *LinkedList.Head.
	node := linkedList.Head
	// iterate while current *LinkedList.Node
	// contains an adjacent *LinkedList.Node.
	for node.Next != nil && node.Next != linkedList.Tail {
		// check whether current adjacent *LinkedList.Node.Value
		// is the provided value.
		if node.Next.Value == property {
			return node.Next
		}
		// update current node and continue moving through the
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
	// perform first insertion when
	// there is no connections.
	if linkedList.IsEmpty() {
		// return accessed LinkedList pointer.
		return linkedList.AddHead(node.New(property))
	}
	// return the accessed LinkedList pointer.
	return linkedList.PrependHead(node.New(property))
}

// PrependHead performs the LinkedList insertion operation.
// Method focuses on updating the *LinkedList.Head
// to reference the provided *LinkedList.Node.
// Manual invocation requires operator to manage *LinkedList.Tail state.
func (linkedList *LinkedList) PrependHead(node *node.Node) *LinkedList {
	// set *LinkedList.Node to contain an
	// adjacent reference to current *LinkedList.Head.
	node.Next = linkedList.Head
	// set provided *LinkedList.Node
	// as the new *LinkedList.Head.
	linkedList.Head = node
	// return the accessed LinkedList pointer.
	return linkedList
}

// SizeOf computes the current size
// of the accessed *LinkedList. Size
// is computed from 0.
func (linkedList *LinkedList) SizeOf() int {
	// end early if there are
	// no adjacent LinkedList.Nodes.
	if linkedList.IsEmpty() {
		return 0
	}
	n := 0
	node := linkedList.Head
	// iterate across adjacent *LinkedList.Nodes
	// until the cursor has reached the *LinkedList.Tail.
	for node.Next != nil {
		n = n + 1
		node = node.Next
	}
	// return the computed size.
	return n
}

//var _ definition.LinkedList = (*LinkedList)(nil)

*/

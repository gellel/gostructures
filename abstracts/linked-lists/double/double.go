// Package double exports a Double Linked-List structure. Double Linked-List
// provides the insertion, searching and deletion methods while offering
// some extensions to the base Double Linked-List interface. Package also
// exports a Double Linked-List pointer instantiation function.
package double

import node "github.com/gellel/gostructures/abstracts/linked-lists/double/double-node"

// LinkedList declares the interface for a doubly Linked-List.
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

// LinkedList defines the Double Linked-List data structure.
// Double Linked-List holds a sequence of Double Linked-List nodes.
// Each connection can be traversed from one node to another provided
// that they have a non Nil previous or next. Double Linked-List can
// contain any mixture of data.
type LinkedList struct {
	Head *node.Node // First Double Linked-List node.
    Tail *node.Node // Last Double Linked-List node.
}

func (double *LinkedList) Append(value interface{}) *LinkedList {
    if double.IsEmpty() {
        return double.Set(node.New(value))
    }
	return double.SetTail(double.Tail.AddNext(value, true))
}

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

func (double *LinkedList) HasHead() bool {
    return double.Head != nil
}

func (double *LinkedList) HasTail() bool {
    return double.Tail != nil
}

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

func (double *LinkedList) InsertBefore(d *node.Double, value interface{}) *LinkedList {
    if double.IsEmpty() {
        double.SetHead(d.AddPrevious(value, true).AssignNext(d, false))
    } else if d.HasPrevious() {
        d.Previous.AddNext(value, true).AssignNext(d, false) // access d.Previous node. create & assign new node. append d to new node as next.
    } else {
        double.Prepend(value)
    }
}

func (double *LinkedList) IsEmpty() bool {
    return double.Head == nil && double.Tail == nil
}

func (double *LinkedList) IsPopulated() bool {
    return double.Head != nil && double.Tail != nil
}


func (double *LinkedList) Prepend(value interface{}) *LinkedList {
    if double.IsEmpty() {
        return double.Set(value)
    }
	return double.SetTail(double.Tail.AddNext(value, true))
}

func (double *LinkedList) Remove(value interface{}) bool {

}

func (double *LinkedList) Set(d *node.Double) *LinkedList {
    return double.SetHead(d).SetTail(double.Head)
}

func (double *LinkedList) SetHead(d *node.Double) *LinkedList {
    double.Head = d
    return double
}

func (double *LinkedList) SetTail(d *node.Double) *LinkedList {
    double.Tail = d
    return double
}

func (double *LinkedList) Size() int {
    n := double.Head
	s := 0
	for n != nil {
		n = n.Next
		s = s + 1
	}
	return s
}

func (double *LinkedList) Walk() {
    n := double.Head
	for n != nil {
		fmt.Println(n)
		n = n.Next
	}
}

func (double *LinkedList) WalkReverse() {
    n := single.Tail
	for n != nil {
		fmt.Println(n)
		n = n.Previous
	}
}

var _ linkedList = (*LinkedList)(nil)

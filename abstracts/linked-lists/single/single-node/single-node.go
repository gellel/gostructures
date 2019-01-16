// Package node exports a singly Linked-List-Node. A singly Linked-List-Node
// is a struct contained within a singly Linked-List. Linked-List-Node
// can contain a value attribute of any type as well as an optional reference to an
// adjacent Linked-List-Node. Package also exports a Linked-List-Node
// pointer instantiation method (New).
package node

// SINGLE declares the interface for a singly Linked-List-Node.
type SINGLE interface {
	AddNext(value interface{}, next bool) *Single
	AssignNext(s *Single, next bool) *Single
	AssignValue(value interface{}) *Single
	EmptyNext() bool
	HasNext() bool
	HasValue(value interface{}) bool
	IsEmpty() bool
	IsSet() bool
	RemoveNext() *Single
}

// Single declares the struct for a Linked-List-Node for a singly Linked-List.
type Single struct {
	Next  *Single     // Connection to an adjacent Linked-List-Node.
	Value interface{} // Property assigned to the Linked-List-Node.
}

// New instantiates a new singly Linked-List-Node.
func New(property interface{}) *Single {
	return &Single{Value: property}
}

// AddNext sets the accessed Linked-List-Node's adjcent Linked-List-Node;
// returns the new Linked-List-Node if next is set to true.
func (single *Single) AddNext(value interface{}, next bool) *Single {
	single.Next = New(value)
	if next {
		return single.Next
	}
	return single
}

// AssignNext sets the argument Linked-List-Node as the accessed Linked-List-Node's adjacent Linked-List-Node; returns the argument Linked-List-Node if next is set to true.
func (single *Single) AssignNext(s *Single, next bool) *Single {
	single.Next = s
	if next {
		return single.Next
	}
	return single
}

// AssignValue sets the accessed Linked-List-Node's value.
func (single *Single) AssignValue(value interface{}) *Single {
	single.Value = value
	return single
}

// EmptyNext checks whether the Linked-List-Node has an empty adjacent Linked-List-Node.
func (single *Single) EmptyNext() bool {
	return single.Next == nil
}

// HasNext checks whether the Linked-List-Node has an adjacent Linked-List-Node.
func (single *Single) HasNext() bool {
	return single.Next != nil
}

// HasValue checks whether the value argument is stored in the Linked-List-Node.
func (single *Single) HasValue(value interface{}) bool {
	return single.Value == value
}

// IsEmpty checks whether the Linked-List-Node contains a Nil value.
func (single *Single) IsEmpty() bool {
	return single.Value == nil
}

// IsSet checks whether the Linked-List-Node contains a non Nil value.
func (single *Single) IsSet() bool {
	return single.Value != nil
}

// RemoveNext removes the accessed Linked-List-Node's adjacent Linked-List-Node; returns accessed Linked-List-Node.
func (single *Single) RemoveNext() *Single {
	single.Next = nil
	return single
}

var _ SINGLE = (*Single)(nil)

// Package node exports a Double-Linked-List-Node. A doublly Linked-List-Node
// is a struct contained within a Double-Linked-List. Linked-List-Node
// can contain a value attribute of any type as well as an optional reference to both
// a previous and adjacent Linked-List-Node. Package also exports a Linked-List-Node
// pointer instantiation method (New).
package node

// Double declares the interface for a Double-Linked-List-Node.
type double interface {
	AddPrevious(value interface{}, previous bool) *Double
	AssignPrevious(d *Double, previous bool) *Double
	AddNext(value interface{}, next bool) *Double
	AssignNext(d *Double, next bool) *Double
	AssignValue(value interface{}) *Double
	EmptyPrevious() bool
	EmptyNext() bool
	HasPrevious() bool
	HasNext() bool
	HasValue(value interface{}) bool
	IsEmpty() bool
	IsSet() bool
	RemovePrevious() *Double
	RemoveNext() *Double
}

// Double declares the struct for a Linked-List-Node for a Double-Linked-List.
type Double struct {
	Previous *Double     // Connection to a previous Linked-List-Node.
	Next     *Double     // Connection to an adjacent Linked-List-Node.
	Value    interface{} // Property assigned to the Linked-List-Node.
	// maybe should have a 'list' reference so it can confirm whether it can be added?
}

// New instantiates a new Double-Linked-List-Node.
func New(property interface{}) *Double {
	return &Double{Value: property}
}

// AddPrevious sets the accessed Linked-List-Node's previous Linked-List-Node;
// returns the new Linked-List-Node if previous is set to true.
func (double *Double) AddPrevious(value interface{}, previous bool) *Double {
	double.Previous = New(value)
	if previous {
		return double.Previous
	}
	return double
}

// AddNext sets the accessed Linked-List-Node's adjcent Linked-List-Node;
// returns the new Linked-List-Node if next is set to true.
func (double *Double) AddNext(value interface{}, next bool) *Double {
	double.Next = New(value)
	if next {
		return double.Next
	}
	return double
}

// AssignPrevious sets the argument Linked-List-Node as the accessed Linked-List-Node's previous Linked-List-Node; returns the argument Linked-List-Node if previous is set to true.
func (double *Double) AssignPrevious(p *Double, previous bool) *Double {
	double.Previous = p
	if previous {
		return double.Previous
	}
	return double
}

// AssignNext sets the argument Linked-List-Node as the accessed Linked-List-Node's adjacent Linked-List-Node; returns the argument Linked-List-Node if next is set to true.
func (double *Double) AssignNext(p *Double, next bool) *Double {
	double.Next = p
	if next {
		return double.Next
	}
	return double
}

// AssignValue sets the accessed Linked-List-Node's value.
func (double *Double) AssignValue(value interface{}) *Double {
	double.Value = value
	return double
}

// EmptyPrevious checks whether the Linked-List-Node has an empty previous Linked-List-Node.
func (double *Double) EmptyPrevious() bool {
	return double.Previous == nil
}

// EmptyNext checks whether the Linked-List-Node has an empty adjacent Linked-List-Node.
func (double *Double) EmptyNext() bool {
	return double.Next == nil
}

// HasPrevious checks whether the Linked-List-Node has a previous Linked-List-Node.
func (double *Double) HasPrevious() bool {
	return double.Previous != nil
}

// HasNext checks whether the Linked-List-Node has an adjacent Linked-List-Node.
func (double *Double) HasNext() bool {
	return double.Next != nil
}

// HasValue checks whether the value argument is stored in the Linked-List-Node.
func (double *Double) HasValue(value interface{}) bool {
	return double.Value == value
}

// IsEmpty checks whether the Linked-List-Node contains a Nil value.
func (double *Double) IsEmpty() bool {
	return double.Value == nil
}

// IsSet checks whether the Linked-List-Node contains a non Nil value.
func (double *Double) IsSet() bool {
	return double.Value != nil
}

// RemovePrevious removes the accessed Linked-List-Node's previous Linked-List-Node; returns accessed Linked-List-Node.
func (double *Double) RemovePrevious() *Double {
	double.Previous = nil
	return double
}

// RemoveNext removes the accessed Linked-List-Node's adjacent Linked-List-Node; returns accessed Linked-List-Node.
func (double *Double) RemoveNext() *Double {
	double.Next = nil
	return double
}

var _ double = (*Double)(nil)

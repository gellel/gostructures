// Package node exports a Min-Node for use in a Min-Priorty Queue.
// Min-Node is intended to be contained with the Heap used by
// the main Priority Queue.
package node

// Min declares the interface for a Min-Priority Queue.
type min interface {
	SetPriority(p int) *Min
	SetValue(value interface{}) *Min
}

// Min declares the structure for a Min-Priortiy Queue Node.
type Min struct {
	Priority int
	Value    interface{}
}

// New instantiates a new Min-Priority Queue Node pointer.
func New(priority int, value interface{}) *Min {
	return &Min{
		Priority: priority,
		Value:    value}
}

// SetPriority assigns the priorty of the accessed Min-Priority Queue Node.
func (min *Min) SetPriority(p int) *Min {
	min.Priority = p
	return min
}

// SetValue assigns the value held by the accessed Min-Priority Queue node.
func (min *Min) SetValue(value interface{}) *Min {
	min.Value = value
	return min
}

var _ min = (*Min)(nil)

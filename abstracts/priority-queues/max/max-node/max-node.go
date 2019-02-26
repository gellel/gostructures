// Package node exports a Max-Node for use in a Max-Priorty Queue.
// Max-Node is intended to be contained with the Heap used by
// the main Priority Queue.
package node

// Max declares the interface for a Max-Priority Queue.
type max interface {
	SetPriority(p int) *Max
	SetValue(value interface{}) *Max
}

// Max declares the structure for a Max-Priortiy Queue Node.
type Max struct {
	Priority int
	Value    interface{}
}

// New instantiates a new Max-Priority Queue Node pointer.
func New(priority int, value interface{}) *Max {
	return &Max{
		Priority: priority,
		Value:    value}
}

// SetPriority assigns the priorty of the accessed Max-Priority Queue Node.
func (max *Max) SetPriority(p int) *Max {
	max.Priority = p
	return max
}

// SetValue assigns the value held by the accessed Max-Priority Queue node.
func (max *Max) SetValue(value interface{}) *Max {
	max.Value = value
	return max
}

var _ max = (*Max)(nil)

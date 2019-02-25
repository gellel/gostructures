package node

type max interface {
	SetPriority(p int) *Max
	SetValue(value interface{}) *Max
}

type Max struct {
	Priority int
	Value    interface{}
}

func New(priority int, value interface{}) *Max {
	return &Max{
		Priority: priority,
		Value:    value}
}

func (max *Max) SetPriority(p int) *Max {
	max.Priority = p
	return max
}

func (max *Max) SetValue(value interface{}) *Max {
	max.Value = value
	return max
}

var _ max = (*Max)(nil)

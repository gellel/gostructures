package node

const (
	SIZE int = 2
)

type closed interface {
	IsEmpty() bool
	IsNotEmpty() bool
	Length() int
	PositionOf() int
	RemoveValue() *Closed
	SetValue(value interface{}) *Closed
	ValueOf() string
}

type Closed [SIZE]interface{}

func New(i int, value interface{}) *Closed {
	return &Closed{i, value}
}

func (closed *Closed) PositionOf() int {
}

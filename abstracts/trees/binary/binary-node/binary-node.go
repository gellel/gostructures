package node

const (
	LEFT  string = "LEFT"
	RIGHT string = "RIGHT"
)

type Binary struct {
	Left  *Binary
	Right *Binary
	Side  string
	Value float64
}

func (binary *Binary) IsEqual(value float64) bool {
	return binary.Value == value
}

func (binary *Binary) IsLess(value float64) bool {
	return binary.Value < value
}

func (binary *Binary) IsLeft() bool {
	return binary.Side == LEFT
}

func (binary *Binary) IsRight() bool {
	return binary.Side == RIGHT
}

func (binary *Binary) IsMore(value float64) bool {
	return binary.Value > value
}

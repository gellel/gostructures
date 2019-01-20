package node

import "math"

const (
	// LEFT declares constant string that provides a namespace to annotate that a Binary-node is assigned as a left child of its parent.
	LEFT string = "LEFT"
	// RIGHT declares constant string that provides a namespace to annotate that a Binary-node is assigned as a right child of its parent.
	RIGHT string = "RIGHT"
)

type binary interface {
	EmptyLeft() bool
	EmptyRight() bool
	HasLeft() bool
	HasRight() bool
	Height() int
	HeightOf() float64
	HeightLeft() float64
	HeightRight() float64
	IsChild(b *Binary) bool
	IsEqual(value float64) bool
	IsLess(value float64) bool
	IsMore(value float64) bool
}

type Binary struct {
	Left  *Binary
	Right *Binary
	Side  string
	Value float64
}

func New(value float64) *Binary {
	return &Binary{Value: value}
}

func (binary *Binary) EmptyLeft() bool {
	return binary.Left == nil
}

func (binary *Binary) EmptyRight() bool {
	return binary.Right == nil
}

func (binary *Binary) HasLeft() bool {
	return binary.Left != nil
}

func (binary *Binary) HasRight() bool {
	return binary.Right != nil
}

func (binary *Binary) Height() int {
	return int(binary.HeightOf())
}

func (binary *Binary) HeightOf() float64 {
	return math.Max(binary.HeightLeft(), binary.HeightRight())
}

func (binary *Binary) HeightLeft() float64 {
	if binary.HasLeft() {
		return (binary.Left.HeightOf() + 1.0)
	}
	return 0.0
}

func (binary *Binary) HeightRight() float64 {
	if binary.HasRight() {
		return (binary.Right.HeightOf() + 1.0)
	}
	return 0.0
}

func (binary *Binary) IsChild(b *Binary) bool {
	return binary.Left == b || binary.Right == b
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

var _ binary = (*Binary)(nil)

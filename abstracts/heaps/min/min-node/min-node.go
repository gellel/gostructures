package node

import (
	"fmt"
	"log"
	"math"
)

const (
	// LEFT declares Min.Side as LEFT.
	LEFT = "LEFT"
	// RIGHT declares Min.Side as RIGHT.
	RIGHT = "RIGHT"
	// ROOT declares Min.Side as ROOT.
	ROOT = "ROOT"
)

var (
	// Sides declares map of Min.Side sides.
	Sides = map[string]bool{
		LEFT:  true,
		RIGHT: true,
		ROOT:  true}
)

// Min declares the implementation for a Min-Heap-Node
// within the Min-Heap.
type min interface {
	AssignLeft(m *Min) *Min
	AssignParent(m *Min) *Min
	AssignRight(m *Min) *Min
	AssignSide(side string) *Min
	BubbleUp() *Min
	BubbleDown() *Min
	EmptyLeft() bool
	EmptyParent() bool
	EmptyRight() bool
	EmptySide() bool
	Insert(value float64) *Min
	InsertBinary(m *Min) *Min
	IsChild(m *Min) bool
	IsEqual(value float64) bool
	IsLess(value float64) bool
	IsLeft() bool
	IsMore(value float64) bool
	IsParent(m *Min) bool
	IsRed() bool
	IsRight() bool
	IsRoot() bool
	HasLeft() bool
	HasParent() bool
	HasRight() bool
	HasSide() bool
	Remove(value float64) *Min
	RemoveLeft() *Min
	RemoveParent() *Min
	RemoveRight() *Min
	ToFloatSlice() []float64
	ToMinSlice() []*Min
	UnsafelyAssignLeft(m *Min) *Min
	UnsafelyAssignParent(m *Min) *Min
	UnsafelyAssignRight(m *Min) *Min
	UnsafelyAssignSide(side string) *Min
	ViolatesLeft(m *Min) error
	ViolatesParent(m *Min) error
	ViolatesRight(m *Min) error
	ViolatesSide(side string) error
}

type Min struct {
	Left	*Min
	Parent	*Min
	Right 	*Min
	Side    string
	Value   float64
}

func New(value float64) *Min {
	return &Min{
		Side:   ROOT
		Value:  value
	}
}

// UnsafelyAssignLeft assigns a Left child to the accessed Min without performing validity checks.
func (min *Min) UnsafelyAssignLeft(m *Min) *Min {

	m.UnsafelyAssignSide(LEFT)

	m.UnsafelyAssignParent(min)

	min.Left = m

	return min
}

// UnsafelyAssignParent assigns a Parent Min to the accessed Min without performing validity checks.
func (min *Min) UnsafelyAssignLeft(m *Min) *Min {

	min.Parent = m

	return min
}

// UnsafelyAssignRight assigns a Right child to the accessed Min without performing validity checks.
func (min *Min) UnsafelyAssignRight(m *Min) *Min {

	m.UnsafelyAssignSide(RIGHT)

	m.UnsafelyAssignParent(min)

	min.Right = m

	return min
}

// ViolatesLeft checks that the argument Min is a valid candidate for a Left child of the accessed Min.
func (min *Min) ViolatesLeft(m *Min) error {

	if m.Value < min.Value {
		return fmt.Errorf("address (*%p) cannot hold pointer (*%p). argument struct for (*%p).Left must contain value greater than %f", &m, &min, &min, min.Value)
	}
	return nil
}

// ViolatesParent checks that the argument Min is a valid candidate for the Parent of the accessed Min.
func (min *Min) ViolatesParent(m *Min) error {

	if redBlack.Parent == nil {
		return fmt.Errorf("address (*%p) has no parent pointer", &m)
	} else if m.Value > min.Parent.Value {
		return fmt.Errorf("address (*%p) cannot hold pointer (*%p). argument Min for (*%p).Parent must not exceed %f", &m, &min, &min, min.Parent.Value)
	}
	return nil
}

// ViolatesRight checks that the argument Min is a valid candidate for a Right child of the accessed Min.
func (min *Min) ViolatesRight(m *Min) error {

	if m.Value < min.Value {
		return fmt.Errorf("address (*%p) cannot hold pointer (*%p). argument struct for (*%p).Right must contain value greater than %f", &m, &min, &min, min.Value)
	}
	return nil
}

// ViolatesSide checks that the argument Side is a supported Min Side.
func (min *Min) ViolatesSide(side string) error {

	_, ok := Sides[side]

	if ok == false {
		return fmt.Errorf("unsupported string value. argument side must be either %s, %s or %s", LEFT, RIGHT, ROOT)
	}
	return nil
}

var _ min = (*Min)(nil)

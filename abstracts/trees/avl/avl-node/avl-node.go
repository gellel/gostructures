package node

import (
	"fmt"
	"log"
	"math"
)

const (
	LEFT  string = "LEFT"
	RIGHT string = "RIGHT"
)

var (
	sides = map[string]bool{LEFT: true, RIGHT: true}
)

type avl interface {
	AssignLeft(a *AVL) *AVL
	AssignParent(a *AVL) *AVL
	AssignRight(a *AVL) *AVL
	Balance() int
	HasLeft() bool
	HasParent() bool
	HasRight() bool
	Height() int
	HeightOf() float64
	HeightLeft() float64
	HeightRight() float64
	Insert(value float64) *AVL
	IsChild(b *AVL) bool
	IsEqual(value float64) bool
	IsLess(value float64) bool
	IsMore(value float64) bool
	Remove(value float64) *AVL
	RemoveLeft() *AVL
	RemoveRight() *AVL
	RotateLeft() *AVL
	RotateRight() *AVL
	RotateLeftRight() *AVL
	RotateRightLeft() *AVL
	ToAVLSlice() []*AVL
	ToFloatSlice() []float64
	UnsafelyAssignLeft(b *AVL) *AVL
	UnsafelyAssignRight(b *AVL) *AVL
	ViolatesLeft(b *AVL) error
	ViolatesRight(b *AVL) error
	Walk()
}

type AVL struct {
	Left   *AVL
	Parent *AVL
	Right  *AVL
	Side   string
	Value  float64
}

func New(value float64) *AVL {
	return &AVL{Value: value}
}

func (avl *AVL) AssignLeft(a *AVL) *AVL {

	err := avl.ViolatesLeft(a)

	if err != nil {
		log.Panicln(err)
	}

	return avl.UnsafelyAssignLeft(a)
}

func (avl *AVL) AssignParent(a *AVL) *AVL {

	err := avl.ViolatesParent(a)

	if err != nil {
		log.Panicln(err)
	}

	return avl.UnsafelyAssignParent(a)
}

func (avl *AVL) AssignRight(a *AVL) *AVL {

	err := avl.ViolatesRight(a)

	if err != nil {
		log.Panicln(err)
	}

	return avl.UnsafelyAssignRight(a)
}

func (avl *AVL) AssignSide(side string) *AVL {

	err := avl.ViolatesSide(side)

	if err != nil {
		log.Panicln(err)
	}

	avl.Side = side

	return avl
}

func (avl *AVL) Balance() int {
	return int(avl.HeightLeft() - avl.HeightRight())
}

func (avl *AVL) HasLeft() bool {
	return avl.Left != nil
}

func (avl *AVL) HasParent() bool {
	return avl.Parent != nil
}

func (avl *AVL) HasRight() bool {
	return avl.Right != nil
}

func (avl *AVL) Height() int {
	return int(avl.HeightOf())
}

func (avl *AVL) HeightOf() float64 {
	return math.Max(avl.HeightLeft(), avl.HeightRight())
}

func (avl *AVL) HeightLeft() float64 {
	if avl.HasLeft() {
		return (avl.Left.HeightOf() + 1.0)
	}
	return 0.0
}

func (avl *AVL) HeightRight() float64 {
	if avl.HasRight() {
		return (avl.Right.HeightOf() + 1.0)
	}
	return 0.0
}

func (avl *AVL) IsChild(a *AVL) bool {
	return avl.Left == a || avl.Right == a
}

func (avl *AVL) IsEqual(value float64) bool {
	return avl.Value == value
}

func (avl *AVL) IsLess(value float64) bool {
	return value < avl.Value
}

func (avl *AVL) IsMore(value float64) bool {
	return value > avl.Value
}

func (avl *AVL) IsLeft() bool {
	return avl.Side == LEFT
}

func (avl *AVL) IsRight() bool {
	return avl.Side == RIGHT
}

func (avl *AVL) RemoveLeft() *AVL {

	avl.Left = nil

	return avl
}

func (avl *AVL) RemoveParent() *AVL {

	avl.Parent = nil

	return avl
}

func (avl *AVL) RemoveRight() *AVL {

	avl.Right = nil

	return avl
}

func (avl *AVL) ToAVLSlice() []*AVL {

	nodes := make([]*AVL, 0)

	if avl.HasLeft() {
		nodes = append(nodes, avl.Left.ToAVLSlice()...)
	}
	nodes = append(nodes, avl)
	if avl.HasRight() {
		nodes = append(nodes, avl.Right.ToAVLSlice()...)
	}
	return nodes
}

func (avl *AVL) ToFloatSlice() []float64 {

	floats := make([]float64, 0)

	if avl.HasLeft() {
		floats = append(floats, avl.Left.ToFloatSlice()...)
	}
	floats = append(floats, avl.Value)
	if avl.HasRight() {
		floats = append(floats, avl.Right.ToFloatSlice()...)
	}
	return floats
}

func (avl *AVL) UnsafelyAssignLeft(a *AVL) *AVL {
	avl.Left = a.UnsafelyAssignParent(avl)
	return avl
}

func (avl *AVL) UnsafelyAssignParent(a *AVL) *AVL {
	avl.Parent = a
	return avl
}

func (avl *AVL) UnsafelyAssignRight(a *AVL) *AVL {
	avl.Right = a.UnsafelyAssignParent(avl)
	return avl
}

func (avl *AVL) ViolatesLeft(a *AVL) error {
	if a.Value > avl.Value {
		return fmt.Errorf("address (*%p) cannot hold pointer (*%p). argument struct must contain value less than %f", &avl, &a, avl.Value)
	}
	return nil
}

func (avl *AVL) ViolatesParent(a *AVL) error {
	if avl.Parent == nil {
		return fmt.Errorf("address (*%p) has no parent pointer.", &avl)
	} else if avl.Side == LEFT && avl.Parent.Value < avl.Value {
		return fmt.Errorf("address (*%p) cannot hold pointer (*%p). argument struct value must exceed %f", &avl, &a, avl.Parent.Value)
	} else if avl.Side == RIGHT && avl.Parent.Value > avl.Value {
		return fmt.Errorf("address (*%p) cannot hold pointer (*%p). argument struct must not exceed %f", &avl, &a, avl.Parent.Value)
	}
	return nil
}

func (avl *AVL) ViolatesRight(a *AVL) error {
	if a.Value < avl.Value {
		return fmt.Errorf("address (*%p) cannot hold pointer (*%p). argument struct must contain value greater than %f", &avl, &a, avl.Value)
	}
	return nil
}

func (avl *AVL) ViolatesSide(side string) error {

	_, ok := sides[side]

	if ok == false {
		return fmt.Errorf("unsupported string value. argument side must be either %s or %s", LEFT, RIGHT)
	}
	return nil
}

var _ avl = (*AVL)(nil)

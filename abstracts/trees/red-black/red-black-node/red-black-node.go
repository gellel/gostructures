package node

import "log"

const (
	BLACK = "BLACK"
	LEFT  = "LEFT"
	RED   = "RED"
	RIGHT = "RIGHT"
	ROOT  = "ROOT"
)

var (
	colors = map[string]bool{
		BLACK: true,
		RED:   true}
	sides = map[string]bool{
		LEFT:  true,
		RIGHT: true}
)

type redBlack interface {
	AssignColor(color string) *RedBlack
	AssignLeft(r *RedBlack) *RedBlack
	AssignParent(r *RedBlack) *RedBlack
	AssignRight(r *RedBlack) *RedBlack
	AssignSide(side string) *RedBlack
	EmptyColor() bool
	EmptyLeft() bool
	EmptyParent() bool
	EmptyRight() bool
	EmptySide() bool
	Insert(value float64) *RedBlack
	IsBlack() bool
	IsEqual(value float64) bool
	IsLess(value float64) bool
	IsLeft() bool
	IsMore(value float64) bool
	IsRed() bool
	IsRight() bool
	IsRoot() bool
	Remove(value float64) *RedBlack
	RemoveLeft() *RedBlack
	RemoveParent() *RedBlack
	RemoveRight() *RedBlack
	ToRedBlackSlice() []*RedBlack
	ToFloatSlice() []float64
	UnsafelyAssignLeft(a *RedBlack) *RedBlack
	UnsafelyAssignParent(a *RedBlack) *RedBlack
	UnsafelyAssignRight(a *RedBlack) *RedBlack
	UnsafelyAssignSide(side string) *RedBlack
	ViolatesColor(color string) error
	ViolatesLeft(a *RedBlack) error
	ViolatesParent(a *RedBlack) error
	ViolatesRight(a *RedBlack) error
	ViolatesSide(side string) error
	Walk() *RedBlack
}

type RedBlack struct {
	Color  string
	Left   *RedBlack
	Parent *RedBlack
	Right  *RedBlack
	Side   string
	Value  float64
}

func (redBlack *RedBlack) AssignColor(color string) *RedBlack {

	err := redBlack.ViolatesColor(color)

	if err != nil {
		log.Panicln(err)
	}

	return redBlack.UnsafelyAssignColor(color)
}

func (redBlack *RedBlack) AssignLeft(r *RedBlack) *RedBlack {

	err := redBlack.ViolatesLeft(r)

	if err != nil {
		log.Panicln(err)
	}

	return redBlack.UnsafelyAssignLeft(r)
}

func (redBlack *RedBlack) AssignParent(r *RedBlack) *RedBlack {

	err := redBlack.ViolatesParent()

	if err != nil {
		log.Panicln(err)
	}

	return redBlack.UnsafelyAssignParent(r)
}

func (redBlack *RedBlack) AssignRight(r *RedBlack) *RedBlack {

	err := redBlack.ViolatesRight(r)

	if err != nil {
		log.Panicln(err)
	}

	return redBlack.UnsafelyAssignRight(r)
}

func (redBlack *RedBlack) AssignSide(side string) *RedBlack {

	err := redBlack.ViolatesSide(side)

	if err != nil {
		log.Panicln(err)
	}

	return redBlack.UnsafelyAssignSide(side)
}

func (redBlack *RedBlack) EmptyColor() bool {
	return redBlack.Color == ""
}

func (redBlack *RedBlack) EmptyLeft() bool {
	return redBlack.Left == nil
}

func (redBlack *RedBlack) EmptyParent() bool {
	return redBlack.Parent == nil
}

func (redBlack *RedBlack) EmptyRight() bool {
	return redBlack.Right == nil
}

func (redBlack *RedBlack) EmptySide() bool {
	return redBlack.Side == ""
}

func (redBlack *RedBlack) Insert(value float64) *RedBlack {
	return redBlack
}

func (redBlack *RedBlack) IsBlack() bool {
	return redBlack.Color == BLACK
}

func (redBlack *RedBlack) IsEqual(value float64) bool {
	return value == redBlack.Value
}

func (redBlack *RedBlack) IsLess(value float64) bool {
	return value < redBlack.Value
}

func (redBlack *RedBlack) IsLeft() bool {
	return redBlack.Side == LEFT
}

func (redBlack *RedBlack) IsMore(value float64) bool {
	return value > redBlack.Value
}

func (redBlack *RedBlack) IsRed() bool {
	return redBlack.Color == RED
}

func (redBlack *RedBlack) IsRight() bool {
	return redBlack.Side == RIGHT
}

func (redBlack *RedBlack) IsRoot() bool {
	return redBlack.Side == ROOT
}

func (redBlack *RedBlack) Remove(value float64) *RedBlack {
	return redBlack
}

var _ redBlack = (*RedBlack)(nil)

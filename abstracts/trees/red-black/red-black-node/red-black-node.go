package node

import (
	"fmt"
	"log"
)

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
        RIGHT: true,
        ROOT:  true}
)

type Rb interface {
    AssignBlack() *RedBlack
	AssignColor(color string) *RedBlack
	AssignLeft(r *RedBlack) *RedBlack
    AssignParent(r *RedBlack) *RedBlack
    AssignRed() *RedBlack
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
	MaxValue() float64
	MinValue() float64
	Remove(value float64) *RedBlack
	RemoveLeft() *RedBlack
	RemoveParent() *RedBlack
	RemoveRight() *RedBlack
	ToRedBlackSlice() []*RedBlack
	ToFloatSlice() []float64
	UnsafelyAssignLeft(r *RedBlack) *RedBlack
	UnsafelyAssignParent(r *RedBlack) *RedBlack
	UnsafelyAssignRight(r *RedBlack) *RedBlack
	UnsafelyAssignSide(side string) *RedBlack
	ViolatesColor(color string) error
	ViolatesLeft(r *RedBlack) error
	ViolatesParent(r *RedBlack) error
	ViolatesRight(r *RedBlack) error
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

func (redBlack *RedBlack) AssignBlack() *RedBlack {
    return redBlack.AssignColor(BLACK)
}

func (redBlack *RedBlack) AssignColor(color string) *RedBlack {

	err := redBlack.ViolatesColor(color)

	if err != nil {
		log.Panicln(err)
	}

	return redBlack.UnsafelyAssignColor(color)
}

func (redBlack *RedBlack) AssignLeft(rb *RedBlack) *RedBlack {

	err := redBlack.ViolatesLeft(rb)

	if err != nil {
		log.Panicln(err)
	}

	return redBlack.UnsafelyAssignLeft(rb)
}

func (redBlack *RedBlack) AssignParent(rb *RedBlack) *RedBlack {

	err := redBlack.ViolatesParent(rb)

	if err != nil {
		log.Panicln(err)
	}

	return redBlack.UnsafelyAssignParent(rb)
}

func (redBlack *RedBlack) AssignRed() *RedBlack {
    return redBlack.AssignColor(RED)
}

func (redBlack *RedBlack) AssignRight(rb *RedBlack) *RedBlack {

	err := redBlack.ViolatesRight(rb)

	if err != nil {
		log.Panicln(err)
	}

	return redBlack.UnsafelyAssignRight(rb)
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

func (redBlack *RedBlack) HasLeft() bool {
	return redBlack.Left != nil
}

func (redBlack *RedBlack) HasParent() bool {
	return redBlack.Parent != nil
}

func (redBlack *RedBlack) HasRight() bool {
	return redBlack.Right != nil
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

func (redBlack *RedBlack) MaxValue() float64 {

	r := redBlack

	for r != nil {
		r = r.Right
	}
	return r.Value
}

func (redBlack *RedBlack) MinValue() float64 {

	r := redBlack

	for r != nil {
		r = r.Left
	}
	return r.Value
}

func (redBlack *RedBlack) Remove(value float64) *RedBlack {
	return redBlack
}

func (redBlack *RedBlack) RemoveLeft() *RedBlack {

	redBlack.Left = nil

	return redBlack
}

func (redBlack *RedBlack) RemoveParent() *RedBlack {

	redBlack.Parent = nil

	return redBlack
}

func (redBlack *RedBlack) RemoveRight() *RedBlack {

	redBlack.Right = nil

	return redBlack
}

func (redBlack *RedBlack) ToFloatSlice() []float64 {

	floats := make([]float64, 0)

	if redBlack.HasLeft() {
		floats = append(floats, redBlack.Left.ToFloatSlice()...)
	}
	floats = append(floats, redBlack.Value)
	if redBlack.HasRight() {
		floats = append(floats, redBlack.Right.ToFloatSlice()...)
	}
	return floats
}

func (redBlack *RedBlack) ToRedBlackSlice() []*RedBlack {

	nodes := make([]*RedBlack, 0)

	if redBlack.HasLeft() {
		nodes = append(nodes, redBlack.Left.ToRedBlackSlice()...)
	}
	nodes = append(nodes, redBlack)
	if redBlack.HasRight() {
		nodes = append(nodes, redBlack.Right.ToRedBlackSlice()...)
	}
	return nodes
}

func (redBlack *RedBlack) UnsafelyAssignColor(color string) *RedBlack {
	redBlack.Color = color
	return redBlack
}

func (redBlack *RedBlack) UnsafelyAssignLeft(rb *RedBlack) *RedBlack {
	redBlack.Left = rb.UnsafelyAssignSide(LEFT)
	return redBlack
}

func (redBlack *RedBlack) UnsafelyAssignParent(rb *RedBlack) *RedBlack {
	redBlack.Parent = rb
	return redBlack
}

func (redBlack *RedBlack) UnsafelyAssignRight(rb *RedBlack) *RedBlack {
	redBlack.Right = rb.UnsafelyAssignSide(RIGHT)
	return redBlack
}

func (redBlack *RedBlack) UnsafelyAssignSide(side string) *RedBlack {
	redBlack.Side = side
	return redBlack
}

func (redBlack *RedBlack) ViolatesColor(color string) error {

	_, ok := colors[color]

	if ok == false {
		return fmt.Errorf("unsupported string value. argument color must be either %s or %s", RED, BLACK)
	}
	return nil
}

func (redBlack *RedBlack) ViolatesLeft(rb *RedBlack) error {
	if rb.Value > redBlack.Value {
		return fmt.Errorf("address (*%p) cannot hold pointer (*%p). argument struct must contain value less than %f", &redBlack, &rb, redBlack.Value)
	}
	return nil
}

func (redBlack *RedBlack) ViolatesParent(rb *RedBlack) error {
	if redBlack.Parent == nil {
		return fmt.Errorf("address (*%p) has no parent pointer", &redBlack)
	} else if redBlack.Side == LEFT && redBlack.Parent.Value < redBlack.Value {
		return fmt.Errorf("address (*%p) cannot hold pointer (*%p). argument struct value must exceed %f", &redBlack, &rb, redBlack.Parent.Value)
	} else if redBlack.Side == RIGHT && redBlack.Parent.Value > redBlack.Value {
		return fmt.Errorf("address (*%p) cannot hold pointer (*%p). argument struct must not exceed %f", &redBlack, &rb, redBlack.Parent.Value)
	}
	return nil
}

func (redBlack *RedBlack) ViolatesRight(rb *RedBlack) error {
	if rb.Value < redBlack.Value {
		return fmt.Errorf("address (*%p) cannot hold pointer (*%p). argument struct must contain value greater than %f", &redBlack, &rb, redBlack.Value)
	}
	return nil
}

func (redBlack *RedBlack) ViolatesSide(side string) error {

	_, ok := sides[side]

	if ok == false {
		return fmt.Errorf("unsupported string value. argument side must be either %s, %s or %s", LEFT, RIGHT, ROOT)
	}
	return nil
}

func (redBlack *RedBlack) Walk() *RedBlack {
	if redBlack.HasLeft() {
		redBlack.Left.Walk()
	}
	log.Println(redBlack.Value, redBlack.Side)
	if redBlack.HasRight() {
		redBlack.Right.Walk()
	}
	return redBlack
}

var _ Rb = (*RedBlack)(nil)

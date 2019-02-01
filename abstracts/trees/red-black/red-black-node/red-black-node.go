package node

import (
	"fmt"
	"log"
)

const (
	// BLACK declares Rb.Color as BLACK.
	BLACK = "BLACK"
	// LEFT declares Rb.Side as LEFT.
	LEFT = "LEFT"
	// RED declares Rb.Color as RED.
	RED = "RED"
	// RIGHT declares Rb.Side as RIGHT.
	RIGHT = "RIGHT"
	// ROOT declares Rb.Side as ROOT.
	ROOT = "ROOT"
)

var (
	// Colors declares map of Rb.Color colors.
	Colors = map[string]bool{
		BLACK: true,
		RED:   true}
	// Sides declares map of Rb.Side sides.
	Sides = map[string]bool{
		LEFT:  true,
		RIGHT: true,
		ROOT:  true}
)

// Rb declares the interface for a RedBlack-Tree-Node.
type Rb interface {
	About() *RedBlack
	AssignBlack() *RedBlack
	AssignColor(color string) *RedBlack
	AssignLeft(rb *RedBlack) *RedBlack
	AssignParent(rb *RedBlack) *RedBlack
	AssignRed() *RedBlack
	AssignRight(rb *RedBlack) *RedBlack
	AssignSide(side string) *RedBlack
	EmptyAdjacent() bool
	EmptyAncestors() bool
	EmptyChildren() bool
	EmptyColor() bool
	EmptyGrandParent() bool
	EmptyLeft() bool
	EmptyParent() bool
	EmptyRight() bool
	EmptySide() bool
	EmptyUncle() bool
	Find(value float64) *RedBlack
	HasAdjacent() bool
	HasAncestors() bool
	HasChildren() bool
	HasColor() bool
	HasGrandParent() bool
	HasLeft() bool
	HasParent() bool
	HasRight() bool
	HasSide() bool
	HasUncle() bool
	Insert(value float64) *RedBlack
	InsertBinary(rb *RedBlack) *RedBlack
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
	Relate() *RedBlack
	Remove(value float64) *RedBlack
	RemoveAncestors() *RedBlack
	RemoveLeft() *RedBlack
	RemoveParent() *RedBlack
	RemoveRight() *RedBlack
	RemoveUncle() *RedBlack
	Rotate() *RedBlack
	RotateLeft() *RedBlack
	RotateRight() *RedBlack
	ToRedBlackSlice() []*RedBlack
	ToFloatSlice() []float64
	UnsafelyAssignColor(color string) *RedBlack
	UnsafelyAssignGrandParent(rb *RedBlack) *RedBlack
	UnsafelyAssignLeft(rb *RedBlack) *RedBlack
	UnsafelyAssignParent(rb *RedBlack) *RedBlack
	UnsafelyAssignRight(rb *RedBlack) *RedBlack
	UnsafelyAssignSide(side string) *RedBlack
	ViolatesColor(color string) error
	ViolatesLeft(rb *RedBlack) error
	ViolatesParent(rb *RedBlack) error
	ViolatesRight(rb *RedBlack) error
	ViolatesSide(side string) error
	Walk() *RedBlack
}

// RedBlack declares the struct for a RedBlack-Tree-Node.
type RedBlack struct {
	Color       string
	GrandParent *RedBlack
	Left        *RedBlack
	Parent      *RedBlack
	Right       *RedBlack
	Side        string
	Uncle       *RedBlack
	Value       float64
}

// New instantiates a new Rb pointer.
func New(value float64) *RedBlack {
	return &RedBlack{
		Color: BLACK,
		Side:  ROOT,
		Value: value}
}

// About displays plain details about accessed Rb.
func (redBlack *RedBlack) About() *RedBlack {

	uncle := -0.0

	grandparent := -0.0

	parent := -0.0

	left := -0.0

	right := -0.0

	if redBlack.HasUncle() {
		uncle = redBlack.Uncle.Value
	}
	if redBlack.HasGrandParent() {
		grandparent = redBlack.GrandParent.Value
	}
	if redBlack.HasParent() {
		parent = redBlack.Parent.Value
	}
	if redBlack.HasLeft() {
		left = redBlack.Left.Value
	}
	if redBlack.HasRight() {
		right = redBlack.Right.Value
	}

	message := "(Color: %s, Side: %s) Left: %f. This: %f. Right: %f. Parent: %f. GrandParent: %f. Uncle: %f"

	s := fmt.Sprintf(message, redBlack.Side, redBlack.Color, left, redBlack.Value, right, parent, grandparent, uncle)

	fmt.Println(s)

	return redBlack
}

// AssignBlack sets the accessed Rb.Color as BLACK.
func (redBlack *RedBlack) AssignBlack() *RedBlack {
	return redBlack.AssignColor(BLACK)
}

// AssignColor safely assigns a supported color to the accessed Rb. Enforces that the argument color is RED or BLACK.
func (redBlack *RedBlack) AssignColor(color string) *RedBlack {

	err := redBlack.ViolatesColor(color)

	if err != nil {
		log.Panicln(err)
	}
	return redBlack.UnsafelyAssignColor(color)
}

// AssignLeft safely assigns a Left to the accessed Rb. Enforces that the argument Rb is of a lesser value.
func (redBlack *RedBlack) AssignLeft(rb *RedBlack) *RedBlack {

	err := redBlack.ViolatesLeft(rb)

	if err != nil {
		log.Panicln(err)
	}
	return redBlack.UnsafelyAssignLeft(rb)
}

// AssignParent safely assigns a Parent to the accessed Rb. Enforces that the argument Rb is greater than accessed Rb when Rb is a Left child.
// When Rb is a Right child AssignParent enforces that the argument Rb is less than the accessed Rb.
func (redBlack *RedBlack) AssignParent(rb *RedBlack) *RedBlack {

	err := redBlack.ViolatesParent(rb)

	if err != nil {
		log.Panicln(err)
	}
	return redBlack.UnsafelyAssignParent(rb)
}

// AssignRed sets the accessed Rb.Color as RED.
func (redBlack *RedBlack) AssignRed() *RedBlack {
	return redBlack.AssignColor(RED)
}

// AssignRight safely assigns a Right to the accessed Rb. Enforces that the argument Rb is of a greater value.
func (redBlack *RedBlack) AssignRight(rb *RedBlack) *RedBlack {

	err := redBlack.ViolatesRight(rb)

	if err != nil {
		log.Panicln(err)
	}
	return redBlack.UnsafelyAssignRight(rb)
}

// AssignSide safely assigns a Side to the accessed Rb. Enforces that the argument Side is either LEFT, RIGHT or ROOT.
func (redBlack *RedBlack) AssignSide(side string) *RedBlack {

	err := redBlack.ViolatesSide(side)

	if err != nil {
		log.Panicln(err)
	}
	return redBlack.UnsafelyAssignSide(side)
}

// EmptyColor checks that the assigned Rb.Color to the accessed Rb is empty.
func (redBlack *RedBlack) EmptyColor() bool {
	return redBlack.Color == ""
}

// EmptyAdjacent checks that the accessed Rb has no relationships (Rb.Parent, Rb.GrandParent, Rb.Left, Rb.Right)
func (redBlack *RedBlack) EmptyAdjacent() bool {
	return redBlack.EmptyChildren() && redBlack.EmptyAncestors()
}

// EmptyAncestors checks that the accessed Rb has no Parent or GrandParent.
func (redBlack *RedBlack) EmptyAncestors() bool {
	return redBlack.EmptyParent() && redBlack.EmptyGrandParent()
}

// EmptyChildren checks that the accessed Rb has no Children.
func (redBlack *RedBlack) EmptyChildren() bool {
	return redBlack.EmptyLeft() && redBlack.EmptyRight()
}

// EmptyGrandParent checks that the assigned Rb.GrandParent to the accessed Rb is empty.
func (redBlack *RedBlack) EmptyGrandParent() bool {
	return redBlack.GrandParent == nil
}

// EmptyLeft checks that the assigned Rb.Left to the accessed Rb is empty.
func (redBlack *RedBlack) EmptyLeft() bool {
	return redBlack.Left == nil
}

// EmptyParent checks that the assigned Rb.Parent to the accessed Rb is empty.
func (redBlack *RedBlack) EmptyParent() bool {
	return redBlack.Parent == nil
}

// EmptyRight checks that the assigned Rb.Right to the accessed Rb is empty.
func (redBlack *RedBlack) EmptyRight() bool {
	return redBlack.Right == nil
}

// EmptySide checks that the assigned Rb.Side to the accessed Rb is empty.
func (redBlack *RedBlack) EmptySide() bool {
	return redBlack.Side == ""
}

// EmptyUncle checks that the assigned Rb.Uncle to the accessed Rb is empty.
func (redBlack *RedBlack) EmptyUncle() bool {
	return redBlack.Uncle == nil
}

// Find searches for a Rb that holds the argument value.
func (redBlack *RedBlack) Find(value float64) *RedBlack {

	if redBlack.IsEqual(value) {
		return redBlack
	} else if redBlack.IsLess(value) && redBlack.HasLeft() {
		return redBlack.Left.Find(value)
	} else if redBlack.IsMore(value) && redBlack.HasRight() {
		return redBlack.Right.Find(value)
	}
	return nil
}

// HasAdjacent checks that the accessed Rb contains relationships (Rb.Parent, Rb.GrandParent, Rb.Left, Rb.Right).
func (redBlack *RedBlack) HasAdjacent() bool {
	return redBlack.HasChildren() && redBlack.HasAncestors()
}

// HasAncestors checks that the assigned Rb.Parent and Rb.GrandParent to the accessed Rb is not Nil.
func (redBlack *RedBlack) HasAncestors() bool {
	return redBlack.HasParent() && redBlack.HasGrandParent()
}

// HasChildren checks that the assigned Rb.Left and assigned Rb.Right to the accessed Rb is not Nil.
func (redBlack *RedBlack) HasChildren() bool {
	return redBlack.HasLeft() && redBlack.HasRight()
}

// HasColor checks that the assigned Rb.Color to the accessed Rb is not Nil.
func (redBlack *RedBlack) HasColor() bool {
	return redBlack.Color != ""
}

// HasGrandParent checks that the assigned Rb.GrandParent to the accessed Rb is not Nil.
func (redBlack *RedBlack) HasGrandParent() bool {
	return redBlack.GrandParent != nil
}

// HasLeft checks that the assigned Rb.Left to the accessed Rb is not Nil.
func (redBlack *RedBlack) HasLeft() bool {
	return redBlack.Left != nil
}

// HasParent checks that the assigned Rb.Parent to the accessed Rb is not Nil.
func (redBlack *RedBlack) HasParent() bool {
	return redBlack.Parent != nil
}

// HasRight checks that the assigned Rb.Right to the accessed Rb is not Nil.
func (redBlack *RedBlack) HasRight() bool {
	return redBlack.Right != nil
}

// HasSide checks that the assigned Rb.Side to the accessed Rb is not Nil.
func (redBlack *RedBlack) HasSide() bool {
	return redBlack.Side != ""
}

// HasUncle checks that the assigned Rb.Uncle to the accessed Rb is not Nil.
func (redBlack *RedBlack) HasUncle() bool {
	return redBlack.Uncle != nil
}

// Insert creates and adds new Rb to the accessed Rb (or its descendants) when the value is not already contained in the Rb.
func (redBlack *RedBlack) Insert(value float64) *RedBlack {

	rb := redBlack.InsertBinary(New(value).AssignRed())

	rb.Rotate()

	return redBlack
}

// InsertBinary assigns a new Rb to the accessed Rb (or its descendants) after a performing a Binary Search and returns the assigned Rb.
func (redBlack *RedBlack) InsertBinary(rb *RedBlack) *RedBlack {

	if redBlack.IsLess(rb.Value) {
		if redBlack.EmptyLeft() {
			redBlack.UnsafelyAssignLeft(rb)
		} else {
			redBlack.Left.InsertBinary(rb)
		}
	} else if redBlack.IsMore(rb.Value) {
		if redBlack.EmptyRight() {
			redBlack.UnsafelyAssignRight(rb)
		} else {
			redBlack.Right.InsertBinary(rb)
		}
	}
	return rb
}

// IsBlack checks that the accessed Rb is colored BLACK.
func (redBlack *RedBlack) IsBlack() bool {
	return redBlack.Color == BLACK
}

// IsEqual checks that the Rb.Value is greater than the Value of the accessed Rb.
func (redBlack *RedBlack) IsEqual(value float64) bool {
	return value == redBlack.Value
}

// IsLess checks that the Rb.Value is less than the Value of the accessed Rb.
func (redBlack *RedBlack) IsLess(value float64) bool {
	return value < redBlack.Value
}

// IsLeft checks that the accessed Rb is a Left child of its assigned Rb.Parent.
func (redBlack *RedBlack) IsLeft() bool {
	return redBlack.Side == LEFT
}

// IsMore checks that the Rb.Value is greater than the Value of the accessed Rb.
func (redBlack *RedBlack) IsMore(value float64) bool {
	return value > redBlack.Value
}

// IsRed checks that the accessed Rb is colored RED.
func (redBlack *RedBlack) IsRed() bool {
	return redBlack.Color == RED
}

// IsRight checks that the accessed Rb is a Right child of its assigned Rb.Parent.
func (redBlack *RedBlack) IsRight() bool {
	return redBlack.Side == RIGHT
}

// IsRoot checks that the accessed Rb is the Root of all Rb ancestors.
func (redBlack *RedBlack) IsRoot() bool {
	return redBlack.Side == ROOT
}

// MaxValue finds the largest Value stored in the descendants of the accessed Rb.
func (redBlack *RedBlack) MaxValue() float64 {

	r := redBlack

	for r != nil {
		r = r.Right
	}
	return r.Value
}

// MinValue finds the smallest Value stored in the descendants of the accessed Rb.
func (redBlack *RedBlack) MinValue() float64 {

	r := redBlack

	for r != nil {
		r = r.Left
	}
	return r.Value
}

// Relate iterates and sets the relationships of the accessed Rb descendents.
func (redBlack *RedBlack) Relate() *RedBlack {

	if redBlack.HasLeft() {
		redBlack.AssignLeft(redBlack.Left.Relate())
	}
	if redBlack.HasRight() {
		redBlack.AssignRight(redBlack.Right.Relate())
	}
	return redBlack
}

// Remove deletes the accessed Rb or one of its descendants if Rb.Value matches the argument Value.
func (redBlack *RedBlack) Remove(value float64) *RedBlack {
	return redBlack
}

// RemoveAncestors removes the references to the accessed Rb's Parent, GrandParent and Uncle.
func (redBlack *RedBlack) RemoveAncestors() *RedBlack {

	redBlack.Parent = nil

	redBlack.GrandParent = nil

	redBlack.Uncle = nil

	return redBlack
}

// RemoveGrandParent removes the reference to the GrandParent of the accessed Rb.
func (redBlack *RedBlack) RemoveGrandParent() *RedBlack {

	redBlack.GrandParent = nil

	return redBlack
}

// RemoveLeft removes the reference to the Left child of the accessed Rb.
func (redBlack *RedBlack) RemoveLeft() *RedBlack {

	redBlack.Left = nil

	return redBlack
}

// RemoveParent removes the reference to the Parent of the accessed Rb.
func (redBlack *RedBlack) RemoveParent() *RedBlack {

	redBlack.Parent = nil

	return redBlack
}

// RemoveRight removes the reference to the Right child of the accessed Rb.
func (redBlack *RedBlack) RemoveRight() *RedBlack {

	redBlack.Right = nil

	return redBlack
}

// RemoveUncle removes the reference to the Uncle of the accessed Rb.
func (redBlack *RedBlack) RemoveUncle() *RedBlack {

	redBlack.Uncle = nil

	return redBlack
}

func (redBlack *RedBlack) Rotate() *RedBlack {

	if redBlack.IsRoot() && redBlack.IsBlack() {
		return redBlack
	}
	if redBlack.IsRed() && redBlack.Parent.IsBlack() {
		return redBlack
	}
	if redBlack.HasUncle() && redBlack.Uncle.IsRed() {

		redBlack.AssignBlack()

		redBlack.Uncle.AssignRed()

		redBlack.Parent.AssignRed()

		return redBlack.GrandParent.Rotate()
	}

	if redBlack.HasUncle() && redBlack.Uncle.IsBlack() {

	}
	/*
		if redBlack.IsRoot() {
			return redBlack
		}
		if redBlack.Parent.IsBlack() {
			return redBlack
		}
		if redBlack.HasUncle() && redBlack.Uncle.IsRed() {

			redBlack.Uncle.AssignBlack()

			redBlack.Parent.AssignBlack()

			if redBlack.GrandParent.IsRoot() {
				return redBlack
			}

			redBlack.GrandParent.AssignRed()

			redBlack.GrandParent.Rotate()
		}
	*/

	return redBlack
}

// RotateLeft changes the accessed Rb address. Moves Rb.Right address as Rb address. Changes Rb.Left address to accessed Rb address.
func (redBlack *RedBlack) RotateLeft() *RedBlack {

	root := *redBlack // eg. 5

	left := redBlack.Left // eg. 2

	right := redBlack.Right // eg.10

	if root.HasParent() {
		right.AssignParent(root.Parent)
	} else {
		right.RemoveAncestors()
	}

	*redBlack = *right // now 10.

	redBlack.AssignSide(root.Side)

	redBlack.AssignLeft(&root) // now 5

	redBlack.AssignRight(redBlack.Right)

	root.AssignLeft(left)

	x := root.Right.Left

	root.AssignRight(x) // then update relationships as these will be wrong

	x.Relate()

	return redBlack
}

func (redBlack *RedBlack) RotateRight() *RedBlack {

	return redBlack
}

// ToFloatSlice creates a slice of all Rb.Value's stored at the accessed Rb.
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

// ToRedBlackSlice creates a slice of all Rb's stored at the accessed Rb.
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

// UnsafelyAssignColor assigns a Color to the accessed Rb without performing validity checks.
func (redBlack *RedBlack) UnsafelyAssignColor(color string) *RedBlack {

	redBlack.Color = color

	return redBlack
}

// UnsafelyAssignGrandParent assigns a GrandParent to the accessed Rb without performing validity checks.
func (redBlack *RedBlack) UnsafelyAssignGrandParent(rb *RedBlack) *RedBlack {

	redBlack.GrandParent = rb

	if rb.HasChildren() {
		redBlack.UnsafelyAssignUncle(rb)
	} else {
		redBlack.RemoveUncle()
	}
	return redBlack
}

// UnsafelyAssignLeft assigns a Left child to the accessed Rb without performing validity checks.
func (redBlack *RedBlack) UnsafelyAssignLeft(rb *RedBlack) *RedBlack {

	rb.UnsafelyAssignSide(LEFT)

	rb.UnsafelyAssignParent(redBlack)

	redBlack.Left = rb

	return redBlack
}

// UnsafelyAssignParent assigns a Parent to the accessed Rb without performing validity checks.
func (redBlack *RedBlack) UnsafelyAssignParent(rb *RedBlack) *RedBlack {

	redBlack.Parent = rb

	if rb.HasParent() {
		redBlack.UnsafelyAssignGrandParent(rb.Parent)
	} else {
		redBlack.RemoveGrandParent().RemoveUncle()
	}
	return redBlack
}

// UnsafelyAssignRight assigns a Right child to the accessed Rb without performing validity checks.
func (redBlack *RedBlack) UnsafelyAssignRight(rb *RedBlack) *RedBlack {

	rb.UnsafelyAssignSide(RIGHT)

	rb.UnsafelyAssignParent(redBlack)

	redBlack.Right = rb

	return redBlack
}

// UnsafelyAssignSide assigns a Side to the accessed Rb without performing validity checks.
func (redBlack *RedBlack) UnsafelyAssignSide(side string) *RedBlack {

	redBlack.Side = side

	return redBlack
}

// UnsafelyAssignUncle assigns a Uncle to the accessed Rb without performing validity checks.
func (redBlack *RedBlack) UnsafelyAssignUncle(rb *RedBlack) *RedBlack {

	if redBlack.IsLeft() {
		if redBlack.Parent.Value == rb.Right.Value {
			redBlack.Uncle = rb.Left
		} else {
			redBlack.Uncle = rb.Right
		}
	} else if redBlack.IsRight() {
		if redBlack.Parent.Value == rb.Left.Value {
			redBlack.Uncle = rb.Right
		} else {
			redBlack.Uncle = rb.Left
		}
	}
	return redBlack
}

// ViolatesColor checks that the argument Color is a supported Rb Color.
func (redBlack *RedBlack) ViolatesColor(color string) error {

	_, ok := Colors[color]

	if ok == false {
		return fmt.Errorf("unsupported string value. argument color must be either %s or %s", RED, BLACK)
	}
	return nil
}

// ViolatesLeft checks that the argument Rb is a valid candidate for a Left child of the accessed Rb.
func (redBlack *RedBlack) ViolatesLeft(rb *RedBlack) error {

	if rb.Value > redBlack.Value {
		return fmt.Errorf("address (*%p) cannot hold pointer (*%p). argument struct must contain value less than %f", &redBlack, &rb, redBlack.Value)
	}
	return nil
}

// ViolatesParent checks that the argument Rb is a valid candidate for the Parent of the accessed Rb.
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

// ViolatesRight checks that the argument Rb is a valid candidate for a Right child of the accessed Rb.
func (redBlack *RedBlack) ViolatesRight(rb *RedBlack) error {

	if rb.Value < redBlack.Value {
		return fmt.Errorf("address (*%p) cannot hold pointer (*%p). argument struct must contain value greater than %f", &redBlack, &rb, redBlack.Value)
	}
	return nil
}

// ViolatesSide checks that the argument Side is a supported Rb Side.
func (redBlack *RedBlack) ViolatesSide(side string) error {

	_, ok := Sides[side]

	if ok == false {
		return fmt.Errorf("unsupported string value. argument side must be either %s, %s or %s", LEFT, RIGHT, ROOT)
	}
	return nil
}

// Walk visits all of the descendants of the accessed Rb in an ordered fashion, logging the Rb.Value at each interval once all sequential options are exhausted.
func (redBlack *RedBlack) Walk() *RedBlack {

	if redBlack.HasLeft() {
		redBlack.Left.Walk()
	}
	log.Println(redBlack.Value, redBlack.Side, redBlack.Color)
	if redBlack.HasRight() {
		redBlack.Right.Walk()
	}
	return redBlack
}

var _ Rb = (*RedBlack)(nil)

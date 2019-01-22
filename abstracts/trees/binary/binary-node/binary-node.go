// Package node exports a Binary-Tree-Node for a Binary-Search-Tree.
// Binary-Tree-Node implements most of the core functionality that the
// Binary-Search-Tree provides but at a Binary-Tree-Node level. Package
// also exports a function "New" that simplifies instantiating
// Binary-Tree-Node pointers. Leaf-node's are float64's for min-max
// comparison.
package node

import (
	"fmt"
	"log"
	"math"
)

const (
	// LEFT declares constant string that provides a namespace to annotate that
	// a Binary-node is assigned as a left child of its parent.
	LEFT string = "LEFT"
	// RIGHT declares constant string that provides a namespace to annotate that
	// a Binary-node is assigned as a right child of its parent.
	RIGHT string = "RIGHT"
)

var (
	// SIDES declares constant map that conatins supported child-node positions.
	sides = map[string]bool{LEFT: true, RIGHT: true}
)

// Binary declares the implementation for a Binary-Tree-Node
// within the Binary-Search-Tree.
type binary interface {
	AssignLeft(b *Binary) *Binary
	AssignRight(b *Binary) *Binary
	AssignSide(side string) *Binary
	EmptyLeft() bool
	EmptyRight() bool
	Find(value float64) *Binary
	HasLeft() bool
	HasRight() bool
	Height() int
	HeightOf() float64
	HeightLeft() float64
	HeightRight() float64
	Insert(value float64) *Binary
	IsChild(b *Binary) bool
	IsEqual(value float64) bool
	IsLess(value float64) bool
	IsMore(value float64) bool
	Remove(value float64) *Binary
	RemoveLeft() *Binary
	RemoveRight() *Binary
	ToBinarySlice() []*Binary
	ToFloatSlice() []float64
	UnsafelyAssignLeft(b *Binary) *Binary
	UnsafelyAssignRight(b *Binary) *Binary
	ViolatesLeft(b *Binary) error
	ViolatesRight(b *Binary) error
	Walk()
}

// Binary declares the struct for a Leaf-node within the Binary-Search-Tree.
// Each Leaf-node contains a floating point value for describing it's assigned
// numerical value. This was chosen to enable greater precision for more
// complex (in terms of numerical complexity) Binary-Search-Tree's. In addition,
// a Leaf-node struct contains an optional left and right Leaf-node, which acts
// as the branches of the Binary-Search-Tree. Each Binary-Tree-Node that is created
// using the Leaf-node's insert method will have it's literal side position
// assigned to the new child-node as a string vale of "left" or "right".
type Binary struct {
	Left  *Binary // Left is the left Binary-Tree-Node of the accessed Binary-Tree-Node.
	Right *Binary // Right is the right Binary-Tree-Node of the accessed Left-Node.
	Side  string  // Side expresses the position (relative to the parent) that the accessed Binary-Tree-Node is assigned.
	Value float64 // Value is the numeric weight of the Binary-Tree-Node. In a Binary-Search-Tree, a value lower than the accessed Binary-Tree-Node will be assigned to the left position. Expectidly, the opposite is true for larger sums.
}

// New instantiates a new Leaf-node. A Leaf-nodes generated using the New
// function do not require a reference to either a left or right child-node.
// Additionally, the side reference is not required.
func New(value float64) *Binary {
	return &Binary{Value: value}
}

// AssignLeft assigns by pointer a left child-node to the accessed Binary-Tree-Node. Method
// expects that the provided Binary-Tree-Node pointer contains a value that is of a lesser
// value than that of the accessed Binary-Tree-Node. If a violation occurs method throws
// an invalid argument exception.
func (binary *Binary) AssignLeft(b *Binary) *Binary {

	err := binary.ViolatesLeft(b)

	if err != nil {
		log.Panicln(err)
	}

	return binary.UnsafelyAssignLeft(b)
}

// AssignRight assigns by pointer a right child-node to the accessed Binary-Tree-Node. Method
// expects that the provided Binary-Tree-Node pointer contains a value that is of a greater
// value than that of the accessed Binary-Tree-Node. If a violation occurs method throws
// an invalid argument exception.
func (binary *Binary) AssignRight(b *Binary) *Binary {

	err := binary.ViolatesRight(b)

	if err != nil {
		log.Panicln(err)
	}

	return binary.UnsafelyAssignRight(b)
}

// AssignSide assigns the accessed Binary-Tree-Node a string value that expresses
// the position it sits within its parent Binary-Tree-Node. A Binary-Tree-Node that is
// stored in Binary.Left must contain a string value of "LEFT". Expectidly,
// a Binary-Tree-Node stored in Binary.Right must contain a string value of "RIGHT".
func (binary *Binary) AssignSide(side string) *Binary {

	err := binary.ViolatesSide(side)

	if err != nil {
		log.Panicln(err)
	}

	binary.Side = side

	return binary
}

// EmptyLeft checks whether the accessed Binary-Tree-Node has an unassigned Left-Child-Node.
func (binary *Binary) EmptyLeft() bool {
	return binary.Left == nil
}

// EmptyRight checks whether the accessed Binary-Tree-Node has an unassigned Right-Child-Node.
func (binary *Binary) EmptyRight() bool {
	return binary.Right == nil
}

// Find checks whether the accessed Binary-Tree-Node or its Child-Nodes contains the argument value.
func (binary *Binary) Find(value float64) *Binary {
	if binary.IsEqual(value) {
		return binary
	} else if binary.IsLess(value) && binary.HasLeft() {
		return binary.Left.Find(value)
	} else if binary.IsMore(value) && binary.HasRight() {
		return binary.Right.Find(value)
	}
	return nil
}

// HasLeft checks whether the accessed Binary-Tree-Node has an assigned Left-Child-Node.
func (binary *Binary) HasLeft() bool {
	return binary.Left != nil
}

// HasRight checks whether the accessed Binary-Tree-Node has an assigned Right-Child-Node.
func (binary *Binary) HasRight() bool {
	return binary.Right != nil
}

// Height computes the current depth of the accessed Binary-Tree-Node and converts the
// floating point sum to an integer.
func (binary *Binary) Height() int {
	return int(binary.HeightOf())
}

// HeightOf computes the current depth of the accessed Binary-Tree-Node.
func (binary *Binary) HeightOf() float64 {
	return math.Max(binary.HeightLeft(), binary.HeightRight())
}

// HeightLeft computes the nested depth of the accessed Binary-Tree-Node's Left-Children.
func (binary *Binary) HeightLeft() float64 {
	if binary.HasLeft() {
		return (binary.Left.HeightOf() + 1.0)
	}
	return 0.0
}

// HeightRight computes the nested depth of the accessed Binary-Tree-Node's Right-Children.
func (binary *Binary) HeightRight() float64 {
	if binary.HasRight() {
		return (binary.Right.HeightOf() + 1.0)
	}
	return 0.0
}

// Insert creates and assigns a new Binary-Tree-Node to the accessed Binary-Tree-Node.
// When an argument value is smaller than the accessed Binary-Tree-Node value, the new instance
// is stored on the accessed Binary-Tree-Node's left position. Alternatively, when
// a larger value is provided, it is stored than the right position. Equal sums are
// discarded and no new Binary-Tree-Node is created.
func (binary *Binary) Insert(value float64) *Binary {
	if binary.IsLess(value) {
		if binary.HasLeft() {
			binary.Left.Insert(value)
		} else {
			binary.UnsafelyAssignLeft(New(value))
		}
	} else if binary.IsMore(value) {
		if binary.HasRight() {
			binary.Right.Insert(value)
		} else {
			binary.UnsafelyAssignRight(New(value))
		}
	}
	return binary
}

// IsChild checks whether the pointer address is mapped to the accessed Binary-Tree-Node's
// left or right Child-Node positions.
func (binary *Binary) IsChild(b *Binary) bool {
	return binary.Left == b || binary.Right == b
}

// IsEqual checks whether the argument value is equal to the accessed
// Binary-Tree-Node's assigned value.
func (binary *Binary) IsEqual(value float64) bool {
	return binary.Value == value
}

// IsLess checks whether the argument value is less than the accessed
// Binary-Tree-Node's assigned value.
func (binary *Binary) IsLess(value float64) bool {
	return binary.Value < value
}

// IsMore checks whether the argument value is greater than the accessed
// Binary-Tree-Node's assigned value.
func (binary *Binary) IsMore(value float64) bool {
	return binary.Value > value
}

// IsLeft checks whether the accessed Binary-Tree-Node is assigned to
// another Binary-Tree-Node's left position.
func (binary *Binary) IsLeft() bool {
	return binary.Side == LEFT
}

// IsRight checks whether the accessed Binary-Tree-Node is assigned to
// another Binary-Tree-Node's right position.
func (binary *Binary) IsRight() bool {
	return binary.Side == RIGHT
}

// Remove deletes a child-node contained within the accessed Binary-Tree-Node
// or within the accessed Binary-Tree-Node's current child-nodes.
func (binary *Binary) Remove(value float64) *Binary {
	if binary.IsLess(value) && binary.HasLeft() {
		if binary.Left.IsEqual(value) {
			binary.RemoveLeft()
		} else {
			binary.Left.Remove(value)
		}
	} else if binary.IsMore(value) && binary.HasRight() {
		if binary.Right.IsEqual(value) {
			binary.RemoveRight()
		} else {
			binary.Right.Remove(value)
		}
	}
	return binary
}

// RemoveLeft unassigns the accessed Binary-Tree-Node's left Binary-Tree-Node.
func (binary *Binary) RemoveLeft() *Binary {

	binary.Left = nil

	return binary
}

// RemoveRight unassigns the accessed Binary-Tree-Node's right Binary-Tree-Node.
func (binary *Binary) RemoveRight() *Binary {

	binary.Right = nil

	return binary
}

// ToBinarySlice creates an ordered slice of the assigned Binary-Tree-Node's child-nodes.
func (binary *Binary) ToBinarySlice() []*Binary {

	binaries := make([]*Binary, 0)

	if binary.HasLeft() {
		binaries = append(binaries, binary.Left.ToBinarySlice()...)
	}
	binaries = append(binaries, binary)
	if binary.HasRight() {
		binaries = append(binaries, binary.Right.ToBinarySlice()...)
	}
	return binaries
}

// ToFloatSlice creates an ordered slice of the assigned Binary-Tree-Node's child-nodes values.
func (binary *Binary) ToFloatSlice() []float64 {

	floats := make([]float64, 0)

	if binary.HasLeft() {
		floats = append(floats, binary.Left.ToFloatSlice()...)
	}
	floats = append(floats, binary.Value)
	if binary.HasRight() {
		floats = append(floats, binary.Right.ToFloatSlice()...)
	}
	return floats
}

// UnsafelyAssignLeft sets the argument pointer to the accessed
// Binary-Tree-Node left child position without performing
// the Binary-Search-Tree insert operation.
func (binary *Binary) UnsafelyAssignLeft(b *Binary) *Binary {

	binary.Left = b.AssignSide(LEFT)

	return binary
}

// UnsafelyAssignRight sets the argument pointer to the accessed
// Binary-Tree-Node right child position without performing
// the Binary-Search-Tree insert operation.
func (binary *Binary) UnsafelyAssignRight(b *Binary) *Binary {

	binary.Right = b.AssignSide(RIGHT)

	return binary
}

// ViolatesLeft checks that the argument Binary-Tree-Node pointer contains
// a floating point value is greater than the accessed Binary-Tree-Node.
func (binary *Binary) ViolatesLeft(b *Binary) error {
	if b.Value > binary.Value {
		return fmt.Errorf("address (*%p) cannot hold pointer (*%p). argument struct must contain value less than %f", &binary, &b, binary.Value)
	}
	return nil
}

// ViolatesRight checks that the argument Binary-Tree-Node pointer contains
// a floating point value is less than the accessed Binary-Tree-Node.
func (binary *Binary) ViolatesRight(b *Binary) error {
	if b.Value < binary.Value {
		return fmt.Errorf("address (*%p) cannot hold pointer (*%p). argument struct must contain value greater than %f", &binary, &b, binary.Value)
	}
	return nil
}

// ViolatesSide checks that the argument value is not either "LEFT" or "RIGHT".
func (binary *Binary) ViolatesSide(side string) error {

	_, ok := sides[side]

	if ok == false {
		return fmt.Errorf("unsupported string value. argument side must be either %s or %s", LEFT, RIGHT)
	}
	return nil
}

// Walk accesses all assigned child-nodes to accessed Binary-Tree-Node.
func (binary *Binary) Walk() {
	if binary.HasLeft() {
		binary.Left.Walk()
	}
	log.Println(binary.Value, binary.Side)
	if binary.HasRight() {
		binary.Right.Walk()
	}
}

var _ binary = (*Binary)(nil)

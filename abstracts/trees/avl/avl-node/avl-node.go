// Package node exports a AVL-Tree-Node for a AVL-Search-Tree.
// AVL-Tree-Node implements most of the core functionality that the
// AVL-Search-Tree provides but at a AVL-Tree-Node level. Package
// also exports a function "New" that simplifies instantiating
// AVL-Tree-Node pointers. Leaf-Node's are float64's for min-max
// comparison.
package node

import (
	"fmt"
	"log"
	"math"
)

const (
	// LEFT declares constant string that provides a namespace to annotate that
	// a AVL-node is assigned as a left child of its parent.
	LEFT string = "LEFT"
	// RIGHT declares constant string that provides a namespace to annotate that
	// a AVL-node is assigned as a right child of its parent.
	RIGHT string = "RIGHT"
)

var (
	// SIDES declares constant map that conatins supported child-node positions.
	sides = map[string]bool{LEFT: true, RIGHT: true}
)

// AVL declares the implementation for a AVL-Tree-Node
// within the AVL-Search-Tree.
type avl interface {
	AssignLeft(a *AVL) *AVL
	AssignParent(a *AVL) *AVL
	AssignRight(a *AVL) *AVL
	AssignSide(side string) *AVL
	Balance() int
	EmptyLeft() bool
	EmptyParent() bool
	EmptyRight() bool
	EmptySide() bool
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
	MaxValue() float64
	MinValue() float64
	Remove(value float64) *AVL
	RemoveLeft() *AVL
	RemoveRight() *AVL
	Rotate(value float64) *AVL
	RotateLeft() *AVL
	RotateRight() *AVL
	RotateLeftRight() *AVL
	RotateRightLeft() *AVL
	ToAVLSlice() []*AVL
	ToFloatSlice() []float64
	UnsafelyAssignLeft(a *AVL) *AVL
	UnsafelyAssignParent(a *AVL) *AVL
	UnsafelyAssignRight(a *AVL) *AVL
	UnsafelyAssignSide(side string) *AVL
	ViolatesLeft(a *AVL) error
	ViolatesParent(a *AVL) error
	ViolatesRight(a *AVL) error
	ViolatesSide(side string) error
	Walk() *AVL
}

// AVL declares the struct for a Leaf-Node within the AVL-Search-Tree.
// Each Leaf-Node contains a floating point value for describing it's assigned
// numerical value. This was chosen to enable greater precision for more
// complex (in terms of numerical complexity) AVL-Search-Tree's. In addition,
// a Leaf-Node struct contains an optional left and right Leaf-Node, which acts
// as the branches of the AVL-Search-Tree. Each AVL-Tree-Node that is created
// using the Leaf-Node's insert method will have it's literal side position
// assigned to the new child-node as a string vale of "left" or "right".
// As new Leaf-Node's are inserted into an accessed AVL-Node,
// an AVL rotation may occur to satisfy the properties of an AVL-Tree.
type AVL struct {
	Left   *AVL    // Left is the left AVL-Tree-Node of the accessed AVL-Tree-Node.
	Parent *AVL    // Parent is the AVL-Tree-Node that holds the accessed AVL-Tree-Node in either the left or right position.
	Right  *AVL    // Right is the right AVL-Tree-Node of the accessed Left-Node.
	Side   string  // Side expresses the position (relative to the parent) that the accessed AVL-Tree-Node is assigned.
	Value  float64 // Value is the numeric weight of the AVL-Tree-Node. In a AVL-Search-Tree, a value lower than the accessed AVL-Tree-Node will be assigned to the left position. Expectidly, the opposite is true for larger sums.
}

// New instantiates a new Leaf-node. A Leaf-nodes generated using the New
// function do not require a reference to either a left or right child-node.
// Additionally, the side or parent reference is not required.
func New(value float64) *AVL {
	return &AVL{Value: value}
}

// AssignLeft assigns by pointer a left child-node to the accessed AVL-Tree-Node. Method
// expects that the provided AVL-Tree-Node pointer contains a value that is of a lesser
// value than that of the accessed AVL-Tree-Node. If a violation occurs method throws
// an invalid argument exception.
func (avl *AVL) AssignLeft(a *AVL) *AVL {

	err := avl.ViolatesLeft(a)

	if err != nil {
		log.Panicln(err)
	}

	return avl.UnsafelyAssignLeft(a)
}

// AssignParent assigns by pointer a parent AVL-Tree-Node to the accessed AVL-Tree-Node. Method
// expects that the provided AVL-Tree-Node pointer contains a value that is of a greater
// value than that of the accessed AVL-Tree-Node. If a violation occurs method throws
// an invalid argument exception.
func (avl *AVL) AssignParent(a *AVL) *AVL {

	err := avl.ViolatesParent(a)

	if err != nil {
		log.Panicln(err)
	}

	return avl.UnsafelyAssignParent(a)
}

// AssignRight assigns by pointer a right child-node to the accessed AVL-Tree-Node. Method
// expects that the provided AVL-Tree-Node pointer contains a value that is of a greater
// value than that of the accessed AVL-Tree-Node. If a violation occurs method throws
// an invalid argument exception.
func (avl *AVL) AssignRight(a *AVL) *AVL {

	err := avl.ViolatesRight(a)

	if err != nil {
		log.Panicln(err)
	}

	return avl.UnsafelyAssignRight(a)
}

// AssignSide assigns the accessed AVL-Tree-Node a string value that expresses
// the position it sits within its parent AVL-Tree-Node. A AVL-Tree-Node that is
// stored in AVL.Left must contain a string value of "LEFT". Expectidly,
// a AVL-Tree-Node stored in AVL.Right must contain a string value of "RIGHT".
func (avl *AVL) AssignSide(side string) *AVL {

	err := avl.ViolatesSide(side)

	if err != nil {
		log.Panicln(err)
	}

	avl.Side = side

	return avl
}

// Balance computes the weighting of the accessed AVL-Tree-Node. Negative integers
// suggest that the AVL-Tree-Node holds more child-nodes on the right of the of
// node. A positive value suggests the left contains more left children.
func (avl *AVL) Balance() int {
	return int(avl.HeightLeft() - avl.HeightRight())
}

// EmptyLeft checks whether the accessed AVL-Tree-Node has an unassigned Left-Child-Node.
func (avl *AVL) EmptyLeft() bool {
	return avl.Left == nil
}

// EmptyParent checks whether the accessed AVL-Tree-Node has an unassigned Parent-Node.
func (avl *AVL) EmptyParent() bool {
	return avl.Parent == nil
}

// EmptyRight checks whether the accessed AVL-Tree-Node has an unassigned Right-Child-Node.
func (avl *AVL) EmptyRight() bool {
	return avl.Right == nil
}

// EmptySide checks whether the accessed AVL-Tree-Node has an unassigned side reference.
func (avl *AVL) EmptySide() bool {
	return avl.Side == ""
}

// Find checks whether the accessed AVL-Tree-Node or its Child-Nodes contains the argument value.
func (avl *AVL) Find(value float64) *AVL {
	if avl.IsEqual(value) {
		return avl
	} else if avl.IsLess(value) && avl.HasLeft() {
		return avl.Left.Find(value)
	} else if avl.IsMore(value) && avl.HasRight() {
		return avl.Right.Find(value)
	}
	return nil
}

// HasLeft checks whether the accessed AVL-Tree-Node has an assigned Left-Child-Node.
func (avl *AVL) HasLeft() bool {
	return avl.Left != nil
}

// HasParent checks whether the accessed AVL-Tree-Node has an assigned Parent-Node.
func (avl *AVL) HasParent() bool {
	return avl.Parent != nil
}

// HasRight checks whether the accessed AVL-Tree-Node has an assigned Right-Child-Node.
func (avl *AVL) HasRight() bool {
	return avl.Right != nil
}

// Height computes the current depth of the accessed AVL-Tree-Node and converts the
// floating point sum to an integer.
func (avl *AVL) Height() int {
	return int(avl.HeightOf())
}

// HeightOf computes the current depth of the accessed AVL-Tree-Node.
func (avl *AVL) HeightOf() float64 {
	return math.Max(avl.HeightLeft(), avl.HeightRight())
}

// HeightLeft computes the nested depth of the accessed AVL-Tree-Node's Left-Children.
func (avl *AVL) HeightLeft() float64 {
	if avl.HasLeft() {
		return (avl.Left.HeightOf() + 1.0)
	}
	return 0.0
}

// HeightRight computes the nested depth of the accessed AVL-Tree-Node's Right-Children.
func (avl *AVL) HeightRight() float64 {
	if avl.HasRight() {
		return (avl.Right.HeightOf() + 1.0)
	}
	return 0.0
}

// Insert creates and assigns a new AVL-Tree-Node to the accessed AVL-Tree-Node.
// When an argument value is smaller than the accessed AVL-Tree-Node value, the new instance
// is stored on the accessed AVL-Tree-Node's left position. Alternatively, when
// a larger value is provided, it is stored than the right position. Equal sums are
// discarded and no new AVL-Tree-Node is created. When an insertion is performed,
// accessed AVL-Tree-Node checks that the it is evenly distributed. When bias left,
// the child nodes at the left position are shuffled one over so that they are weighted right
// evenly. The alternative case moves nodes from the right to the left.
func (avl *AVL) Insert(value float64) *AVL {
	if avl.IsEqual(value) {
		return avl
	}
	if avl.IsLess(value) {
		if avl.HasLeft() {
			avl.Left.Insert(value)
		} else {
			avl.UnsafelyAssignLeft(New(value))
		}
	} else if avl.IsMore(value) {
		if avl.HasRight() {
			avl.Right.Insert(value)
		} else {
			avl.UnsafelyAssignRight(New(value))
		}
	}
	return avl.Rotate(value)
}

// IsChild checks whether the pointer address is mapped to the accessed AVL-Tree-Node's
// left or right Child-Node positions.
func (avl *AVL) IsChild(a *AVL) bool {
	return avl.Left == a || avl.Right == a
}

// IsEqual checks whether the argument value is equal to the accessed
// AVL-Tree-Node's assigned value.
func (avl *AVL) IsEqual(value float64) bool {
	return avl.Value == value
}

// IsLess checks whether the argument value is less than the accessed
// AVL-Tree-Node's assigned value.
func (avl *AVL) IsLess(value float64) bool {
	return value < avl.Value
}

// IsMore checks whether the argument value is greater than the accessed
// AVL-Tree-Node's assigned value.
func (avl *AVL) IsMore(value float64) bool {
	return value > avl.Value
}

// IsLeft checks whether the accessed AVL-Tree-Node is assigned to
// another AVL-Tree-Node's left position.
func (avl *AVL) IsLeft() bool {
	return avl.Side == LEFT
}

// IsRight checks whether the accessed AVL-Tree-Node is assigned to
// another AVL-Tree-Node's right position.
func (avl *AVL) IsRight() bool {
	return avl.Side == RIGHT
}

// MaxValue finds the largest value from within the accessed AVL-Tree-Node.
func (avl *AVL) MaxValue() float64 {

	a := avl

	for a != nil {
		a = a.Right
	}
	return a.Value
}

// MinValue finds the smallest value from within the accessed AVL-Tree-Node.
func (avl *AVL) MinValue() float64 {

	a := avl

	for a != nil {
		a = a.Left
	}
	return a.Value
}

// Remove deletes a child-node contained within the accessed AVL-Tree-Node
// or within the accessed AVL-Tree-Node's current child-nodes.
func (avl *AVL) Remove(value float64) *AVL {
	if avl.IsLess(value) && avl.HasLeft() {
		return avl.Left.Remove(value)
	}
	if avl.IsMore(value) && avl.HasRight() {
		return avl.Right.Remove(value)
	}
	if avl.IsEqual(value) && avl.IsLeft() && avl.HasParent() {
		return avl.Parent.RemoveLeft().Rotate(value)
	}
	if avl.IsEqual(value) && avl.IsRight() && avl.HasParent() {
		return avl.Parent.RemoveRight().Rotate(value)
	}
	return avl
}

// RemoveLeft unassigns the accessed AVL-Tree-Node's left AVL-Tree-Node.
func (avl *AVL) RemoveLeft() *AVL {

	avl.Left = nil

	return avl
}

// RemoveParent unassigns the accessed AVL-Tree-Node's parent AVL-Tree-Node.
func (avl *AVL) RemoveParent() *AVL {

	avl.Parent = nil

	return avl
}

// RemoveRight unassigns the accessed AVL-Tree-Node's right AVL-Tree-Node.
func (avl *AVL) RemoveRight() *AVL {

	avl.Right = nil

	return avl
}

// Rotate performs the AVL-Tree-Node balancing rotations.
func (avl *AVL) Rotate(value float64) *AVL {

	balance := avl.Balance()

	if (balance > 1) && value < avl.Left.Value {
		return avl.RotateRight()
	} else if (balance < -1) && value > avl.Right.Value {
		return avl.RotateLeft()
	} else if (balance > 1) && value > avl.Left.Value {
		return avl.RotateLeftRight()
	} else if (balance < -1) && value < avl.Right.Value {
		return avl.RotateRightLeft()
	}
	return avl
}

// RotateLeft shifts the accessed AVL-Tree-Node to the right child position
// of the accessed AVL-Tree-Node's left child. This effective makes the
// the following operation; before rotation Z(root.Left.Left)<-Y(root.Left)<-X(root);
// after rotation Z(root.Left)<-Y(root)->(root.Right)X.
func (avl *AVL) RotateLeft() *AVL {

	root := *avl

	root.RemoveRight()

	avl.Right.AssignParent(root.Parent)

	*avl = *avl.Right

	avl.UnsafelyAssignSide(root.Side)

	if avl.HasLeft() {
		root.AssignRight(avl.Left)
	}

	avl.AssignLeft(&root)

	return avl
}

// RotateLeftRight performs a combination of AVL.LeftRotate and
// AVL.RightRotate. The first operation requires the accessed AVL-Tree-Node
// to rotate its left-child left. The accessed AVL-Tree-Node then
// rotates itself right.
func (avl *AVL) RotateLeftRight() *AVL {

	avl.Left.RotateLeft()

	return avl.RotateRight()
}

// RotateRight shifts the accessed AVL-Tree-Node to the left child position
// of the accessed AVL-Tree-Node's right child. This effective makes the
// the following operation; before rotation X(root)->Y(root.Right)->Z(root.Right.Right);
// after rotation X(root.Left)<-Y(root)->(root.Right)Z.
func (avl *AVL) RotateRight() *AVL {

	root := *avl

	root.RemoveLeft()

	avl.Left.AssignParent(root.Parent)

	*avl = *avl.Left

	avl.UnsafelyAssignSide(root.Side)

	if avl.HasRight() {
		root.AssignLeft(avl.Right)
	}

	avl.AssignRight(&root)

	return avl
}

// RotateRightLeft performs a combination of AVL.RotateRight and
// AVL.RotateLeft. The first operation requires the accessed AVL-Tree-Node
// to rotate its right-child right. The accessed AVL-Tree-Node then
// rotates itself left.
func (avl *AVL) RotateRightLeft() *AVL {

	avl.Right.RotateRight()

	return avl.RotateLeft()
}

// ToAVLSlice creates an ordered slice of the assigned AVL-Tree-Node's child-nodes.
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

// ToFloatSlice creates an ordered slice of the assigned AVL-Tree-Node's child-nodes values.
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

// UnsafelyAssignLeft sets the argument pointer to the accessed
// AVL-Tree-Node left child position without performing
// the AVL-Search-Tree insert operation.
func (avl *AVL) UnsafelyAssignLeft(a *AVL) *AVL {
	avl.Left = a.UnsafelyAssignParent(avl).AssignSide(LEFT)
	return avl
}

// UnsafelyAssignParent sets the argument pointer to the accessed
// AVL-Tree-Node parent position without performing
// the AVL-Search-Tree insert operation.
func (avl *AVL) UnsafelyAssignParent(a *AVL) *AVL {
	avl.Parent = a
	return avl
}

// UnsafelyAssignRight sets the argument pointer to the accessed
// AVL-Tree-Node right child position without performing
// the AVL-Search-Tree insert operation.
func (avl *AVL) UnsafelyAssignRight(a *AVL) *AVL {
	avl.Right = a.UnsafelyAssignParent(avl).AssignSide(RIGHT)
	return avl
}

// UnsafelyAssignSide sets the argument string to the accessed
// AVL-Tree-Node sid reference without validating the argument string.
func (avl *AVL) UnsafelyAssignSide(side string) *AVL {
	avl.Side = side
	return avl
}

// ViolatesLeft checks that the argument AVL-Tree-Node pointer contains
// a floating point value is greater than the accessed AVL-Tree-Node.
func (avl *AVL) ViolatesLeft(a *AVL) error {
	if a.Value > avl.Value {
		return fmt.Errorf("address (*%p) cannot hold pointer (*%p). argument struct must contain value less than %f", &avl, &a, avl.Value)
	}
	return nil
}

// ViolatesParent checks that the argument AVL-Tree-Node pointer
// satisfies the left and right constraints of a AVL-Search-Tree;
// enforcing that the parent of a left node is greater than the
// accessed AVL-Tree-Node. The alternative check is performed for
// a right node.
func (avl *AVL) ViolatesParent(a *AVL) error {
	if avl.Parent == nil {
		return fmt.Errorf("address (*%p) has no parent pointer", &avl)
	} else if avl.Side == LEFT && avl.Parent.Value < avl.Value {
		return fmt.Errorf("address (*%p) cannot hold pointer (*%p). argument struct value must exceed %f", &avl, &a, avl.Parent.Value)
	} else if avl.Side == RIGHT && avl.Parent.Value > avl.Value {
		return fmt.Errorf("address (*%p) cannot hold pointer (*%p). argument struct must not exceed %f", &avl, &a, avl.Parent.Value)
	}
	return nil
}

// ViolatesRight checks that the argument AVL-Tree-Node pointer contains
// a floating point value is less than the accessed AVL-Tree-Node.
func (avl *AVL) ViolatesRight(a *AVL) error {
	if a.Value < avl.Value {
		return fmt.Errorf("address (*%p) cannot hold pointer (*%p). argument struct must contain value greater than %f", &avl, &a, avl.Value)
	}
	return nil
}

// ViolatesSide checks that the argument value is not either "LEFT" or "RIGHT".
func (avl *AVL) ViolatesSide(side string) error {

	_, ok := sides[side]

	if ok == false {
		return fmt.Errorf("unsupported string value. argument side must be either %s or %s", LEFT, RIGHT)
	}
	return nil
}

// Walk accesses all assigned child-nodes to accessed AVL-Tree-Node.
func (avl *AVL) Walk() *AVL {
	if avl.HasLeft() {
		avl.Left.Walk()
	}
	log.Println(avl.Value, avl.Side)
	if avl.HasRight() {
		avl.Right.Walk()
	}
	return avl
}

var _ avl = (*AVL)(nil)

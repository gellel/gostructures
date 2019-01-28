package node

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
	AssignLeft(r *RedBlack) *RedBlack
	AssignParent(r *RedBlack) *RedBlack
	AssignRight(r *RedBlack) *RedBlack
	AssignSide(side string) *RedBlack
	EmptyLeft() bool
	EmptyParent() bool
	EmptyRight() bool
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
	Left   *RedBlack
	Parent *RedBlack
	Right  *RedBlack
	Side   string
	Value  float64
}

var _ redBlack = (*RedBlack)(nil)

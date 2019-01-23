package node

const (
	// LEFT declares constant string reference for AVL-Leaf left position.
	LEFT string = "LEFT"
	// RIGHT declares constant string reference for AVL-Leaf right position.
	RIGHT string = "RIGHT"
)

var (
	// SIDES declares constant map that conatins supported child-node positions.
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
	ToBinarySlice() []*AVL
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
		return fmt.Errorf("address (*%p) cannot hold pointer (*%p). argument struct value must exceed %f", &avl, %a, avl.Parent.Value)
	} else if avl.Side == RIGHT && avl.Parent.Value > avl.Value {
		return fmt.Errorf()
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

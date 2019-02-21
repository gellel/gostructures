package container

type container interface {
	Access(p int) interface{}
	AccessLeftOf(p int) interface{}
	AccessParentOf(p int) interface{}
	AccessRightOf(p int) interface{}
	Bounds(p int) bool
	BoundsLeftOf(p int) bool
	BoundsParentOf(p int) bool
	BoundsRightOf(p int) bool
	CanMerge(c *Container) bool
	CannotMerge(c *Container) bool
	Contains(value interface{}) bool
	CutAfter(p int) *Container
	CutBefore(p int) *Container
	Empty() *Container
	IsEmpty() bool
	IsNotEmpty() bool
	LeftOf(p int) int
	Length() int
	Merge(c *Container) *Container
	ParentOf(p int) int
	Peek() interface{}
	PeekAt(p int) interface{}
	PeekAtLeft(p int) interface{}
	PeekAtParent(p int) interface{}
	PeekAtRight(p int) interface{}
	PeekLast() interface{}
	Poll() interface{}
	Push(value interface{}) int
	RightOf(p int) int
	Search(value interface{}) int
	Swap(a int, b int) *Container
	ToSlice() []interface{}
	ToTypeOf(value interface{}) []interface{}
	Type()
	TypeOf(value interface{}) string
}

type Container []interface{}

func (container *Container) Access(p int) interface{} {
	return (*container)[p]
}

func (container *Container) AccessLeftOf(p int) interface{} {
	return container.Access(container.LeftOf(p))
}

func (container *Container) AccessParentOf(p int) interface{} {
	return container.Access(container.ParentOf(p))
}

func (container *Container) AccessRightOf(p int) interface{} {
	return container.Access(container.RightOf(p))
}

func (container *Container) Bounds(p int) bool {
	return ((p > -1) && (p < (len((*container)) - 1)))
}

func (container *Container) BoundsLeftOf(p int) bool {
	return container.Bounds(container.LeftOf(p))
}

func (container *Container) BoundsParentOf(p int) bool {
	return container.Bounds(container.ParentOf(p))
}

func (container *Container) BoundsRightOf(p int) bool {
	return container.Bounds(container.RightOf(p))
}

func (container *Container) CanMerge(c *Container) bool {
}

func (container *Container) CannotMerge(c *Container) bool {
}

func (container *Container) Contains(value interface{}) bool {
}

func (container *Container) CutAfter(p int) *Container {
	*container = (*container)[p:]
	return container
}

func (container *Container) CutBefore(p int) *Container {
	*container = (*container)[:p]
	return container
}

func (container *Container) Empty() *Container {
	*container = (*container)[:0]
	return container
}

func (container *Container) IsEmpty() bool {
	return (len((*container)) == 0)
}

func (container *Container) IsNotEmpty() bool {
	return (len((*container)) != 0)
}

func (container *Container) LeftOf(p int) int {
	return ((p * 2) + 1)
}

func (container *Container) Length() int {
	return len((*container))
}

func (container *Container) Merge(c *Container) *Container {
}

func (container *Container) ParentOf(p int) int {
	return ((p - 1) / 2)
}

func (container *Container) Peek() interface{} {
	if container.IsNotEmpty() {
		return container.Access(0)
	}
	return nil
}

func (container *Container) PeekAt(p int) interface{} {
	if container.Bounds(p) {
		return container.Access(p)
	}
	return nil
}

func (container *Container) PeekAtLeft(p int) interface{} {
	if container.BoundsLeftOf(p) {
		return container.AccessLeftOf(p)
	}
	return nil
}

func (container *Container) PeekAtParent(p int) interface{} {
	if container.BoundsParentOf(p) {
		return container.AccessParentOf(p)
	}
	return nil
}

func (container *Container) PeekAtRight(p int) interface{} {
	if container.BoundsRightOf(p) {
		return container.AccessRightOf(p)
	}
	return nil
}

func (container *Container) PeekLast() interface{} {
	if container.IsNotEmpty() {
		return container.PeekAt(container.Length() - 1)
	}
	return nil
}

func (container *Container) Poll() interface{} {
	if container.IsNotEmpty() {
		value := container.Access(0)
		container.CutAfter(0)
		return value
	}
	return nil
}

func (container *Container) Push(value interface{}) int {
	*container = append((*container), value)
	return container.Length()
}

var _ container = (*Container)(nil)

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
	PeekFirst() interface{}
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
	return container.Access(container.RIghtOf(p))
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

var _ container = (*Container)(nil)

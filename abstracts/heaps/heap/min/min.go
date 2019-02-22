package min

import (
	"github.com/gellel/gostructures/abstracts/heaps/heap/container"
	"github.com/gellel/gostructures/abstracts/heaps/heap/core"
)

type heap interface {
	BubbleDown(p int) int
	BubbleUp(p int) int
	Push(value interface{}) int
}

type Heap struct {
	core.Heap
}

func New() *Heap {
	h := &Heap{}
	h.Container = container.New()
	return h
}

func (heap *Heap) BubbleDown(p int) int {

}

func (heap *Heap) BubbleUp(p int) int {

}

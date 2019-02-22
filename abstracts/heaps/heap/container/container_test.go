package container_test

import (
	"fmt"
	"testing"

	"github.com/gellel/gostructures/abstracts/heaps/heap/container"
)

func Test(t *testing.T) {

	a := container.New()

	b := container.New()

	a.Push(1.0)

	b.Push(1)

	fmt.Println(a.TypeOf(a.Peek()))

	fmt.Println(b.TypeOf(b.Peek()))

	fmt.Println(b.CanMerge(a))
}

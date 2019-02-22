package core_test

import (
	"fmt"
	"testing"

	"github.com/gellel/gostructures/abstracts/heaps/heap/core"
)

func Test(t *testing.T) {

	a := core.New()

	a.Push(uint8(1))

	fmt.Println(a.ToSlice())
}

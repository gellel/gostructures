package min_test

import (
	"fmt"
	"testing"

	"github.com/gellel/gostructures/abstracts/heaps/heap/min"
)

func Test(t *testing.T) {

	a := min.New()

	fmt.Println(a.Push(1))

	fmt.Println(a.ToSlice())
}

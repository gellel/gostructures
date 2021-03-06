package postorder_test

import (
	"math/rand"
	"testing"

	binary "github.com/gellel/gostructures/abstracts/trees/binary"
	postorder "github.com/gellel/gostructures/algorithms/tree-traversal/post-order"
)

func Test(t *testing.T) {

	b := binary.New(rand.Float64())

	for i := 0; i < 10; i++ {
		b.Insert(rand.Float64())
	}
	postorder.Traverse(b.Root)
}

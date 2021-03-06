package preorder_test

import (
	"math/rand"
	"testing"

	binary "github.com/gellel/gostructures/abstracts/trees/binary"
	preorder "github.com/gellel/gostructures/algorithms/tree-traversal/pre-order"
)

func Test(t *testing.T) {

	b := binary.New(rand.Float64())

	for i := 0; i < 10; i++ {
		b.Insert(rand.Float64())
	}
	preorder.Traverse(b.Root)
}

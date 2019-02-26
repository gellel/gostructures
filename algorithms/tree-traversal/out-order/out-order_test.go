package outorder_test

import (
	"math/rand"
	"testing"

	binary "github.com/gellel/gostructures/abstracts/trees/binary"
	inorder "github.com/gellel/gostructures/algorithms/tree-traversal/in-order"
)

func Test(t *testing.T) {

	b := binary.New(rand.Float64())

	for i := 0; i < 10; i++ {
		b.Insert(rand.Float64())
	}
	inorder.Traverse(b.Root)
}

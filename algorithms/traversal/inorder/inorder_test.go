package inorder_test

import (
	"math/rand"
	"testing"

	binary "github.com/gellel/gostructures/abstracts/trees/binary"
	"github.com/gellel/gostructures/algorithms/traversal/inorder"
)

func Test(t *testing.T) {

	b := binary.New(rand.Float64())

	max := 10

	for i := 0; i < max; i++ {
		b.Insert(rand.Float64())
	}

	b.Root.Walk()

	inorder.Traverse(b.Root)
}

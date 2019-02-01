package node_test

import (
	"math/rand"
	"testing"

	node "github.com/gellel/gostructures/abstracts/trees/red-black/red-black-node"
)

func Test(t *testing.T) {

	rand.Float64()

	a := node.New(5.0)

	a.InsertBinary(node.New(2.0))
	a.InsertBinary(node.New(10.0))
	a.InsertBinary(node.New(12.0))
	a.InsertBinary(node.New(8.0))
	a.InsertBinary(node.New(6.0))
	a.InsertBinary(node.New(9.0))

	a.Walk().RotateLeft()

	a.Find(2.0).About()
}

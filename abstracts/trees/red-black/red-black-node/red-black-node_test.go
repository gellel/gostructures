package node_test

import (
	"fmt"
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

	a.RotateLeft()

	a.Walk()

	fmt.Println("-----")

	b := node.New(10.0)

	b.InsertBinary(node.New(5.0))
	b.InsertBinary(node.New(2.0))
	b.InsertBinary(node.New(8.0))
	b.InsertBinary(node.New(6.0))
	b.InsertBinary(node.New(9.0))
	b.InsertBinary(node.New(12.0))

	b.RotateRight()

	b.Walk()

	fmt.Println("-----")

	c := node.New(15.0)

	c.InsertBinary(node.New(5.0))
	c.InsertBinary(node.New(1.0))

	c.RotateRight()

	c.About()

	fmt.Println("-----")

	d := node.New(1.0)

	d.InsertBinary(node.New(2.0))
	d.InsertBinary(node.New(3.0))

	d.RotateLeft()

	d.About()

	fmt.Println("-----")

	e := node.New(15.0)

	e.Insert(5.0)
	e.Insert(1.0)

	e.About()
}

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

	a.Insert(6.0)

	a.Insert(4.0)

	a.Insert(3.0)

	a.Walk()

	a.Insert(2.0)

	a.Insert(1.0)

	fmt.Println("-")

	a.Walk()

}

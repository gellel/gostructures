package node_test

import (
	"fmt"
	"math/rand"
	"testing"

	node "github.com/gellel/gostructures/abstracts/trees/red-black/red-black-node"
)

func Test(t *testing.T) {

	rand.Float64()

	a := node.New(1.0)

	a.Insert(2.0)

	a.Insert(0.05)

	a.Insert(3.0)

	fmt.Println(a.Find(3.0).Uncle)

}

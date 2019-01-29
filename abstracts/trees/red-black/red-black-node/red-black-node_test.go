package node_test

import (
	"fmt"
	"math/rand"
	"testing"

	node "github.com/gellel/gostructures/abstracts/trees/red-black/red-black-node"
)

func Test(t *testing.T) {

	a := node.New(rand.Float64())

	fmt.Println(a)

}

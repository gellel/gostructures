package node_test

import (
	"fmt"
	"testing"

	node "github.com/gellel/gostructures/abstracts/trees/binary/binary-node"
)

func Test(t *testing.T) {
	fmt.Println(node.New(2.0).AssignRight(node.New(1.0)))
}

package node_test

import (
	"fmt"
	"testing"

	node "github.com/gellel/gostructures/abstracts/lists/linked-list/list-node"
)

func Test(t *testing.T) {

	n := node.New("a")

	fmt.Println(n)

	n.Next = node.New("b")

	fmt.Println(n, n.Next)
}

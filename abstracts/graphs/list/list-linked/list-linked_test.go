package list_test

import (
	"fmt"
	"testing"

	list "github.com/gellel/gostructures/abstracts/graphs/list/list-linked"
	node "github.com/gellel/gostructures/abstracts/graphs/list/list-node"
)

func Test(t *testing.T) {
	l := list.Linked{}

	l.Head = node.New()

	fmt.Println(l)
}

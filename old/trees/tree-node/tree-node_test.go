package treenode_test

import (
	"fmt"
	"testing"

	treenode "github.com/gellel/gostructures/trees/tree-node"
)

func Test(t *testing.T) {

	b := treenode.New(7.0)

	fmt.Println(b)
}

package treenode_test

import (
	"log"
	"testing"

	treenode "github.com/gellel/gostructures/trees/tree-node"
)

func Test(t *testing.T) {

	n := treenode.New(1)

	log.Println(n.SetLeft(treenode.New(2)))

	log.Println(n.ToSlice())

	log.Println(n.ToSliceOfValues())

}

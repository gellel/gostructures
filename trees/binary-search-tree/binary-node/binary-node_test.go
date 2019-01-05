package binarynode_test

import (
	"fmt"
	"log"
	"testing"

	binarynode "github.com/gellel/gostructures/trees/binary-search-tree/binary-node"
)

func Test(t *testing.T) {

	b := binarynode.New(7)

	fmt.Println(b.Insert(2).Insert(10).Insert(3))

	fmt.Println(b.Left, b.Right)

	log.Println(b)
}

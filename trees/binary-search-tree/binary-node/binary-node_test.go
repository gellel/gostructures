package binarynode_test

import (
	"log"
	"testing"

	binarynode "github.com/gellel/gostructures/trees/binary-search-tree/binary-node"
)

func Test(t *testing.T) {

	b := binarynode.New(7)

	log.Println(b)

	log.Println(b.Insert(2).Insert(10).Insert(4).Insert(3).Insert(8).Insert(5))

}

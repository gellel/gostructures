package binarysearchtree_test

import (
	"log"
	"testing"

	binarysearchtree "github.com/gellel/gostructures/trees/binary-search-tree"
)

func Test(t *testing.T) {

	bst := binarysearchtree.New(10.0)

	log.Println(bst)

	bst.Add(8.0).Add(10.0).Add(5.0).Add(12.0).Add(-1.0)

	log.Println(bst.ToSliceFloat64())
}

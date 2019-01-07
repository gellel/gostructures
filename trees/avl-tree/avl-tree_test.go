package avltree_test

import (
	"log"
	"testing"

	avltree "github.com/gellel/gostructures/trees/avl-tree"
)

func Test(t *testing.T) {

	a := avltree.New(1.0).Add(2.0).Add(3.0)

	log.Println(a.Distribution())

	log.Println(a.ToSliceFloat64())
}

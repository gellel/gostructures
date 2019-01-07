package avltree_test

import (
	"log"
	"testing"

	avltree "github.com/gellel/gostructures/trees/avl-tree"
)

func Test(t *testing.T) {

	a := avltree.New(1.0)

	b := avltree.New(3.0)

	a.Add(2.0).Add(3.0)

	b.Add(2.0).Add(1.0)

	log.Println("a:", a.Root.Distribution())

	log.Println("b:", b.Root.Distribution())

	log.Println("a:", a.ToSliceFloat64())

	a.RotateLeft(a.Root)

	//log.Println(a.Distribution())

	//log.Println(a.ToSliceFloat64())
}

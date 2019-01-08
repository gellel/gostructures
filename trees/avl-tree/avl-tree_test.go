package avltree_test

import (
	"log"
	"testing"

	avltree "github.com/gellel/gostructures/trees/avl-tree"
)

func Test(t *testing.T) {

	a := avltree.New(1.0)

	b := avltree.New(3.0)

	c := avltree.New(3.0)

	a.Add(2.0).Add(3.0) //  left-rotate

	b.Add(2.0).Add(1.0) // right-rotate

	c.Add(1.0).Add(2.0)

	log.Println("ll:", a.Root.Distribution())

	log.Println("rr:", b.Root.Distribution())

	log.Println("lr:", c.Root.Distribution())

	log.Println("ll:", a.ToSliceFloat64())

	log.Println("rr:", b.ToSliceFloat64())

	log.Println("lr:", c.ToSliceFloat64())

	a.RotateLeft(a.Root)

	b.RotateRight(b.Root)

	c.RotateLeftRight(c.Root)

	log.Println("ll:", a.Distribution())

	log.Println("ll:", a.ToSliceFloat64())

	log.Println("rr:", b.Distribution())

	log.Println("rr:", b.ToSliceFloat64())

	log.Println("lr:", c.Distribution())

	log.Println("lr:", c.ToSliceFloat64())

}

package binarynode_test

import (
	"fmt"
	"testing"

	binarynode "github.com/gellel/gostructures/trees/binary-search-tree/binary-node"
)

func Test(t *testing.T) {

	b := binarynode.New(7)

	fmt.Println(b)

	/*b.Insert(6)

	b.Insert(5)

	fmt.Println(b.Left, b.Left.Left)

	fmt.Println(b.Minimum())*/
}

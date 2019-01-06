package binarynode_test

import (
	"fmt"
	"testing"

	binarynode "github.com/gellel/gostructures/trees/binary-search-tree/binary-node"
)

func Test(t *testing.T) {

	b := binarynode.New(7.0)

	fmt.Println(b)

}

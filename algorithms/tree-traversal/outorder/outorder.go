package outorder

import (
	"fmt"

	node "github.com/gellel/gostructures/abstracts/trees/binary/binary-node"
)

// Traverse access and prints the Binary-Tree-Node's connected children in reverse-sorted order (RIGHT, ROOT, LEFT).
func Traverse(n *node.Binary) *node.Binary {
	if n.Right != nil {
		Traverse(n.Right)
	}
	fmt.Println(fmt.Sprintf("tree.Node (%s).(%f)", n.Side, n.Value))
	if n.Left != nil {
		Traverse(n.Left)
	}
	return n
}

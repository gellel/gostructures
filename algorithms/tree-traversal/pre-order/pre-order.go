package preorder

import (
	"fmt"

	node "github.com/gellel/gostructures/abstracts/trees/binary/binary-node"
)

// Traverse navigates a binary tree (specifically from the tree's its root node), and print its children according to the "top-down" preorder traversal (ROOT, LEFT, RIGHT).
func Traverse(n *node.Binary) *node.Binary {
	fmt.Println(fmt.Sprintf("tree.Node (%s).(%f)", n.Side, n.Value))
	if n.Left != nil {
		Traverse(n.Left)
	}
	if n.Right != nil {
		Traverse(n.Right)
	}
	return n
}

package postorder

import (
	"fmt"

	node "github.com/gellel/gostructures/abstracts/trees/binary/binary-node"
)

// Traverse navigates a binary tree (specifically from the tree's its root node), and print its children according to the "bottom-up" postorder traversal (LEFT, ROOT, RIGHT).
func Traverse(n *node.Binary) *node.Binary {
	if n.Left != nil {
		Traverse(n.Left)
	}
	if n.Right != nil {
		Traverse(n.Right)
	}
	fmt.Println(fmt.Sprintf("tree.Node (%s).(%f)", n.Side, n.Value))
	return n
}

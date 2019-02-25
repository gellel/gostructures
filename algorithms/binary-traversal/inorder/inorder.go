package inorder

import (
	"fmt"

	node "github.com/gellel/gostructures/abstracts/trees/binary/binary-node"
)

// Traverse navigates a binary tree in its defined order (LEFT, ROOT, RIGHT).
func Traverse(n *node.Binary) *node.Binary {
	if n.Left != nil {
		Traverse(n.Left)
	}
	fmt.Println(fmt.Sprintf("tree.Node (%s).(%f)", n.Side, n.Value))
	if n.Right != nil {
		Traverse(n.Right)
	}
	return n
}

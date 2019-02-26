package level

import (
	"fmt"

	"github.com/gellel/gostructures/abstracts/queues/queue"

	node "github.com/gellel/gostructures/abstracts/trees/binary/binary-node"
)

func Traverse(n *node.Binary) *node.Binary {
	q := queue.New()
	q.Enqueue(n)
	for q.IsPopulated() {
		x := q.Poll().(*node.Binary)
		fmt.Println(fmt.Sprintf("tree.Node (%s).(%f)", x.Side, x.Value))
		if x.Left != nil {
			q.Enqueue(x.Left)
		}
		if x.Right != nil {
			q.Enqueue(x.Right)
		}
	}
	return n
}

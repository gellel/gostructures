package node_test

import (
	"fmt"
	"math/rand"
	"testing"

	node "github.com/gellel/gostructures/abstracts/priority-queues/max/max-node"
)

func Test(t *testing.T) {
	fmt.Println(node.New(int(rand.Intn(10-1)+1), "Go"))
}

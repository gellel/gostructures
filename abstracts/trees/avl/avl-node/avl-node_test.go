package node_test

import (
	"fmt"
	"testing"

	node "github.com/gellel/gostructures/abstracts/trees/avl/avl-node"
)

func Test(t *testing.T) {

	a := node.New(3.0)

	a.Insert(2.0).Insert(1.0)

	fmt.Println(a)

	fmt.Println(a.Find(3.0) != nil)

	fmt.Println(a.Left.Value, a.Value, a.Right.Value)

	b := node.New(1.0)

	/*
		n := node.New(1.0)

		sequence := make([]float64, 100)

		for i := 0; i < len(sequence); i++ {
			sequence[i] = rand.Float64()
			n.Insert(sequence[i])
		}

		for i := 0; i < len(sequence); i++ {
			if n.Find(sequence[i]) == nil {
				t.Fatalf("expected %f. got nil", sequence[i])
			}
		}
	*/
}

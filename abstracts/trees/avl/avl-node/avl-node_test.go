package node_test

import (
	"fmt"
	"math/rand"
	"testing"

	node "github.com/gellel/gostructures/abstracts/trees/avl/avl-node"
)

func Test(t *testing.T) {

	a := node.New(3.0)

	a.Side = "ROOT"

	a.Insert(2.0).Insert(1.0)

	fmt.Println(a)

	fmt.Println(a.Find(3.0) != nil)

	fmt.Println(a.Left.Value, a.Value, a.Right.Value)

	a.Walk()

	fmt.Println("-")

	b := node.New(1.0)

	b.Side = "ROOT"

	b.Insert(2.0).Insert(3.0)

	fmt.Println(b)

	fmt.Println(b.Find(1.0) != nil)

	fmt.Println(b.Left.Value, b.Value, b.Right.Value)

	b.Walk()

	fmt.Println("-")

	c := node.New(3.0)

	c.Side = "ROOT"

	c.Insert(1.0).Insert(2.0)

	fmt.Println(c)

	fmt.Println(c.Find(3.0) != nil)

	fmt.Println(c.Left.Value, c.Value, c.Right.Value)

	c.Walk()

	fmt.Println("-")

	d := node.New(1.0)

	d.Side = "ROOT"

	d.Insert(3.0).Insert(2.0)

	fmt.Println(d)

	fmt.Println(d.Find(3.0) != nil)

	fmt.Println(d.Left.Value, d.Value, d.Right.Value)

	d.Walk()

	fmt.Println("-")

	n := node.New(1.0)

	sequence := make([]float64, 10)

	for i := 0; i < len(sequence); i++ {
		sequence[i] = rand.Float64()
		n.Insert(sequence[i])
	}

	fmt.Println(len(sequence), len(n.ToFloatSlice()))

	fmt.Println(sequence)
	fmt.Println(n.ToFloatSlice())

	for i := 0; i < len(sequence); i++ {
		if n.Find(sequence[i]) == nil {
			t.Fatalf("expected %f. got nil", sequence[i])
		}
	}
	n.Walk()
}

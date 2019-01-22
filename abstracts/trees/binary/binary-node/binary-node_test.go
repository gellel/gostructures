package node_test

import (
	"math/rand"
	"testing"

	node "github.com/gellel/gostructures/abstracts/trees/binary/binary-node"
)

func Test(t *testing.T) {

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
}

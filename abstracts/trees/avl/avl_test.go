package avl_test

import (
	"fmt"
	"math/rand"
	"testing"

	avl "github.com/gellel/gostructures/abstracts/trees/avl"
)

func Test(t *testing.T) {

	a := avl.New(rand.Float64())

	sequence := make([]float64, 100)

	for i := 0; i < len(sequence); i++ {
		sequence[i] = rand.Float64()
		a.Insert(sequence[i])
	}

	for i := 0; i < len(sequence); i++ {
		if a.Find(sequence[i]) == nil {
			t.Fatalf("expected %f. got nil", sequence[i])
		}
	}

	fmt.Println(a.ToFloatSlice())
}

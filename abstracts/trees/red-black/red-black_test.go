package redblack_test

import (
	"fmt"
	"math/rand"
	"testing"

	redblack "github.com/gellel/gostructures/abstracts/trees/red-black"
)

func Test(t *testing.T) {

	a := redblack.New(rand.Float64())

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

	if a.Root.Side != "ROOT" {
		t.Fatalf("expected %s at address (*%p). got %s", "ROOT", &a, a.Root.Side)
	}

	fmt.Println(a.ToFloatSlice())
}

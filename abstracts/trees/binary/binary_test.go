package binary_test

import (
	"fmt"
	"math/rand"
	"testing"

	binary "github.com/gellel/gostructures/abstracts/trees/binary"
)

func Test(t *testing.T) {

	b := binary.New(rand.Float64())

	sequence := make([]float64, 100)

	for i := 0; i < len(sequence); i++ {
		sequence[i] = rand.Float64()
		b.Insert(sequence[i])
	}

	for i := 0; i < len(sequence); i++ {
		if b.Find(sequence[i]) == nil {
			t.Fatalf("expected %f. got nil", sequence[i])
		}
	}

	fmt.Println(b.ToFloatSlice())
}

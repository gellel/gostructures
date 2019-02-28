package continious_test

import (
	"fmt"
	"math/rand"
	"testing"

	continious "github.com/gellel/gostructures/algorithms/other/continious-number"
)

const (
	CEILING int = 3
	FLOOR   int = 1
	N       int = 100
)

func Test(t *testing.T) {

	a := [N]int{}

	for i := 0; i != N; i++ {
		a[i] = (rand.Intn(CEILING-FLOOR) + 1)
	}

	k, c := continious.Number(a[:])

	fmt.Println("array:", a)
	fmt.Println("length:", len(a))

	fmt.Println("longest sequential number:", k)
	fmt.Println("continious for:", c)
}

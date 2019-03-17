package matrix_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/gellel/gostructures/abstracts/graphs/matrix"
)

func Test(t *testing.T) {

	rand.Seed(time.Now().UTC().UnixNano())

	max := rand.Intn(20)

	graph := matrix.New(max)

	fmt.Println(graph)

	fmt.Println(graph.Length())

	for i := 0; i < max; i = i + 1 {
		key := rand.Intn(graph.LengthOf(i))

		success := graph.AddEdge(i, key)

		fmt.Println("key", key)

		fmt.Println("success", success)
	}

	fmt.Println(graph)
}

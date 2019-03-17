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

}

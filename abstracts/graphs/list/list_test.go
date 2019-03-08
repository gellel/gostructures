package list_test

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/gellel/gostructures/abstracts/graphs/list"
)

var (
	ALPHABET []string = []string{
		"A",
		"B",
		"C",
		"D",
		"E",
		"F",
		"G",
		"H",
		"I",
		"J",
		"K",
		"L",
		"M",
		"N",
		"O",
		"P",
		"Q",
		"R",
		"S",
		"T",
		"U",
		"V",
		"W",
		"X",
		"Y",
		"Z"}
	RANGE int = len(ALPHABET)
)

func Test(t *testing.T) {

	rand.Seed(time.Now().UTC().UnixNano())

	max := rand.Intn(RANGE)

	min := rand.Intn(max)

	verticies := ALPHABET[min:max]

	graph := list.Graph{}

	for _, vertex := range verticies {
		if graph.AddVertex(vertex) == false {
			s := fmt.Sprintf("(*%p) unable to receive vertex %s", &graph, vertex)
			log.Panicln(s)
		}
	}

	fmt.Println(graph)
	fmt.Println(verticies)
}

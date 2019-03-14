package matrix

type Graph [][]int

func (graph *Graph) AddEdge(i int, j int) bool {

}

func (graph *Graph) HasCoordinate(i int, j int) bool {
	return graph.HasRowAt(i) && graph.HasColumnAt(j)
}

func (graph *Graph) HasColumnAt(j int) bool {
	return (j > -1) && (len(*graph) > 0) && (j < len((*graph)[0]))
}

func (graph *Graph) HasRowAt(i int) bool {
	return (i > -1) && (i < len(*graph))
}

func (graph *Graph) Length() int {
	return len(*graph)
}

func (graph *Graph) LengthOf(j int) int {
	if graph.HasColumnAt(j) {
		return len((*graph)[j])
	}
	return -1
}

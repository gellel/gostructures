// Package matrix exports an adjacency matrix Graph abstract data structure.
// Adjacency matrix Graphs are a slice of slices, with each address holding
// an integer (either one or zero). If an address holds a one, this means that there
// is a relationship between the accessed row, and a sibling column. The same
// should be true should the accessed addresses be reversed; if
// the accessed address holds a zero, no relationship exists.
// Unlike an adjacency list, the matrix must scale linearly and cannot hold an
// unbalanced length of rows or columns.
package matrix

// Graph declares interfaces for Graph data structure.
type graph interface {
	AddEdge(i int, j int) bool
	Connected(i int, j int) bool
	HasCoordinates(i int, j int) bool
	Length() int
	LengthOf(i int) int
	RemoveEdge(i int, j int) bool
}

// Graph declares the pointer for Graph abstract data structure.
type Graph [][]int

// New instantiates a new adjacency list Graph.
func New(size int) *Graph {
	g := make(Graph, size)
	for i := 0; i < size; i = i + 1 {
		g[i] = make([]int, size)
	}
	return &g
}

// AddEdge assigns the Graph a new connection between one vertex and another.
func (graph *Graph) AddEdge(i int, j int) bool {
	if graph.HasCoordinates(i, j) {
		(*graph)[i][j], (*graph)[j][i] = 1, 1
		return true
	}
	return false
}

// HasCoordinates checks that matrix has both coordinates.
func (graph *Graph) HasCoordinates(i int, j int) bool {
	return i > -1 && i < len((*graph)) && j > -1 && j < len((*graph)[i])
}

// Length returns the number of vertices held by the Graph.
func (graph *Graph) Length() int {
	return len(*graph)
}

// LengthOf returns the number of edges held by the accessed vertex.
func (graph *Graph) LengthOf(i int) int {
	if i > -1 && i < len((*graph)) {
		return len((*graph)[i])
	}
	return -1
}

// RemoveEdge deletes a connection from one vertex to another.
func (graph *Graph) RemoveEdge(i int, j int) bool {
	if graph.HasCoordinates(i, j) {
		(*graph)[i][j], (*graph)[j][i] = 0, 0
		return true
	}
	return false
}

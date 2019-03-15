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
	RemoveEdge(i int, j int) bool
}

// Graph declares the pointer for Graph abstract data structure.
type Graph [][]int

// AddEdge assigns the Graph a new connection between one vertex and another.
func (graph *Graph) AddEdge(i int, j int) bool {
	if i > -1 && i < len((*graph)) && j > -1 && j < len((*graph)[i]) {
		(*graph)[i][j], (*graph)[j][i] = 1, 1
		return true
	}
	return false
}

// RemoveEdge deletes a connection from one vertex to another.
func (graph *Graph) RemoveEdge(i int, j int) bool {
	if i > -1 && i < len((*graph)) && j > -1 && j < len((*graph)[i]) {
		(*graph)[i][j], (*graph)[j][i] = 0, 0
		return true
	}
	return false
}

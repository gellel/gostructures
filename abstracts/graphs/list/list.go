package list

type graph interface {
	AddEdge(source string, destination string) *Graph
	AddVertex(source string) *Graph
	EdgeLength(source string) int
	HasEdge(source string)
	IsConnected(source string, destination string) bool
	IsEdgeEmpty(source string) bool
	IsEdgeNotEmpty(source string) bool
	IsVertexEmpty(source string) bool
	IsVertexNotEmpty(source string) bool
	RemoveEdge(source string, destination string) *Graph
	RemoveVertex(source string) *Graph
	VertexLength(source string) int
}

type Graph map[string][]string

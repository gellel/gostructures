package list

type graph interface {
	AddEdge(source string, destination string) *Graph
	AddVertex(source string) *Graph
	HasVertex(source string)
	IsConnected(source string, destination string) bool
	IsEdgeEmpty(source string) bool
	IsEdgeNotEmpty(source string) bool
	IsVertexEmpty(source string) bool
	IsVertexNotEmpty(source string) bool
	RemoveEdge(source string, destination string) *Graph
	RemoveVertex(source string) *Graph
}

type Graph map[string][]string

func (graph *Graph) AddEdge(source string, destination string) *Graph {}

func (graph *Graph) AddVertex(source string) *Graph {}

func (graph *Graph) HasVertex(source string) bool {
	_, ok := (*graph)[source]
	return ok
}

func (graph *Graph) IsConnected(source string, destination string) bool {
	if graph.HasVertex(source) {
		for _, vertex := range (*graph)[source] {
			if vertex == destination {
				return true
			}
		}
	}
	return false
}

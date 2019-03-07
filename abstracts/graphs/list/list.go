package list

type graph interface {
	AddEdge(source string, destination string) bool
	AddVertex(vertex string) bool
	Connected(source string, destination string) bool
	GetEdges(vertex string) []string
	GetVertices() []string
	Length() int
	LengthOf(vertex string) int
	RemoveEdge(source string, destination string) bool
	RemoveVertex(vertex string) bool
}

type Graph map[string][]string

func (graph *Graph) AddEdge(source string, destination string) bool {
	g := *graph
	if _, ok := g[source]; ok {
		if _, ok := g[destination]; ok {
			g[source] = append(g[source], destination)
			return true
		}
	}
	return false
}

func (graph *Graph) AddVertex(vertex string) bool {
	if _, ok := (*graph)[vertex]; ok == false {
		(*graph)[vertex] = []string{}
		return true
	}
	return false
}

func (graph *Graph) Connected(source string, destination string) bool {
	g := *graph
	if _, ok := g[source]; ok {
		if _, ok := g[destination]; ok {
			for _, vertex := range g[source] {
				if vertex == destination {
					return true
				}
			}
		}
	}
	return false
}

func (graph *Graph) GetEdges(vertex string) []string {
	if _, ok := (*graph)[vertex]; ok {
		return (*graph)[vertex]
	}
	return nil
}

func (graph *Graph) GetVertices() []string {
	vertices := []string{}
	for key := range *graph {
		vertices = append(vertices, key)
	}
	return vertices
}

func (graph *Graph) Length() int {
	return len(*graph)
}

func (graph *Graph) LengthOf(vertex string) int {
	if _, ok := (*graph)[vertex]; ok {
		return len((*graph)[vertex])
	}
	return -1
}

func (graph *Graph) RemoveEdge(source string, destination string) bool {
	g := *graph
	if _, ok := g[source]; ok {
		if _, ok := g[destination]; ok {
			for i, vertex := range g[source] {
				if vertex == destination {
					g[source] = append(g[source][:i], g[source][i+1:]...)
					return true
				}
			}
		}
	}
	return false
}

func (graph *Graph) RemoveVertex(source string) bool {
	if _, ok := (*graph)[source]; ok {
		delete((*graph), source)
		return true
	}
	return false
}

var _ graph = (*Graph)(nil)

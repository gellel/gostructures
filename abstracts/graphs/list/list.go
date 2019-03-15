// Package list exports an adjacency list Graph abstract data structure.
// Adjacency list Graphs are a map based collection of string based keys
// and values of string slices. Each key within this struct represents a
// vertex of the Graph. For each vertex, the list value held by this key
// is this vertex's connections to other vertex's of the Graph.
package list

// Graph declares interfaces for Graph data structure.
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

// Graph declares the pointer for Graph abstract data structure.
type Graph map[string][]string

// AddEdge assigns the Graph a new connection between one vertex and another.
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

// AddVertex adds a new key to the Graph that holds its connections.
func (graph *Graph) AddVertex(vertex string) bool {
	if _, ok := (*graph)[vertex]; ok == false {
		(*graph)[vertex] = []string{}
		return true
	}
	return false
}

// Connected checks whether two vertices are connected.
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

// GetEdges collects all the edges for the accessed vertex.
func (graph *Graph) GetEdges(vertex string) []string {
	if _, ok := (*graph)[vertex]; ok {
		return (*graph)[vertex]
	}
	return nil
}

// GetVertices collects all verticies from the Graph.
func (graph *Graph) GetVertices() []string {
	vertices := []string{}
	for key := range *graph {
		vertices = append(vertices, key)
	}
	return vertices
}

// Length returns the number of vertices held by the Graph.
func (graph *Graph) Length() int {
	return len(*graph)
}

// LengthOf returns the number of edges held by the accessed vertex.
func (graph *Graph) LengthOf(vertex string) int {
	if _, ok := (*graph)[vertex]; ok {
		return len((*graph)[vertex])
	}
	return -1
}

// RemoveEdge deletes a connection from one vertex to another. If there are two connections, the destination relationship is not modified.
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

// RemoveVertex deletes a vertex from the Graph.
func (graph *Graph) RemoveVertex(source string) bool {
	if _, ok := (*graph)[source]; ok {
		delete((*graph), source)
		for k := range *graph {
			for i := range (*graph)[k] {
				a := []string{}
				if (*graph)[k][i] != source {
					a = append(a, (*graph)[k][i])
				}
				(*graph)[k] = a
			}
		}
		return true
	}
	return false
}

var _ graph = (*Graph)(nil)

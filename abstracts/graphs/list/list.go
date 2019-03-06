// Package list exports an Adjacency-List Graph. Adjacency-List Graph manages
// a collection of nodes and edges representing a connected series of data.
// Graph vertices are stored structs of Linked-Lists, and every vertex stores a collection of
// its of adjacent vertices.
package list

import list "github.com/gellel/gostructures/abstracts/graphs/list/list-linked"

type Graph []*list.Linked

func (graph *Graph) AddEdge(source int, destination int) {}

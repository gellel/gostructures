package list

import node "github.com/gellel/gostructures/abstracts/graphs/list/list-node"

type linked interface{}

type Linked struct {
	Head *node.List
}

func New() *Linked {
	return &Linked{}
}

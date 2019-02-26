package closed

import (
	node "github.com/gellel/gostructures/abstracts/hash-tables/closed-table/closed-node"
	"github.com/gellel/gostructures/abstracts/linked-lists/single"
)

const (
	SIZE int = 8
)

type table interface {
	Add(key string, value interface{}) int
	Append(n *node.Closed) int
	Contains(key string) bool
	Create() *single.LinkedList
	Delete(key string) bool
	Hash(key string) int
	IsEmpty(p int) bool
	IsNotEmpty(p int) bool
	Length() int
	LengthAt(p int) int
	LengthBy(key string) int
	Search(key string) interface{}
}

type Table [SIZE]*single.LinkedList

func (table *Table) Add(key string, value interface{}) int {
	i := table.Hash(key)
	n := node.Closed{i, value}
}

var _ table = (*Table)(nil)

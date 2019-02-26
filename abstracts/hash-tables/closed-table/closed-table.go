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
	Contains(key string) bool
	Delete(key string) bool
	Hash(key string) int
	Length() int
	LengthAt(p int) int
	LengthBy(key string) int
	Search(key string) interface{}
}

type Table [SIZE]*single.LinkedList

var _ table = (*Table)(nil)
// Package closed exports a closed-address hash table. 
package closed

import (
	"github.com/gellel/gostructures/abstracts/linked-lists/single"
)

const (
	// SIZE declares the maximum address space for the Closed-Addressed Hash-Table.
	SIZE int = 8
)

type Hash [SIZE]*single.LinkedList

// Access retrieves the pointer stored at the argument address. Does not check that the position is in the table range.
func (hash *Hash) Access(p int) *single.LinkedList {
	return (*hash)[p]
}

// Add pushes a new hashed entry into the table. Value may be stored at the root or nth position with the address linked-list.
func (hash *Hash) Add(key string, value interface{}) int {
	p := hash.Hash(key)
	if hash.Empty(p) {
		(*hash)[p] = single.New()
	}
	hash.Access(p).Append(key)
	return p
}

// Bounds checks that the argument integer falls within the supported range of addresses in the table.
func (hash *Hash) Bounds(p int) bool {
	return ((p > -1) && (p < len((*hash))))
}

// Delete removes a value from the table.
func (hash *Hash) Delete(key string) bool {
	p := hash.Hash(key)
	if hash.Empty(p) {
		return false
	}
	return hash.Access(p).Remove(key)
}

// Empty checks the ponter stored at the argument address and evaluates whether the stored item there is currently nil.
func (hash *Hash) Empty(p int) bool {
	return (hash.Access(p) == nil)
}

// Get accesses the corresponding address where the value is held by the table.
func (hash *Hash) Get(key string) interface{} {
	p := hash.Hash(key)
	if hash.Empty(p) {
		return nil
	}
	return hash.Access(p).Find(key)
}

// Hash computes and converts key to a hash value so that it may be assigned position in the linked-list. Hashing method is quite primative and address size is limited.
func (hash *Hash) Hash(key string) int {
	s := 0
	for _, k := range []byte(key) {
		s = (s + int(k))
	}
	return (s % SIZE)
}

// Package definition contains interface
// required for implementing functional
// Single Linked-List.
package definition

import node "github.com/gellel/gostructures/abstracts/lists/linked-list/list-node"

// LinkedList is a single Linked-List interface decleration.
type LinkedList interface {
	Append(property interface{}) *LinkedList
	AppendTail(node *node.Node) *LinkedList
	Contains(property interface{}) bool
	Delete(property interface{}) *LinkedList
	Find(property interface{}) *node.Node
	HasEmptyHead() bool
	HasHead() bool
	HasEmptyTail() bool
	HasTail() bool
	IsEmpty() bool
	IsPopulated() bool
	Prepend(property interface{}) *LinkedList
	PrependHead(node *node.Node) *LinkedList
	SizeOf() int
}

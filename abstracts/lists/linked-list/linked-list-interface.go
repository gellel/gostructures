package definition

type LinkedList interface {
    Append(property interface{}) *LinkedList
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
}
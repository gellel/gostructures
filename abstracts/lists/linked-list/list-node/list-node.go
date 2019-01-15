package node

type Node struct {
    Next *Node
	Value interface{}
}

func New(property interface{}) *Node {
    return &Node{Value: property}
}
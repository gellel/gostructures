package doublelinkedlist

type DoubleLinkedList interface {
	Add() *DoubleLinkedList
}

type DoubleLinkedList struct{}

func (doubleLinkedList *DoubleLinkedList) Add() *DoubleLinkedList {
	return doubleLinkedList
}

var _ LinkedList = (*DoubleLinkedList)(nil)

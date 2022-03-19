package linkedlist

import (
	"fmt"

	"com.github/MadhavaAdiga/GolangDS/standard"
)

type LinkedList interface {
	GetHead() (interface{}, error)
	GetTail() (interface{}, error)
	Add(interface{}) error
	AddFirst(interface{}) error
	Remove() (interface{}, error)
	RemoveAt(int) (interface{}, error)
	RemoveFirst() (interface{}, error)
	IndexOf(interface{}) (int, error)
	Contains(interface{}) bool
	elementAt(int) (interface{}, error)
	Size() int
	IsEmpty() bool
	standard.Collection
}

type linkedListIterator struct {
	index int
	list  LinkedList
}

func newLinkedListIterator(list LinkedList) *linkedListIterator {
	return &linkedListIterator{
		index: 0,
		list:  list,
	}
}

// iterator method for looping
func (iterator *linkedListIterator) HasNext() bool {
	return iterator.index < iterator.list.Size()
}

// iterator method to get node data
func (iterator *linkedListIterator) GetNext() interface{} {
	v, error := iterator.list.elementAt(iterator.index)
	if error != nil {
		fmt.Print(error)
		return nil
	}
	iterator.index++
	return v
}

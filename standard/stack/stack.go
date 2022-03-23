package stack

import (
	"fmt"

	"com.github/MadhavaAdiga/GolangDS/standard"
)

type Stack interface {
	Push(interface{}) error
	Pop() (interface{}, error)
	Peek() (interface{}, error)
	IsEmpty() bool
	Size() int
	elementAt(int) (interface{}, error)
	standard.Collection
}

type ListStackIterator struct {
	index int
	stack Stack
}

func newStackIterator(stack Stack) *ListStackIterator {
	return &ListStackIterator{
		index: stack.Size() - 1,
		stack: stack,
	}
}

func (iterator *ListStackIterator) HasNext() bool {
	return iterator.index >= 0
}

func (iterator *ListStackIterator) GetNext() interface{} {
	v, err := iterator.stack.elementAt(iterator.index)
	if err != nil {
		fmt.Print("stack is empty")
		return nil
	}
	iterator.index--
	return v
}

package stack

import (
	"fmt"

	"com.github/MadhavaAdiga/GolangDS/standard"
	linkedlist "com.github/MadhavaAdiga/GolangDS/standard/linkedList"
)

type ListStack struct {
	list linkedlist.LinkedList
}

func NewListStack() Stack {
	return &ListStack{
		list: linkedlist.NewSLinkedList(),
	}
}

func (stack *ListStack) Push(ele interface{}) error {
	return stack.list.Add(ele)
}

func (stack *ListStack) Pop() (interface{}, error) {
	v, err := stack.list.Remove()
	if err != nil {
		return nil, fmt.Errorf("stack is empty")
	}
	return v, nil
}

func (stack *ListStack) Peek() (interface{}, error) {
	return stack.list.GetTail()
}

func (stack *ListStack) IsEmpty() bool {
	return stack.list.Size() == 0
}

func (stack *ListStack) Size() int {
	return stack.list.Size()
}

func (stack *ListStack) CreateIterator() standard.Iterator {
	return newStackIterator(stack.list)
}

type ListStackIterator struct {
	index int
	list  linkedlist.LinkedList
}

func newStackIterator(list linkedlist.LinkedList) *ListStackIterator {
	return &ListStackIterator{
		index: list.Size() - 1,
		list:  list,
	}
}

func (iterator *ListStackIterator) HasNext() bool {
	return iterator.index >= 0
}

func (iterator *ListStackIterator) GetNext() interface{} {
	v, err := iterator.list.ElementAt(iterator.index)
	if err != nil {
		fmt.Print("stack is empty")
		return nil
	}
	iterator.index--
	return v
}

var _ Stack = (*ListStack)(nil)
var _ standard.Collection = (*ListStack)(nil)

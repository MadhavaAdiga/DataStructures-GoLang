package stack

import (
	"fmt"

	"com.github/MadhavaAdiga/GolangDS/standard/internal/types"
	linkedlist "com.github/MadhavaAdiga/GolangDS/standard/linkedList"
)

type LinkedListImpl struct {
	list linkedlist.LinkedList
}

func NewListStack() Stack {
	return &LinkedListImpl{
		list: linkedlist.NewSLinkedList(),
	}
}

func (stack *LinkedListImpl) Push(ele interface{}) error {
	return stack.list.Add(ele)
}

func (stack *LinkedListImpl) Pop() (interface{}, error) {
	v, err := stack.list.Remove()
	if err != nil {
		return nil, fmt.Errorf("stack is empty")
	}
	return v, nil
}

func (stack *LinkedListImpl) Peek() (interface{}, error) {
	return stack.list.GetTail()
}

func (stack *LinkedListImpl) IsEmpty() bool {
	return stack.list.Size() == 0
}

func (stack *LinkedListImpl) Size() int {
	return stack.list.Size()
}

func (stack *LinkedListImpl) elementAt(index int) (interface{}, error) {
	return stack.list.ElementAt(index)
}

func (stack *LinkedListImpl) CreateIterator() types.Iterator {
	return newStackIterator(stack)
}

var _ Stack = (*LinkedListImpl)(nil)
var _ types.Collection = (*LinkedListImpl)(nil)

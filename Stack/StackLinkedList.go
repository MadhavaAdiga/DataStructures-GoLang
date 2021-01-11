package stack

import (
	linkedlist "github.com/MadhavaAdiga/DataStructure-GoLang/LinkedList"
)

/*
LinkedListImpl is an array implementation of Stack
*/
type LinkedListImpl struct {
	element *linkedlist.LinkedList
}

/*
Push is a method to insert an element on top of stack
*/
func (stack *LinkedListImpl) Push(element interface{}) error {
	if stack.element == nil {
		stack.element = &linkedlist.LinkedList{}
	}

	err := stack.element.Insert(element)
	return err
}

/*
Pop is a method to remove an element from top of stack
*/
func (stack *LinkedListImpl) Pop() (interface{}, error) {
	value, err := stack.element.Remove()

	return value, err
}

/*
Size is a method which returns size of the stack
*/
func (stack *LinkedListImpl) Size() int {
	return stack.element.Length()
}

/*
Peek is a method which returns top element of stack
*/
func (stack LinkedListImpl) Peek() interface{} {
	return stack.element.GetHead()
}

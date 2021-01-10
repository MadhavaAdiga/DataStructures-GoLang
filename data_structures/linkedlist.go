package linkedlist

import (
	"fmt"
	"reflect"
)

/*
LinkedList is a singly linked list
*/
type LinkedList struct {
	head   *node
	length int
}

/*
Insert is method to add value to begining of a linkedlist
*/
func (list *LinkedList) Insert(value interface{}) error {
	var newNode *node = &node{
		_node{value}, nil,
	}

	if list.head == nil {
		list.head = newNode
		list.length++
		return nil
	}

	// type check
	if reflect.TypeOf(list.head.data) != reflect.TypeOf(value) {
		return fmt.Errorf("not same type \n required type : %s", reflect.TypeOf(list.head.data))
	}

	var temp *node = list.head

	list.head = newNode
	list.head.next = temp

	list.length++

	return nil
}

/*
InsertAtEnd is method to add value to end of a linkedlist
*/
func (list *LinkedList) InsertAtEnd(value interface{}) error {
	var node *node = &node{
		_node{data: value}, nil,
	}

	if list.length == 0 {
		list.head = node
		return nil
	}
	// type check
	if reflect.TypeOf(list.head.data) != reflect.TypeOf(value) {
		return fmt.Errorf("not same type \n required type : %s", reflect.TypeOf(list.head.data))
	}

	current := list.head

	for current.next != nil {
		current = current.next
	}

	current.next = node

	return nil
}

/*
Pop is a method to delete first node in the linked list and returns a value
*/
func (list *LinkedList) Pop() (interface{}, error) {
	if list.length == 0 {
		return nil, fmt.Errorf("No value in list to pop")
	}

	value := list.head.data

	list.head = list.head.next
	list.length--
	return value, nil
}

/*
DelelteByValue is a method to delete a node in the linked list by value
*/
func (list *LinkedList) DelelteByValue(value interface{}) (bool, error) {
	if list.length == 0 {
		return false, fmt.Errorf("No value in list to delete")
	}

	if list.head.data == value {
		list.head = list.head.next
		list.length--
		return true, nil
	}

	var current *node = list.head

	for current.next.data != value {
		if current.next.next == nil {
			return false, fmt.Errorf("Value not found in list")
		}

		current = current.next
	}

	current.next = current.next.next
	list.length--

	return true, nil
}

/*
PrintList is a method that prints a linked list values
*/
func (list LinkedList) PrintList() {
	var current *node = list.head

	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Printf("\n")
}

/*
GetLength is a mthod to return length of a linked list
*/
func (list LinkedList) GetLength() int {
	return list.length
}

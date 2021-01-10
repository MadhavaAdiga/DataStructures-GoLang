package linkedlist

import (
	"fmt"
	"reflect"
)

/*
DoublyLinkedList is a doubly linked list
*/
type DoublyLinkedList struct {
	head   *doublyNode
	tail   *doublyNode
	length int
}

/*
Insert is method to add value to begining of a DoublyLinkedList
*/
func (list *DoublyLinkedList) Insert(value interface{}) error {
	var newNode *doublyNode = &doublyNode{_node{value}, nil, nil}

	if list.length == 0 {
		list.head = newNode
		list.tail = newNode
		list.length++
		return nil
	}

	// type check
	if reflect.TypeOf(list.head.data) != reflect.TypeOf(value) {
		return fmt.Errorf("not same type \n required type : %s", reflect.TypeOf(list.head.data))
	}

	temp := list.head

	list.head = newNode
	list.head.next = temp
	temp.prev = list.head
	list.length++

	return nil
}

/*
InsertAtTail is method to add value to begining of a DoublyLinkedList
*/
func (list *DoublyLinkedList) InsertAtTail(value interface{}) error {
	var newNode *doublyNode = &doublyNode{_node{value}, nil, nil}

	if list.length == 0 {
		list.head = newNode
		list.tail = newNode
		list.length++
		return nil
	}

	// type check
	if reflect.TypeOf(list.head.data) != reflect.TypeOf(value) {
		return fmt.Errorf("not same type \n required type : %s", reflect.TypeOf(list.head.data))
	}

	temp := list.tail

	list.tail = newNode
	list.tail.prev = temp
	temp.next = list.tail
	list.length++

	return nil
}

/*
Pop is a method to delete first node in the  doubly linked list and returns a value
*/
func (list *DoublyLinkedList) Pop() (interface{}, error) {
	if list.length == 0 {
		return nil, fmt.Errorf("No value in list to pop")
	}

	value := list.head.data

	list.head = list.head.next
	list.head.prev = nil
	list.length--

	return value, nil
}

/*
DeleteByValue is a method to delete a node fromdoubly linked list by value
*/
func (list *DoublyLinkedList) DeleteByValue(value interface{}) (bool, error) {
	if list.length == 0 {
		return false, fmt.Errorf("No value in list to delete")
	}

	if list.head.data == value {
		list.head = list.head.next
		list.head.prev = nil
		list.length--
		return true, nil
	}

	var current *doublyNode = list.head

	for current.data != value {
		if current.next == nil {
			return false, fmt.Errorf("Value not found in list")
		}

		current = current.next
	}

	previous := current.prev
	next := current.next

	previous.next = next

	//edge case for deleting a tail node
	if next != nil {
		next.prev = previous
	} else {
		previous.next = nil
		list.tail = previous
	}

	list.length--

	return true, nil
}

/*
Poll is a method to delete last node in the  doubly linked list and returns a value
*/
func (list *DoublyLinkedList) Poll() (interface{}, error) {
	if list.length == 0 {
		return nil, fmt.Errorf("No value in list to delete")
	}

	value := list.tail.data

	list.tail = list.tail.prev
	list.tail.next = nil
	list.length--

	return value, nil
}

/*
PrintList is a method that prints a doubly linked list values
*/
func (list DoublyLinkedList) PrintList() {
	var current *doublyNode = list.head

	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Printf("\n")
}

/*
PrintReverse is a method that prints a doubly linked list values in reverse order
*/
func (list DoublyLinkedList) PrintReverse() {
	current := list.tail

	for current != nil {
		fmt.Printf("<- %d  ", current.data)
		current = current.prev
	}
	fmt.Printf("\n")
}

/*
PeekHead is a method to return the head of a linked list
*/
func (list DoublyLinkedList) PeekHead() interface{} {
	return list.head.data
}

/*
PeekTail is a method to return the head of a linked list
*/
func (list DoublyLinkedList) PeekTail() interface{} {
	return list.tail.data
}

/*
PrintHead is a method to print the head of a linked list
*/
func (list DoublyLinkedList) PrintHead() {
	fmt.Println(list.head.data)
}

/*
PrintTail is a method to print the head of a linked list
*/
func (list DoublyLinkedList) PrintTail() {
	fmt.Println(list.tail.data)
}

/*
GetLength is a mthod to return length of a linked list
*/
func (list DoublyLinkedList) GetLength() int {
	return list.length
}

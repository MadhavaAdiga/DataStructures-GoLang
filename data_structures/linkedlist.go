package linkedlist

import "fmt"

/*
Node is a singly linked list node
*/
type Node struct {
	data int
	next *Node
}

/*
LinkedList is a singly linked list
*/
type LinkedList struct {
	head   *Node
	length int
}

/*
Insert is method to add value to begining of a linkedlist
*/
func (list *LinkedList) Insert(value int) {
	var node *Node = &Node{data: value}

	if list.head == nil {
		list.head = node
		list.length++
		return
	}

	var temp *Node = list.head

	list.head = node
	list.head.next = temp

	list.length++
}

/*
InsertAtEnd is method to add value to end of a linkedlist
*/
func (list *LinkedList) InsertAtEnd(value int) {
	var node *Node = &Node{data: value}

	if list.length == 0 {
		list.head = node
		return
	}

	current := list.head

	for current.next != nil {
		current = current.next
	}

	current.next = node
}

/*
DelelteByValue is a method to delete a node in the linked list by value
*/
func (list *LinkedList) DelelteByValue(value int) error {
	if list.length == 0 {
		return fmt.Errorf("No value in list to delete")
	}

	if list.head.data == value {
		list.head = list.head.next
		list.length--
		return nil
	}

	var current *Node = list.head

	for current.next.data != value {
		if current.next.next == nil {
			return fmt.Errorf("Value not found in list")
		}

		current = current.next
	}

	current.next = current.next.next
	list.length--

	return nil
}

/*
PrintList is a method that prints a linked list values
*/
func (list LinkedList) PrintList() {
	var current *Node = list.head

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

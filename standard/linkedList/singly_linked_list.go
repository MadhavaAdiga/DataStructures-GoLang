package linkedlist

import (
	"fmt"
	"reflect"

	"com.github/MadhavaAdiga/GolangDS/standard"
)

type SinglyLinkedList struct {
	head *node
	tail *node
	size int
}

func NewSLinkedList() *SinglyLinkedList {

	return &SinglyLinkedList{}
}
func (list *SinglyLinkedList) GetHead() (interface{}, error) {
	if list.head == nil {
		return nil, fmt.Errorf("empty list")
	}
	return list.head.data, nil
}

func (list *SinglyLinkedList) GetTail() (interface{}, error) {
	if list.head == nil {
		return nil, fmt.Errorf("empty list")
	}
	return list.tail.data, nil
}

// inserts element to the end of lists
func (list *SinglyLinkedList) Add(ele interface{}) error {
	if !list.typeCheck(ele) {
		return fmt.Errorf("type mismatch")
	}

	node := newNode(ele)
	if list.head == nil {
		list.head, list.tail = node, node
		list.size++
		return nil
	}
	list.tail.next = node
	list.tail = node
	list.size++

	return nil
}

func (list *SinglyLinkedList) AddFirst(ele interface{}) error {
	if !list.typeCheck(ele) {
		return fmt.Errorf("type mismatch")
	}

	node := newNode(ele)
	if list.head == nil {
		list.head, list.tail = node, node
		list.size++
		return nil
	}

	node.next = list.head
	list.head = node
	list.size++

	return nil
}

// Removes last node of the list
func (list *SinglyLinkedList) Remove() (interface{}, error) {
	if list.head == nil {
		return nil, fmt.Errorf("empty list")
	}

	data := list.tail.data

	pointer := list.head
	// iterate till last but one one node
	for pointer.next.next != nil {
		pointer = pointer.next
	}

	pointer.next = nil
	list.tail = pointer
	list.size--

	if list.IsEmpty() {
		list.head = nil
	}

	return data, nil
}

func (list *SinglyLinkedList) RemoveFirst() (interface{}, error) {
	if list.head == nil {
		return nil, fmt.Errorf("empty list")
	}

	data := list.head.data
	list.head = list.head.next
	list.size--

	if list.IsEmpty() {
		list.tail = nil
	}

	return data, nil
}

// position is assumend index value starting from 0
func (list *SinglyLinkedList) RemoveAt(position int) (interface{}, error) {
	if position < 0 || position >= list.size {
		return nil, fmt.Errorf("index is out of bound")
	}
	// if list.head == nil {
	// 	return nil, fmt.Errorf("empty list")
	// }

	var data interface{}
	if position == 0 {
		data = list.head.data
		list.head = list.head.next
		list.size--
		return data, nil
	}

	var pointer *node = list.head // index 0 node
	// loops till previous node of position to remove
	for i := 1; i < position; i++ {
		pointer = pointer.next
	}
	// pointer.next node is one to be removed
	data = pointer.next.data
	pointer.next = pointer.next.next
	list.size--

	return data, nil
}

func (list *SinglyLinkedList) IndexOf(ele interface{}) (int, error) {
	if !list.typeCheck(ele) {
		return -1, fmt.Errorf("type mismatch")
	}
	if list.head == nil {
		return -1, fmt.Errorf("empty list")
	}

	pointer := list.head

	for i := 0; pointer.next != nil; i++ {
		if reflect.DeepEqual(pointer.data, ele) {
			return i, nil
		}
		pointer = pointer.next
	}

	return -1, fmt.Errorf("%v doest not exist in the list", ele)
}

func (list *SinglyLinkedList) Contains(ele interface{}) bool {
	v, _ := list.IndexOf(ele)
	return v != -1
}

func (list *SinglyLinkedList) ElementAt(position int) (interface{}, error) {
	if position < 0 || position >= list.size {
		return nil, fmt.Errorf("index is out of bound")
	}
	if list.head == nil {
		return nil, fmt.Errorf("empty list")
	}

	pointer := list.head
	for i := 0; i != position; i++ {
		pointer = pointer.next
	}

	return pointer.data, nil
}

// get size of the linked list
func (list *SinglyLinkedList) Size() int {
	return list.size
}

// check weather linked list is empty
func (list *SinglyLinkedList) IsEmpty() bool {
	return list.size == 0
}

func (list *SinglyLinkedList) CreateIterator() standard.Iterator {
	return newLinkedListIterator(list)
}

func (list *SinglyLinkedList) typeCheck(ele interface{}) bool {
	if list.head == nil {
		return true
	}
	return reflect.TypeOf(ele) == reflect.TypeOf(list.head.data)
}

// checking for implementation
var _ LinkedList = (*SinglyLinkedList)(nil)
var _ standard.Collection = (*SinglyLinkedList)(nil)

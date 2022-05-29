package linkedlist

import (
	"fmt"
	"reflect"

	"com.github/MadhavaAdiga/GolangDS/standard/internal/types"
)

type DoublyLinkedList struct {
	head *node
	tail *node
	size int
}

func NewDLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (list *DoublyLinkedList) GetHead() (interface{}, error) {
	if list.head == nil {
		return nil, fmt.Errorf("empty list")
	}
	return list.head.data, nil
}

func (list *DoublyLinkedList) GetTail() (interface{}, error) {
	if list.head == nil {
		return nil, fmt.Errorf("empty list")
	}
	return list.tail.data, nil
}

// inserts element to the end of list
func (list *DoublyLinkedList) Add(ele interface{}) error {
	if !list.typeCheck(ele) {
		return fmt.Errorf("type mismatch")
	}

	var node *node = newNode(ele)
	if list.tail == nil {
		list.head, list.tail = node, node
		list.size++
		return nil
	}

	node.prev = list.tail
	list.tail.next = node
	list.tail = list.tail.next
	list.size++

	return nil
}

func (list *DoublyLinkedList) AddFirst(ele interface{}) error {
	if !list.typeCheck(ele) {
		return fmt.Errorf("type mismatch")
	}

	node := newNode(ele)
	if list.head == nil {
		list.head, list.tail = node, node
		list.size++
		return nil
	}

	list.head.prev = node
	node.next = list.head
	list.head = node
	list.size++

	return nil
}

// Removes last node of the list
func (list *DoublyLinkedList) Remove() (interface{}, error) {
	if list.head == nil {
		return nil, fmt.Errorf("empty list")
	}

	data := list.tail.data

	list.tail = list.tail.prev
	list.size--

	if list.IsEmpty() {
		list.head = nil
	} else {
		list.tail.next = nil
	}
	return data, nil
}

func (list *DoublyLinkedList) RemoveFirst() (interface{}, error) {
	if list.head == nil {
		return nil, fmt.Errorf("empty list")
	}

	data := list.head.data

	list.head = list.head.next
	// list.head.prev = nil
	list.size--

	if list.IsEmpty() {
		list.tail = nil
	} else {
		list.head.prev = nil
	}

	return data, nil
}

func (list *DoublyLinkedList) RemoveAt(position int) (interface{}, error) {
	if position < 0 || position >= list.size {
		return nil, fmt.Errorf("index is out of bound")
	}
	if list.head == nil {
		return nil, fmt.Errorf("empty list")
	}

	var pointer *node
	if position < list.size/2 {
		pointer = list.head
		// loops from head till position to remove
		for i := 0; i != position; i++ {
			pointer = pointer.next
		}
	} else {
		pointer = list.tail
		// loops from tail till position to remove
		for i := list.size - 1; i != position; i-- {
			pointer = pointer.prev
		}
	}

	if pointer.next == nil {
		return list.Remove()
	}
	if pointer.prev == nil {
		return list.RemoveFirst()
	}

	data := pointer.data
	pointer.next.prev = pointer.prev
	pointer.prev.next = pointer.next

	pointer.data = nil
	pointer.next = nil
	pointer.prev = nil

	list.size--

	return data, nil
}

func (list *DoublyLinkedList) IndexOf(ele interface{}) (int, error) {
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

func (list *DoublyLinkedList) Contains(ele interface{}) bool {
	v, _ := list.IndexOf(ele)
	return v != -1
}

func (list *DoublyLinkedList) ElementAt(position int) (interface{}, error) {
	if position < 0 || position >= list.size {
		return nil, fmt.Errorf("index is out of bound")
	}
	// if list.head == nil {
	// 	return -1, fmt.Errorf("empty list")
	// }

	var pointer *node
	if position < list.size/2 {
		pointer = list.head
		// loops from head till position to remove
		for i := 0; i != position; i++ {
			pointer = pointer.next
		}
	} else {
		pointer = list.tail
		// loops from tail till position to remove
		for i := list.size - 1; i != position; i-- {
			pointer = pointer.prev
		}
	}

	return pointer.data, nil
}

func (list *DoublyLinkedList) Size() int {
	return list.size
}

func (list *DoublyLinkedList) IsEmpty() bool {
	return list.size == 0
}

func (list *DoublyLinkedList) CreateIterator() types.Iterator {
	return newLinkedListIterator(list)
}

func (list *DoublyLinkedList) typeCheck(ele interface{}) bool {
	if list.head == nil {
		return true
	}
	return reflect.TypeOf(ele) == reflect.TypeOf(list.head.data)
}

// checking for implementation
var _ LinkedList = (*DoublyLinkedList)(nil)
var _ types.Collection = (*DoublyLinkedList)(nil)

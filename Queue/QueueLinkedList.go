package queue

import (
	"fmt"

	linkedlist "github.com/MadhavaAdiga/DataStructure-GoLang/LinkedList"
)

/*
LinkedListImpl is a doubly linked list implementation of Queue
*/
type LinkedListImpl struct {
	element *linkedlist.DoublyLinkedList
}

/*
Enqueue is method to insert element into a queue
*/
func (queue *LinkedListImpl) Enqueue(value interface{}) error {
	if queue.element == nil {
		queue.element = &linkedlist.DoublyLinkedList{}
	}

	err := queue.element.Insert(value)

	return err
}

/*
Dequeue is method to remove the first element of a queue
*/
func (queue *LinkedListImpl) Dequeue() (interface{}, error) {
	if queue.Size() == 0 {
		return nil, fmt.Errorf("No value in queue to dequeue")
	}

	return queue.element.RemoveLast()
}

/*
Poll is method to remove the first element of a queue
*/
func (queue *LinkedListImpl) Poll() (interface{}, error) {
	if queue.Size() == 0 {
		return nil, fmt.Errorf("No value in queue to poll")
	}
	return queue.Dequeue()
}

/*
Size is method to return size of queue
*/
func (queue *LinkedListImpl) Size() int {
	return queue.element.Length()
}

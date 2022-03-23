package queue

import (
	"fmt"

	"com.github/MadhavaAdiga/GolangDS/standard"
	linkedlist "com.github/MadhavaAdiga/GolangDS/standard/linkedList"
)

type LinkedListImpl struct {
	list linkedlist.LinkedList
}

func NewListQueue() Queue {
	return &LinkedListImpl{
		list: linkedlist.NewDLinkedList(),
	}
}

func (queue *LinkedListImpl) Enqueue(ele interface{}) error {
	return queue.list.Add(ele)
}

func (queue *LinkedListImpl) Dequeue() (interface{}, error) {
	if queue.IsEmpty() {
		return nil, fmt.Errorf("No value in queue to dequeue")
	}

	return queue.list.RemoveFirst()
}

func (queue *LinkedListImpl) Peek() (interface{}, error) {
	if queue.IsEmpty() {
		return nil, fmt.Errorf("No value in queue to peek")
	}
	return queue.list.GetHead()
}

func (queue *LinkedListImpl) elementAt(index int) (interface{}, error) {
	return queue.list.ElementAt(index)
}

func (queue *LinkedListImpl) Size() int {
	return queue.list.Size()
}

func (queue *LinkedListImpl) IsEmpty() bool {
	return queue.Size() == 0
}

func (queue *LinkedListImpl) CreateIterator() standard.Iterator {
	return newQueueIterator(queue)
}

var _ Queue = (*LinkedListImpl)(nil)
var _ standard.Collection = (*LinkedListImpl)(nil)

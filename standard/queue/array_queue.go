package queue

import (
	"fmt"
	"reflect"

	"com.github/MadhavaAdiga/GolangDS/standard"
)

type ArrayImpl struct {
	size int
	arr  []interface{}
}

func NewArrayQueue() Queue {
	return &ArrayImpl{
		size: 0,
	}
}

func (queue *ArrayImpl) Enqueue(ele interface{}) error {
	if !queue.checkType(ele) {
		return fmt.Errorf("type mismatch")
	}

	queue.arr = append(queue.arr, ele)
	queue.size++

	return nil
}

func (queue *ArrayImpl) Dequeue() (interface{}, error) {
	if queue.IsEmpty() {
		return nil, fmt.Errorf("queue is empty")
	}

	data := queue.arr[0]
	queue.arr = queue.arr[1:]
	queue.size--

	return data, nil
}

func (queue *ArrayImpl) Peek() (interface{}, error) {
	if queue.IsEmpty() {
		return nil, fmt.Errorf("queue is empty")
	}
	return queue.arr[queue.Size()-1], nil
}

func (queue *ArrayImpl) elementAt(index int) (interface{}, error) {
	if index < 0 || index >= queue.size {
		return nil, fmt.Errorf("index out of bound")
	}
	if len(queue.arr) == 0 {
		return nil, fmt.Errorf("stack is empty")
	}

	return queue.arr[index], nil
}

func (queue *ArrayImpl) Size() int {
	return queue.size
}

func (queue *ArrayImpl) IsEmpty() bool {
	return queue.Size() == 0
}

func (queue *ArrayImpl) CreateIterator() standard.Iterator {
	return newQueueIterator(queue)
}

func (queue *ArrayImpl) checkType(ele interface{}) bool {
	if len(queue.arr) == 0 {
		return true
	}
	return reflect.TypeOf(ele) == reflect.TypeOf(queue.arr[0])
}

var _ Queue = (*ArrayImpl)(nil)
var _ standard.Collection = (*ArrayImpl)(nil)

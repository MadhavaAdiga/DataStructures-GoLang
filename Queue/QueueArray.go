package queue

import (
	"fmt"
	"reflect"
)

/*
ArrayImpl is an array implementation of Queue
*/
type ArrayImpl struct {
	element []interface{}
}

/*
Enqueue is method to insert element into a queue
*/
func (queue *ArrayImpl) Enqueue(value interface{}) error {
	//type check
	if len(queue.element) > 0 {
		if reflect.TypeOf(queue.element[0]) != reflect.TypeOf(value) {
			return fmt.Errorf("not same type \n required type : %s", reflect.TypeOf(queue.element[0]))
		}
	}

	queue.element = append(queue.element, value)
	return nil
}

/*
Dequeue is method to remove the first element of a queue
*/
func (queue *ArrayImpl) Dequeue() (interface{}, error) {
	if queue.Size() == 0 {
		return nil, fmt.Errorf("No value in queue to dequeue")
	}

	value := queue.element[0]

	queue.element = queue.element[1:]

	return value, nil
}

/*
Poll is method to remove the first element of a queue
*/
func (queue *ArrayImpl) Poll() (interface{}, error) {
	if queue.Size() == 0 {
		return nil, fmt.Errorf("No value in queue to poll")
	}
	return queue.Dequeue()
}

/*
Size is method to return size of queue
*/
func (queue *ArrayImpl) Size() int {
	return len(queue.element)
}

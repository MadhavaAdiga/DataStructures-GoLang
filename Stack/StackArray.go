package stack

import (
	"fmt"
	"reflect"
)

/*
ArrayImpl is an array implementation of Stack
*/
type ArrayImpl struct {
	element []interface{}
}

/*
Push is a method to insert an element on top of stack
*/
func (stack *ArrayImpl) Push(element interface{}) error {
	//type check
	if len(stack.element) > 0 {
		if reflect.TypeOf(stack.element[0]) != reflect.TypeOf(element) {
			return fmt.Errorf("not same type \n required type : %s", reflect.TypeOf(stack.element[0]))
		}
	}

	stack.element = append(stack.element, element)
	return nil
}

/*
Pop is a method to remove an element from top of stack
*/
func (stack *ArrayImpl) Pop() (interface{}, error) {
	n := len(stack.element)
	value := stack.element[n-1]

	stack.element = stack.element[:n-1]

	return value, nil
}

/*
Size is a method which returns size of the stack
*/
func (stack *ArrayImpl) Size() int {
	return len(stack.element)
}

/*
Peek is a method which returns top element of stack
*/
func (stack ArrayImpl) Peek() interface{} {
	return stack.element[stack.Size()-1]
}

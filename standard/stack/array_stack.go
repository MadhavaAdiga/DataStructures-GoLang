package stack

import (
	"fmt"
	"reflect"

	"com.github/MadhavaAdiga/GolangDS/standard"
)

type ArrayImpl struct {
	size int
	arr  []interface{}
}

func NewArrayStack() Stack {
	return &ArrayImpl{
		size: 0,
		arr:  make([]interface{}, 0),
	}
}

func (stack *ArrayImpl) Push(ele interface{}) error {
	if !stack.checkType(ele) {
		return fmt.Errorf("type mismatch")
	}

	stack.arr = append(stack.arr, ele)
	stack.size++
	return nil
}

func (stack *ArrayImpl) Pop() (interface{}, error) {
	if len(stack.arr) == 0 {
		return nil, fmt.Errorf("stack is empty")
	}

	data := stack.arr[stack.Size()-1]
	stack.arr = stack.arr[:stack.Size()-1]

	stack.size--

	return data, nil
}

func (stack *ArrayImpl) Peek() (interface{}, error) {
	if len(stack.arr) == 0 {
		return nil, fmt.Errorf("stack is empty")
	}

	return stack.arr[stack.Size()-1], nil
}

func (stack *ArrayImpl) IsEmpty() bool {
	return stack.Size() == 0
}

func (stack *ArrayImpl) Size() int {
	return stack.size
}

func (stack *ArrayImpl) elementAt(index int) (interface{}, error) {
	if index < 0 || index >= stack.size {
		return nil, fmt.Errorf("index out of bound")
	}
	if len(stack.arr) == 0 {
		return nil, fmt.Errorf("stack is empty")
	}

	return stack.arr[index], nil
}

func (stack *ArrayImpl) CreateIterator() standard.Iterator { return newStackIterator(stack) }

func (stack *ArrayImpl) checkType(ele interface{}) bool {
	if len(stack.arr) == 0 {
		return true
	}
	return reflect.TypeOf(ele) == reflect.TypeOf(stack.arr[0])
}

var _ Stack = (*ArrayImpl)(nil)
var _ standard.Collection = (*ArrayImpl)(nil)

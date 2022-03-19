package list

import (
	"fmt"
	"reflect"
)

const defaultCapacity = 10

type ArrayList struct {
	capacity uint32
	len      uint32
	ele      []interface{}
}

func NewArrayList(capacity uint32) *ArrayList {

	return &ArrayList{
		capacity: capacity,
		len:      0,
	}
}

func (list *ArrayList) Add(ele interface{}) error {
	fmt.Println(list.ele)
	if list.IsEmpty() {
		if reflect.TypeOf(ele) != reflect.TypeOf(list.ele[0]) {
			return fmt.Errorf("type does not match given %v required %v", reflect.TypeOf(ele), reflect.TypeOf(list.ele[0]))
		}
	}

	if list.len+1 >= list.capacity {
		newArray := make([]interface{}, list.capacity+defaultCapacity)
		copy(newArray, list.ele)
		list.ele = newArray
	}
	list.ele[list.len] = ele
	list.len++
	return nil
}

func (list *ArrayList) Size() uint32 {
	return list.len
}

func (list *ArrayList) Clear() {
	list.len = 0
	list.ele = make([]interface{}, defaultCapacity)
}

func (list *ArrayList) IsEmpty() bool {
	return len(list.ele) != 0
}

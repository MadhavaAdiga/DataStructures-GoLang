package heap

import (
	"fmt"
	"reflect"

	"com.github/MadhavaAdiga/GolangDS/standard/internal/utils"
)

//MinHeap - A binary min heap
type BinaryHeap struct {
	size       int
	arr        []interface{}
	comparator utils.Comparator
}

func NewBinaryHeap(arr []interface{}) Heap {
	n := len(arr)
	heap := &BinaryHeap{
		size: n,
	}

	if n != 0 {
		heap.arr = arr
		heap.comparator = utils.GetComparator(arr[0])
		// heapify
		for i := utils.Max(0, (n/2)-1); i >= 0; i-- {
			heap.sink(i)
		}
	}
	return heap
}

// NewMinHeapWithComparator - use when required to pass a custom comparator
// ex: creating a heap containg a struct
// note: use NewMinHeap() when heap contains primitive types
func NewBinaryHeapWithComparator(arr []interface{}, comparator utils.Comparator) Heap {
	n := len(arr)
	heap := &BinaryHeap{
		size: n,
	}

	heap.comparator = comparator

	if n != 0 {
		heap.arr = arr
		// heapify
		for i := utils.Max(0, (n/2)-1); i >= 0; i-- {
			heap.sink(i)
		}
	}
	return heap
}

func (heap *BinaryHeap) Add(ele interface{}) error {
	if !heap.typeCheck(ele) {
		return fmt.Errorf("type mismatch")
	}

	if heap.comparator == nil {
		heap.comparator = utils.GetComparator(ele)
	}

	heap.arr = append(heap.arr, ele)
	heap.size++
	heap.swim(heap.Size() - 1)

	return nil
}

func (heap *BinaryHeap) Poll() (interface{}, error) { return heap.removeAt(0) }

func (heap *BinaryHeap) Peek() (interface{}, error) {
	if heap.IsEmpty() {
		return nil, fmt.Errorf("heap is empty")
	}

	return heap.arr[0], nil
}

func (heap *BinaryHeap) Remove(ele interface{}) (interface{}, error) {
	if !heap.typeCheck(ele) {
		return nil, fmt.Errorf("type mismatch")
	}

	for i, v := range heap.arr {
		if heap.comparator(v, ele) == 0 {
			return heap.removeAt(i)
		}
	}

	return nil, fmt.Errorf("%v not found in the heap", ele)
}

func (heap *BinaryHeap) removeAt(k int) (interface{}, error) {
	// if k < 0 || k >= heap.Size(){
	// 	return nil, fmt.Errorf("index out of bound")
	// }
	if heap.IsEmpty() {
		return nil, fmt.Errorf("heap is empty")
	}

	lastIndex := heap.Size() - 1
	data := heap.arr[k]
	// move the kth node to last position[swap with last node]
	heap.swap(k, lastIndex)

	heap.arr = heap.arr[:lastIndex]
	heap.size--

	if k == lastIndex {
		return data, nil
	}

	ele := heap.arr[k]

	heap.sink(k)

	if heap.comparator(heap.arr[k], ele) == 0 {
		heap.swim(k)
	}

	return data, nil
}

func (heap *BinaryHeap) Size() int { return heap.size }

func (heap *BinaryHeap) IsEmpty() bool { return heap.size == 0 }

func (heap *BinaryHeap) Contains(ele interface{}) bool {
	for _, v := range heap.arr {
		if v == ele {
			return true
		}
	}
	return false
}

func (heap *BinaryHeap) Clear() {
	heap.arr = nil
	heap.size = 0
	heap.comparator = nil
}

func (heap *BinaryHeap) sink(i int) {
	size := len(heap.arr)

	for true {
		left := (2 * i) + 1
		right := (2 * i) + 2
		k := left

		if right < size && heap.isLess(right, left) {
			k = right
		}

		if left >= size || heap.isLess(i, k) {
			break
		}

		heap.swap(k, i)
		i = k
	}
}

func (heap *BinaryHeap) swim(k int) {
	parent := (k - 1) / 2

	// iterate while kth node is less than parent node
	for k > 0 && heap.isLess(k, parent) {
		// swap parent with k
		heap.swap(parent, k)

		k = parent
		parent = (k - 1) / 2
	}
}

func (heap *BinaryHeap) swap(i, j int) {
	heap.arr[i], heap.arr[j] = heap.arr[j], heap.arr[i]
}

func (heap *BinaryHeap) isLess(a, b int) bool {
	left := heap.arr[a]
	right := heap.arr[b]

	return heap.comparator(left, right) <= 0
}

func (heap *BinaryHeap) typeCheck(ele interface{}) bool {
	if len(heap.arr) == 0 {
		return true
	}

	return reflect.TypeOf(heap.arr[0]) == reflect.TypeOf(ele)
}

var _ Heap = (*BinaryHeap)(nil)

package heap_test

import (
	"fmt"
	"testing"

	"com.github/MadhavaAdiga/GolangDS/standard/heap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {

	var a []interface{} = []interface{}{5, 3, 4, 2, 1}

	n := heap.NewBinaryHeap(a)
	v, _ := n.Peek()
	assert.Equal(t, 1, v)

	fmt.Println(n.Contains(5))
}

func TestCreateWithComparator(t *testing.T) {

	type ex struct {
		a int
		b string
	}

	var a []interface{} = []interface{}{ex{a: 5, b: "a"}, ex{a: 6, b: "b"}}

	n := heap.NewBinaryHeapWithComparator(a, func(a, b interface{}) int {
		a1 := a.(ex)
		b1 := b.(ex)

		if a1.b > b1.b {
			return 1
		} else if a1.b < b1.b {
			return -1
		} else {
			return 0
		}
	})
	v, _ := n.Peek()
	assert.Equal(t, ex{a: 5, b: "a"}, v)
}

func TestAdd(t *testing.T) {
	h := heap.NewBinaryHeap(nil)

	h.Add(4)

	assert.Error(t, h.Add(5.5))

	h.Add(7)
	h.Add(2)
	h.Add(5)
	h.Add(1)
	h.Add(4)
	h.Add(-1)

	v, err := h.Peek()
	assert.NoError(t, err)
	assert.Equal(t, -1, v)

	assert.Equal(t, 7, h.Size())

}

func TestMaxHeap(t *testing.T) {
	h := heap.NewBinaryHeapWithComparator(nil, func(a, b interface{}) int {
		a1 := a.(int)
		b1 := b.(int)

		if b1 > a1 {
			return 1
		} else if b1 < a1 {
			return -1
		} else {
			return 0
		}
	})

	h.Add(4)
	h.Add(-1)
	h.Add(2)
	h.Add(5)
	h.Add(7)
	h.Add(4)
	h.Add(-1)

	v, err := h.Peek()
	assert.NoError(t, err)
	assert.Equal(t, 7, v)

	v, _ = h.Poll()
	assert.Equal(t, 7, v)
	assert.Equal(t, 6, h.Size())

	v, _ = h.Poll()
	assert.Equal(t, 5, v)
	assert.Equal(t, 5, h.Size())
}

func TestPoll(t *testing.T) {
	h := heap.NewBinaryHeap(nil)

	v, err := h.Poll()
	require.Error(t, err)

	h.Add(4)
	h.Add(7)
	h.Add(2)
	h.Add(5)
	h.Add(1)
	h.Add(4)
	h.Add(-1)

	v, _ = h.Peek()
	assert.Equal(t, -1, v)

	v, _ = h.Poll()
	assert.Equal(t, -1, v)
	assert.Equal(t, 6, h.Size())

	v, _ = h.Poll()
	assert.Equal(t, 1, v)
	assert.Equal(t, 5, h.Size())
}

func TestRemove(t *testing.T) {
	h := heap.NewBinaryHeap(nil)

	v, err := h.Peek()
	require.Error(t, err)

	h.Add(4)
	h.Add(7)
	h.Add(2)
	h.Add(5)

	_, err = h.Remove(0)
	require.Error(t, err, "0 not found in the heap")

	_, err = h.Remove(0.3)
	require.Error(t, err, "type mismatch")

	v, _ = h.Remove(7)
	require.Equal(t, 7, v)
	assert.Equal(t, 3, h.Size())
}

func TestContains(t *testing.T) {
	h := heap.NewBinaryHeap(nil)

	_, err := h.Peek()
	require.Error(t, err)

	b := h.Contains(1)
	require.False(t, b)

	h.Add(4)
	h.Add(7)
	h.Add(2)
	h.Add(5)

	b = h.Contains(0)
	require.False(t, b)

	b = h.Contains(7)
	require.True(t, b)
}

func TestClear(t *testing.T) {
	h := heap.NewBinaryHeap(nil)

	h.Add(4)
	h.Add(7)
	h.Add(2)
	h.Add(5)

	h.Clear()

	_, err := h.Peek()
	require.Error(t, err)
}

package linkedlist_test

import (
	"testing"

	linkedlist "com.github/MadhavaAdiga/GolangDS/standard/linkedList"
	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	list := createLinkedList()

	require.NoError(t, list.Add(5))
	require.Error(t, list.Add("adb"))

	require.Equal(t, 1, list.Size())

	a, _ := list.GetHead()
	b, _ := list.GetTail()
	require.Equal(t, 5, a)
	require.Equal(t, 5, b)
}

func TestAddFirst(t *testing.T) {
	list := createLinkedList()

	require.NoError(t, list.Add(5))
	require.Error(t, list.Add("adb"))

	v, _ := list.GetHead()
	require.Equal(t, 5, v)

	require.NoError(t, list.AddFirst(25))
	require.Error(t, list.AddFirst("adb"))
	v, _ = list.GetHead()
	require.Equal(t, 25, v)

	v, _ = list.GetTail()
	require.Equal(t, 5, v)

	require.Equal(t, 2, list.Size())

}

func TestRemoveFirst(t *testing.T) {
	list := createLinkedList()

	_, err := list.RemoveFirst()
	require.Error(t, err)

	list.Add(5)
	list.Add(15)
	list.Add(25)

	require.Equal(t, 3, list.Size())
	v, _ := list.GetHead()
	require.Equal(t, 5, v)

	v, err = list.RemoveFirst()
	require.Equal(t, 5, v)
	require.NoError(t, err)

	require.Equal(t, 2, list.Size())
}

func TestRemoveLast(t *testing.T) {
	list := createLinkedList()

	_, err := list.Remove()
	require.Error(t, err)

	list.Add(5)
	list.Add(15)
	list.Add(25)

	require.Equal(t, 3, list.Size())
	v, _ := list.GetTail()
	require.Equal(t, 25, v)

	v, err = list.Remove()
	require.Equal(t, 25, v)
	require.NoError(t, err)

	require.Equal(t, 2, list.Size())
}

func TestRemoveAt(t *testing.T) {
	list := createLinkedList()

	_, err := list.RemoveAt(0)
	require.Error(t, err, "index is out of bound")

	list.Add(5)
	list.Add(15)
	list.Add(25)
	list.Add(35)
	list.Add(45)

	require.Equal(t, 5, list.Size())

	v, err := list.RemoveAt(3)
	require.Equal(t, 35, v)
	require.Nil(t, err)
	require.Equal(t, 4, list.Size())

	v, err = list.RemoveAt(3)
	require.Equal(t, 45, v)
	require.Nil(t, err)
	require.Equal(t, 3, list.Size())
}

func TestIndexOf(t *testing.T) {
	list := createLinkedList()

	_, err := list.IndexOf(100)
	require.Error(t, err, "empty list")

	list.Add(5)
	list.Add(15)
	list.Add(25)
	list.Add(35)
	list.Add(45)

	require.Equal(t, 5, list.Size())

	index, err := list.IndexOf(100)
	require.Error(t, err)
	require.Equal(t, -1, index)

	index, err = list.IndexOf("sa")
	require.Error(t, err, "type mismatch")
	require.Equal(t, -1, index)

	index, _ = list.IndexOf(25)
	require.Equal(t, 2, index)
}

func TestContains(t *testing.T) {
	list := createLinkedList()

	list.Add(5)
	list.Add(15)
	list.Add(25)
	list.Add(35)
	list.Add(45)

	require.False(t, list.Contains(3))
	require.True(t, list.Contains(5))
}

func TestIterator(t *testing.T) {
	list := createLinkedList()

	list.Add(5)
	list.Add(15)
	list.Add(25)
	list.Add(35)
	list.Add(45)

	iterator := list.CreateIterator()

	var actual [5]int
	i := 0
	for iterator.HasNext() {
		actual[i] = iterator.GetNext().(int)
		i++
	}

	expected := []int{5, 15, 25, 35, 45}

	for i := range expected {
		require.Equal(t, expected[i], actual[i])
	}
}

func createLinkedList() linkedlist.LinkedList {
	// return linkedlist.NewDLinkedList()
	return linkedlist.NewSLinkedList()
}

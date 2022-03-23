package queue_test

import (
	"testing"

	"com.github/MadhavaAdiga/GolangDS/standard/queue"
	"github.com/stretchr/testify/require"
)

func TestEnqueue(t *testing.T) {
	queue := createQueue()

	v, err := queue.Peek()
	require.Error(t, err)
	require.Nil(t, v)

	require.NoError(t, queue.Enqueue(6))
	v, _ = queue.Peek()
	require.Equal(t, 6, v)
}

func TestDequeue(t *testing.T) {
	queue := createQueue()

	queue.Enqueue(5)
	queue.Enqueue(15)
	queue.Enqueue(25)
	queue.Enqueue(35)
	queue.Enqueue(45)

	require.Equal(t, 5, queue.Size())

	v, err := queue.Dequeue()
	require.Nil(t, err)
	require.Equal(t, 5, v)

	require.Equal(t, 4, queue.Size())
}

func TestIterator(t *testing.T) {
	queue := createQueue()

	queue.Enqueue(5)
	queue.Enqueue(15)
	queue.Enqueue(25)
	queue.Enqueue(35)
	queue.Enqueue(45)

	require.Equal(t, 5, queue.Size())

	expected := [5]int{5, 15, 25, 35, 45}
	var actual [5]int

	iterator := queue.CreateIterator()
	i := 0
	for iterator.HasNext() {
		actual[i] = iterator.GetNext().(int)
		i++
	}

	for i := range expected {
		require.Equal(t, expected[i], actual[i])
	}
}

func createQueue() queue.Queue {
	// return queue.NewListQueue()
	return queue.NewArrayQueue()
}

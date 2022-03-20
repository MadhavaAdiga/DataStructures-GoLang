package stack_test

import (
	"testing"

	stack "com.github/MadhavaAdiga/GolangDS/standard/stack"
	"github.com/stretchr/testify/require"
)

func TestPush(t *testing.T) {
	stack := createStack()

	require.NoError(t, stack.Push(5))
	require.Error(t, stack.Push("123"))

	v, _ := stack.Peek()
	require.Equal(t, 5, v)

	require.NoError(t, stack.Push(6))
	v, _ = stack.Peek()
	require.Equal(t, 6, v)
}

func TestPop(t *testing.T) {
	stack := createStack()

	stack.Push(5)
	stack.Push(15)
	stack.Push(25)
	stack.Push(35)
	stack.Push(45)

	require.Equal(t, 5, stack.Size())

	v, err := stack.Pop()
	require.Nil(t, err)
	require.Equal(t, 45, v)

	require.Equal(t, 4, stack.Size())
}

func TestPeek(t *testing.T) {
	stack := createStack()

	v, err := stack.Peek()
	require.Error(t, err)
	require.Nil(t, v)

	require.NoError(t, stack.Push(6))
	v, _ = stack.Peek()
	require.Equal(t, 6, v)
}

func TestIterator(t *testing.T) {
	stack := createStack()

	stack.Push(5)
	stack.Push(15)
	stack.Push(25)
	stack.Push(35)
	stack.Push(45)

	require.Equal(t, 5, stack.Size())

	expected := [5]int{45, 35, 25, 15, 5}
	var actual [5]int

	iterator := stack.CreateIterator()
	i := 0
	for iterator.HasNext() {
		actual[i] = iterator.GetNext().(int)
		i++
	}

	for i := range expected {
		require.Equal(t, expected[i], actual[i])
	}

}

func createStack() stack.Stack {
	return stack.NewListStack()
}

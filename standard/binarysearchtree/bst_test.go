package binarysearchtree_test

import (
	"testing"

	"com.github/MadhavaAdiga/GolangDS/standard/binarysearchtree"
	"github.com/stretchr/testify/require"
)

// func TestDemo(t *testing.T) {

// 	bst := binarysearchtree.GetBst(binarysearchtree.PRIMITIVE)

// 	require.True(t, bst.IsEmpty())
// 	require.True(t, bst.Add(3))

// 	require.False(t, bst.IsEmpty())

// 	require.False(t, bst.Add(3))

// 	require.True(t, bst.Add(5))
// 	require.False(t, bst.Add(5))
// }

func TestAdd(t *testing.T) {
	bst := binarysearchtree.GetBst(binarysearchtree.PRIMITIVE)

	bst.Add(5)

	require.True(t, bst.Contains(5))
	bst.Add(4)
	require.True(t, bst.Contains(4))
	bst.Add(3)
	require.True(t, bst.Contains(3))
	bst.Add(9)
	require.True(t, bst.Contains(9))
	bst.Add(7)
	require.True(t, bst.Contains(7))
	bst.Add(6)
	require.True(t, bst.Contains(6))
	bst.Add(8)
	require.True(t, bst.Contains(8))

	require.Equal(t, 7, bst.Size())

}

func TestRemove(t *testing.T) {
	bst := binarysearchtree.GetBst(binarysearchtree.PRIMITIVE)

	bst.Add(5)

	require.True(t, bst.Contains(5))
	bst.Add(4)
	require.True(t, bst.Contains(4))
	bst.Add(3)
	require.True(t, bst.Contains(3))
	bst.Add(9)
	require.True(t, bst.Contains(9))
	bst.Add(7)
	require.True(t, bst.Contains(7))
	bst.Add(6)
	require.True(t, bst.Contains(6))
	bst.Add(8)
	require.True(t, bst.Contains(8))

	require.Equal(t, 7, bst.Size())

	require.True(t, bst.Contains(9))

	require.True(t, bst.Remove(9))
	require.Equal(t, 6, bst.Size())
	require.False(t, bst.Contains(9))

	require.True(t, bst.Contains(6))

	require.True(t, bst.Remove(6))
	require.Equal(t, 5, bst.Size())
	require.False(t, bst.Contains(6))

	require.True(t, bst.Remove(5))
	require.Equal(t, 4, bst.Size())
	require.False(t, bst.Contains(5))
}

func TestContains(t *testing.T) {
	bst := binarysearchtree.GetBst(binarysearchtree.PRIMITIVE)

	require.False(t, bst.Contains(0))

	bst.Add(5)

	require.True(t, bst.Contains(5))
	bst.Add(4)
	require.True(t, bst.Contains(4))
	bst.Add(3)
	require.True(t, bst.Contains(3))

	require.False(t, bst.Contains(0))
}

package bst_test

import (
	"testing"

	"com.github/MadhavaAdiga/GolangDS/standard/internal/types"
	"com.github/MadhavaAdiga/GolangDS/standard/internal/utils"
	binarysearchtree "com.github/MadhavaAdiga/GolangDS/standard/tree/bst"
	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	// bst := binarysearchtree.GetBst(binarysearchtree.PRIMITIVE)
	bst := binarysearchtree.NewBstImpl()

	stub := []int{5, 4, 3, 9, 7, 6, 8}

	for _, v := range stub {
		a1 := createComparable(v)
		bst.Add(*a1)
		require.True(t, bst.Contains(*a1))
	}

	require.Equal(t, len(stub), bst.Size())

}

func TestRemove(t *testing.T) {
	// bst := binarysearchtree.GetBst(binarysearchtree.PRIMITIVE)

	bst := binarysearchtree.NewBstImpl()

	a := *createComparable(0)

	bst.Add(a)
	require.True(t, bst.Contains(a))
	require.True(t, bst.Remove(a))
	require.Equal(t, 0, bst.Size())
	require.False(t, bst.Contains(a))

	stub := []int{5, 4, 3, 9, 7, 6, 8}
	var comaprables []types.Comparable

	for _, v := range stub {
		a1 := createComparable(v)
		comaprables = append(comaprables, *a1)

		bst.Add(*a1)
		require.True(t, bst.Contains(*a1))
	}
	require.Equal(t, len(stub), bst.Size())

	// index no to delete
	sequence := []int{3, 5, 0}

	for i, v := range sequence {
		require.True(t, bst.Remove(comaprables[v]))
		require.Equal(t, len(stub)-(i+1), bst.Size())
		require.False(t, bst.Contains(comaprables[v]))
	}

	// bst.Add(5)

	// require.True(t, bst.Contains(5))
	// bst.Add(4)
	// require.True(t, bst.Contains(4))
	// bst.Add(3)
	// require.True(t, bst.Contains(3))
	// bst.Add(9)
	// require.True(t, bst.Contains(9))
	// bst.Add(7)
	// require.True(t, bst.Contains(7))
	// bst.Add(6)
	// require.True(t, bst.Contains(6))
	// bst.Add(8)
	// require.True(t, bst.Contains(8))

	// require.Equal(t, 7, bst.Size())

	// require.True(t, bst.Contains(9))

	// require.True(t, bst.Remove(comaprables[3]))
	// require.Equal(t, 6, bst.Size())
	// require.False(t, bst.Contains(comaprables[3]))

	// require.True(t, bst.Contains(comaprables[5]))

	// require.True(t, bst.Remove(6))
	// require.Equal(t, 5, bst.Size())
	// require.False(t, bst.Contains(6))

	// require.True(t, bst.Remove(5))
	// require.Equal(t, 4, bst.Size())
	// require.False(t, bst.Contains(5))
}

func TestContains(t *testing.T) {
	// bst := binarysearchtree.GetBst(binarysearchtree.PRIMITIVE)

	bst := binarysearchtree.NewBstImpl()

	require.False(t, bst.Contains(*createComparable(0)))

	// bst.Add(5)

	stub := []int{5, 4, 3, 9, 7, 6, 8}

	for _, v := range stub {
		a1 := createComparable(v)
		bst.Add(*a1)
		require.True(t, bst.Contains(*a1))
	}

	require.False(t, bst.Contains(*createComparable(10)))

	// require.True(t, bst.Contains(5))
	// bst.Add(4)
	// require.True(t, bst.Contains(4))
	// bst.Add(3)
	// require.True(t, bst.Contains(3))

	// require.False(t, bst.Contains(0))
}

func BenchmarkBst(b *testing.B) {
	// bst := binarysearchtree.GetBst(binarysearchtree.PRIMITIVE)

	bst := binarysearchtree.NewBstImpl()

	b.Run("", func(b *testing.B) {
		stub := []int{5, 4, 3, 9, 7, 6, 8}
		var comaprables []types.Comparable

		for _, v := range stub {
			a1 := createComparable(v)
			comaprables = append(comaprables, *a1)

			bst.Add(*a1)
			bst.Contains(*a1)
		}
		bst.Size()

		// index no to delete
		sequence := []int{3, 5, 0}

		for _, v := range sequence {
			bst.Remove(comaprables[v])
		}

	})

}

func createComparable(i int) *types.Comparable {
	c1 := utils.GetComparator(i)
	a := func(c types.Comparable) int32 {
		return int32(c1(i, c.Val))
	}

	return types.NewComparable(i, a)
}

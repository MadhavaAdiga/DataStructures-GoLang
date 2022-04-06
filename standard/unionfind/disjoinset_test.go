package unionfind_test

import (
	"testing"

	"com.github/MadhavaAdiga/GolangDS/standard/unionfind"
	"github.com/stretchr/testify/require"
)

func TestComponent(t *testing.T) {

	ds := unionfind.NewDisjoinSet(5)
	require.Equal(t, 5, ds.Components())

	ds.Union(0, 1)
	require.Equal(t, 4, ds.Components())

	ds.Union(1, 0)
	require.Equal(t, 4, ds.Components())

	ds.Union(1, 2)
	require.Equal(t, 3, ds.Components())

	ds.Union(0, 2)
	require.Equal(t, 3, ds.Components())

	ds.Union(4, 3)
	require.Equal(t, 2, ds.Components())

	ds.Union(0, 4)
	require.Equal(t, 1, ds.Components())

	ds.Union(1, 3)
	require.Equal(t, 1, ds.Components())

	ds.Union(4, 7)
	require.Equal(t, 1, ds.Components())
}

func TestComponentSize(t *testing.T) {
	ds := unionfind.NewDisjoinSet(5)
	for i := 0; i < ds.Size(); i++ {
		require.Equal(t, 1, ds.ComponentSize(i))
	}

	ds.Union(0, 1)
	require.Equal(t, 2, ds.ComponentSize(0))
	require.Equal(t, 2, ds.ComponentSize(1))
	require.Equal(t, 1, ds.ComponentSize(2))
	require.Equal(t, 1, ds.ComponentSize(3))
	require.Equal(t, 1, ds.ComponentSize(4))

	ds.Union(1, 0)
	require.Equal(t, 2, ds.ComponentSize(0))
	require.Equal(t, 2, ds.ComponentSize(1))
	require.Equal(t, 1, ds.ComponentSize(2))
	require.Equal(t, 1, ds.ComponentSize(3))
	require.Equal(t, 1, ds.ComponentSize(4))

	ds.Union(1, 3)
	require.Equal(t, 3, ds.ComponentSize(0))
	require.Equal(t, 3, ds.ComponentSize(1))
	require.Equal(t, 1, ds.ComponentSize(2))
	require.Equal(t, 3, ds.ComponentSize(3))
	require.Equal(t, 1, ds.ComponentSize(4))

	ds.Union(2, 4)
	require.Equal(t, 3, ds.ComponentSize(0))
	require.Equal(t, 3, ds.ComponentSize(1))
	require.Equal(t, 2, ds.ComponentSize(2))
	require.Equal(t, 3, ds.ComponentSize(3))
	require.Equal(t, 2, ds.ComponentSize(4))

	ds.Union(0, 4)
	require.Equal(t, 5, ds.ComponentSize(0))
	require.Equal(t, 5, ds.ComponentSize(1))
	require.Equal(t, 5, ds.ComponentSize(2))
	require.Equal(t, 5, ds.ComponentSize(3))
	require.Equal(t, 5, ds.ComponentSize(4))

	ds.Union(0, 2)
	require.Equal(t, 5, ds.ComponentSize(0))
	require.Equal(t, 5, ds.ComponentSize(1))
	require.Equal(t, 5, ds.ComponentSize(2))
	require.Equal(t, 5, ds.ComponentSize(3))
	require.Equal(t, 5, ds.ComponentSize(4))

	ds.Union(1, 4)
	require.Equal(t, 5, ds.ComponentSize(0))
	require.Equal(t, 5, ds.ComponentSize(1))
	require.Equal(t, 5, ds.ComponentSize(2))
	require.Equal(t, 5, ds.ComponentSize(3))
	require.Equal(t, 5, ds.ComponentSize(4))

	ds.Union(3, 4)
	require.Equal(t, 5, ds.ComponentSize(0))
	require.Equal(t, 5, ds.ComponentSize(1))
	require.Equal(t, 5, ds.ComponentSize(2))
	require.Equal(t, 5, ds.ComponentSize(3))
	require.Equal(t, 5, ds.ComponentSize(4))
}

func TestSize(t *testing.T) {
	ds := unionfind.NewDisjoinSet(5)
	require.Equal(t, 5, ds.Size())

	ds.Union(0, 1)
	require.Equal(t, 5, ds.Size())
	ds.Union(2, 1)
	require.Equal(t, 5, ds.Size())
	ds.Union(3, 1)
	require.Equal(t, 5, ds.Size())
	ds.Union(0, 3)
	require.Equal(t, 5, ds.Size())
	ds.Union(4, 1)
	require.Equal(t, 5, ds.Size())
}

func TestConectivity(t *testing.T) {
	ds := unionfind.NewDisjoinSet(5)
	require.Equal(t, 5, ds.Components())

	ds.Union(0, 1)
	require.True(t, ds.Connected(0, 1))
	require.True(t, ds.Connected(1, 0))
	require.True(t, ds.Connected(0, 0))

	require.False(t, ds.Connected(2, 1))
	require.False(t, ds.Connected(3, 1))
	require.False(t, ds.Connected(2, 0))
	require.False(t, ds.Connected(4, 1))

	ds.Union(1, 0)
	require.True(t, ds.Connected(0, 1))
	require.True(t, ds.Connected(1, 0))
	require.True(t, ds.Connected(0, 0))

	require.False(t, ds.Connected(2, 1))
	require.False(t, ds.Connected(3, 1))
	require.False(t, ds.Connected(2, 0))
	require.False(t, ds.Connected(4, 1))

	ds.Union(4, 1)
	require.True(t, ds.Connected(0, 1))
	require.True(t, ds.Connected(1, 0))
	require.True(t, ds.Connected(4, 0))
	require.True(t, ds.Connected(4, 1))

	require.False(t, ds.Connected(2, 1))
	require.False(t, ds.Connected(3, 1))
	require.False(t, ds.Connected(2, 0))

	ds.Union(2, 3)
	require.True(t, ds.Connected(0, 1))
	require.True(t, ds.Connected(1, 0))
	require.True(t, ds.Connected(4, 0))
	require.True(t, ds.Connected(4, 1))
	require.True(t, ds.Connected(3, 2))
	require.True(t, ds.Connected(2, 3))

	require.False(t, ds.Connected(2, 1))
	require.False(t, ds.Connected(3, 1))
	require.False(t, ds.Connected(2, 0))

	ds.Union(4, 3)
	require.True(t, ds.Connected(0, 1))
	require.True(t, ds.Connected(1, 0))
	require.True(t, ds.Connected(4, 0))
	require.True(t, ds.Connected(4, 1))
	require.True(t, ds.Connected(3, 2))
	require.True(t, ds.Connected(2, 3))
	require.True(t, ds.Connected(2, 1))
	require.True(t, ds.Connected(3, 1))
	require.True(t, ds.Connected(2, 0))
}

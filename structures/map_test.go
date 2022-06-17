package structures

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPearsonHashing(t *testing.T) {

	for _, mapSize := range []int{8, 16, 32, 1024, 2048, 4096, 131072} {

		for _, key := range []string{"", "foo", "bar", "baz", "qux"} {
			idx := PearsonHash(key, mapSize)
			require.Lessf(t, idx, mapSize, "hash func is broken: idx (%d) >= size (%d)", idx, mapSize)
		}
	}
}

func TestMap_InsertAndFind(t *testing.T) {
	m := NewMap()

	m.Insert("foo", 1)

	val, ok := m.Find("foo")
	require.True(t, ok)
	require.Equal(t, 1, val, "expected value = 1, but found %d (because collision and overriding the data)", val)

	_, ok = m.Find("bar")
	require.False(t, ok)
}

func TestMap_Collisions_and_Resizing(t *testing.T) {
	m := NewMap()

	m.Insert("foo", 0)
	m.Insert("foo", 1)
	m.Insert("bar", -2)
	m.Insert("qux", 2)
	m.Insert("quux", 3)
	m.Insert("quuux", 4)
	m.Insert("quuuux", 5)
	m.Insert("quuuuux", 6)
	m.Insert("xyz", 7)

	val, ok := m.Find("foo")
	require.True(t, ok)

	require.Equal(t, 1, val, "expected value = 1, but found %d (because collision and overriding the data)", val)

	require.Equal(t, 8, m.count, "there must be 8 elements, but counter shows %d", m.count)

	require.Equal(t, 16, len(m.buckets), "buckets size expected to be 16")
}

func TestMap_NonASCIIcharacters(t *testing.T) {
	require.NotPanics(t, func() {
		m := NewMap()
		m.Insert("フルーツバスケット", 1)

		val, ok := m.Find("フルーツバスケット")
		require.True(t, ok)
		require.Equal(t, 1, val)
	})
}

func TestMap_Remove(t *testing.T) {
	m := NewMap()

	m.Insert("foo", 1)
	m.Insert("bar", -2)
	m.Insert("qux", 2)
	m.Insert("abc", 3)
	m.Insert("def", 4)
	m.Insert("xyz", 5)

	_, ok := m.Find("xyz")
	require.True(t, ok)

	m.Remove("xyz")

	_, ok = m.Find("xyz")
	require.False(t, ok)
}

// functionality comparison with built-in map
func TestMap_functional(t *testing.T) {
	builtinMap := make(map[string]int)
	m := NewMap()

	for i := 0; i < 100; i++ {
		key := strconv.Itoa(i) + "_suffix"

		builtinMap[key] = i
		m.Insert(key, i)

		require.Equal(t, len(builtinMap), m.count, "different sizes of builtin and implemented maps")
	}

	for key, val := range builtinMap {
		v, ok := m.Find(key)
		require.Truef(t, ok, "key '%s' not found in implemented map", key)
		require.Equalf(t, val, v, "data corrupted. expected: %d, but got: %d", val, v)
	}

	for key := range builtinMap {
		m.Remove(key)

		val, ok := m.Find(key)
		require.Falsef(t, ok, "key '%s' val: %d expected to be deleted", key, val)
	}

	require.Zerof(t, m.count, "impl. map expected to be empty but it contains data (%d items)", m.count)
}

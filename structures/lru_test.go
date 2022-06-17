package structures

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLRUCache(t *testing.T) {
	t.Run("test overflow", func(t *testing.T) {
		cache := NewLRUCache(4)

		cache.set("a", 1)
		cache.set("b", 2)
		cache.set("c", 3)
		cache.set("d", 4)
		// [d c b e] --> [e d c b] a (excluded)
		cache.set("e", 5)

		n := cache.head
		// iterate from the newest to the oldest values
		for _, key := range []string{"e", "d", "c", "b"} {
			require.Equal(t, key, n.key)
			n = n.next
		}

		_, ok := cache.get("a")
		require.False(t, ok)
	})

	t.Run("check value", func(t *testing.T) {
		cache := NewLRUCache(4)

		cache.set("a", 1)
		cache.set("b", 2)
		cache.set("c", 3)
		cache.set("d", 4)

		// [d c b a] --> [a d c b]
		val, ok := cache.get("a")
		require.True(t, ok)

		require.Equal(t, 1, val)

		n := cache.head
		require.Equal(t, "a", n.key)

		for _, key := range []string{"a", "d", "c", "b"} {
			require.Equal(t, key, n.key)
			n = n.next
		}
	})

	t.Run("update value", func(t *testing.T) {
		cache := NewLRUCache(4)

		cache.set("a", 1)
		cache.set("b", 2)
		cache.set("c", 3)
		cache.set("d", 4)
		// d c b a

		// after: a d c b
		cache.get("a")

		// after: c a d b
		cache.set("c", 10)

		val, ok := cache.get("c")
		require.True(t, ok)
		require.Equal(t, 10, val)

		n := cache.head
		require.Equal(t, "c", n.key)

		for _, key := range []string{"c", "a", "d", "b"} {
			require.Equal(t, key, n.key)
			n = n.next
		}
	})
}

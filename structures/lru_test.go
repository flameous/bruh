package structures

import (
	"testing"
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
			if key != n.key {
				t.Fatalf("expected %s but got %s", key, n.key)
			}
			n = n.next
		}

		_, ok := cache.get("a")
		if ok {
			t.Fatal("'a' expected to be deleted from the cache")
		}
	})

	t.Run("check value", func(t *testing.T) {
		cache := NewLRUCache(4)

		cache.set("a", 1)
		cache.set("b", 2)
		cache.set("c", 3)
		cache.set("d", 4)

		// [d c b a] --> [a d c b]
		val, ok := cache.get("a")
		if !ok {
			t.Fatal("expected value")
		}
		if val != 1 {
			t.Fatalf("expected 1 but got %d", val)
		}

		n := cache.head
		if n.key != "a" {
			t.Fatal("'a' expected to be the newest key-value pair in the lru cache")
		}

		for _, key := range []string{"a", "d", "c", "b"} {
			if key != n.key {
				t.Fatalf("expected %s but got %s", key, n.key)
			}
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
		if !ok {
			t.Fatal("expected value")
		}
		if val != 10 {
			t.Fatalf("expected 10 but got %d", val)
		}

		n := cache.head
		if n.key != "c" {
			t.Fatal("'c' expected to be the newest key-value pair in the lru cache")
		}

		for _, key := range []string{"c", "a", "d", "b"} {
			if key != n.key {
				t.Fatalf("expected %s but got %s", key, n.key)
			}
			n = n.next
		}
	})
}

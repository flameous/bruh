package structures

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {

	t.Run("", func(t *testing.T) {

		ll := LinkedList{}

		ll.enqueue(1)
		ll.enqueue(2)
		ll.enqueue(3)
		ll.enqueue(4)
		ll.enqueue(5)
		ll.enqueue(6)

		for i := 1; i <= 6; i++ {
			v, ok := ll.dequeue()
			require.True(t, ok)
			require.Equal(t, i, v)
		}

		_, ok := ll.dequeue()
		require.False(t, ok)
	})
}

func TestListRemove(t *testing.T) {

	t.Run("remove last element in queue", func(t *testing.T) {
		ll := LinkedList{}

		n := ll.enqueue(1)
		n.remove()

		val, ok := ll.dequeue()
		require.Falsef(t, ok, "expected to be empty, but got %v", val)
	})

	t.Run("remove head", func(t *testing.T) {
		ll := LinkedList{}
		n := ll.enqueue(1)
		ll.enqueue(2)
		ll.enqueue(3)

		n.remove()

		val, ok := ll.dequeue()
		require.True(t, ok, "expected non-empty queue")
		require.Equal(t, 2, val)
	})

	t.Run("remove node from the middle of queue", func(t *testing.T) {
		ll := LinkedList{}
		ll.enqueue(1)
		n := ll.enqueue(2)
		ll.enqueue(3)
		ll.enqueue(4)
		ll.enqueue(5)

		n.remove()

		for _, expected := range []int{1, 3, 4, 5} {
			val, ok := ll.dequeue()
			require.True(t, ok, "expected non-empty queue")
			require.Equal(t, expected, val)

		}
	})

	t.Run("remove head (#2)", func(t *testing.T) {
		ll := LinkedList{}
		n := ll.enqueue(1)
		ll.enqueue(2)
		ll.enqueue(3)

		n.remove()

		require.Equal(t, 2, ll.head.val)
		require.Equal(t, 3, ll.tail.val)
	})

	t.Run("remove tail", func(t *testing.T) {
		ll := LinkedList{}
		ll.enqueue(1)
		ll.enqueue(2)
		n := ll.enqueue(3)

		n.remove()

		require.Equal(t, 1, ll.head.val)
		require.Equal(t, 2, ll.tail.val)
	})

	t.Run("remove middle node (#2)", func(t *testing.T) {
		ll := LinkedList{}
		head := ll.enqueue(1)
		n := ll.enqueue(2)
		tail := ll.enqueue(3)

		n.remove()
		require.Same(t, head, ll.head)
		require.Same(t, tail, ll.tail)
	})

	t.Run("node with missing queue", func(t *testing.T) {
		n := &Node{}
		n.remove()
	})
}

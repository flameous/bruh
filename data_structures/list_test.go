package structures

import "testing"

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
			if !ok {
				t.Fatal("expected value")
			}
			if v != i {
				t.Fatalf("expected %d but got %d", i, v)
			}
		}

		_, ok := ll.dequeue()
		if ok {
			t.Fatal("queue must be empty")
		}
	})
}

func TestListRemove(t *testing.T) {

	t.Run("remove last element in queue", func(t *testing.T) {
		ll := LinkedList{}

		n := ll.enqueue(1)
		n.remove()

		val, ok := ll.dequeue()
		if ok {
			t.Error("expected empty queue, got val:", val)
		}
	})

	t.Run("remove head", func(t *testing.T) {
		ll := LinkedList{}
		n := ll.enqueue(1)
		ll.enqueue(2)
		ll.enqueue(3)

		n.remove()

		val, ok := ll.dequeue()
		if !ok {
			t.Error("expected empty queue")
		}
		if val != 2 {
			t.Error("expected 2")
		}
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
			if !ok {
				t.Error("expected value")
			}

			if val != expected {
				t.Errorf("expected value %d but got %d", expected, val)
			}
		}
	})

	t.Run("remove head (#2)", func(t *testing.T) {
		ll := LinkedList{}
		n := ll.enqueue(1)
		ll.enqueue(2)
		ll.enqueue(3)

		n.remove()

		if !(ll.head.val == 2 && ll.tail.val == 3) {
			t.Error("broken!")
		}

	})

	t.Run("remove tail", func(t *testing.T) {
		ll := LinkedList{}
		ll.enqueue(1)
		ll.enqueue(2)
		n := ll.enqueue(3)

		n.remove()

		if !(ll.head.val == 1 && ll.tail.val == 2) {
			t.Error("broken!!")
		}
	})

	t.Run("remove middle node (#2)", func(t *testing.T) {
		ll := LinkedList{}
		head := ll.enqueue(1)
		n := ll.enqueue(2)
		tail := ll.enqueue(3)

		n.remove()
		if !(ll.head == head && ll.tail == tail) {
			t.Error("broken!")
		}
	})

	t.Run("node with missing queue", func(t *testing.T) {
		n := &Node{}
		n.remove()
	})
}

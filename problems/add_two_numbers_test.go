package problems

import (
	"testing"
)

func sumLL(n *ListNode) int {
	s := 0
	mult := 1
	for {
		if n == nil {
			break
		}
		s = s + n.Val*mult
		mult *= 10
		n = n.Next
	}
	return s
}

// helper func
func numToReversedLL(n int) *ListNode {
	head := &ListNode{}

	node := head
	for {
		node.Val = n % 10
		if n/10 == 0 {
			break
		}
		n = n / 10
		node.Next = &ListNode{}
		node = node.Next
	}
	return head
}

func TestAddTwoNumbers(t *testing.T) {

	t.Run("two numbers with same length", func(t *testing.T) {
		num1 := 12
		num2 := 23

		l1 := numToReversedLL(num1)
		l2 := numToReversedLL(num2)

		res := addTwoNumbers(l1, l2)

		expected := num1 + num2
		got := sumLL(res)
		if expected != got {
			t.Fatalf("expected %d, got %d", expected, got)
		}
	})

	t.Run("zero plus number", func(t *testing.T) {
		num1 := 0
		num2 := 123

		l1 := numToReversedLL(num1)
		l2 := numToReversedLL(num2)

		res := addTwoNumbers(l1, l2)

		expected := num1 + num2
		got := sumLL(res)
		if expected != got {
			t.Fatalf("expected %d, got %d", expected, got)
		}
	})

	t.Run("two numbers with different length", func(t *testing.T) {
		num1 := 12123
		num2 := 23

		l1 := numToReversedLL(num1)
		l2 := numToReversedLL(num2)

		res := addTwoNumbers(l1, l2)

		expected := num1 + num2
		got := sumLL(res)
		if expected != got {
			t.Fatalf("expected %d, got %d", expected, got)
		}
	})

	t.Run("overflow", func(t *testing.T) {
		num1 := 100009
		num2 := 345

		l1 := numToReversedLL(num1)
		l2 := numToReversedLL(num2)

		res := addTwoNumbers(l1, l2)

		expected := num1 + num2
		got := sumLL(res)
		if expected != got {
			t.Fatalf("expected %d, got %d", expected, got)
		}
	})

	t.Run("overflow at the last digit", func(t *testing.T) {
		num1 := 90000
		num2 := 10000

		l1 := numToReversedLL(num1)
		l2 := numToReversedLL(num2)

		res := addTwoNumbers(l1, l2)

		expected := num1 + num2
		got := sumLL(res)
		if expected != got {
			t.Fatalf("expected %d, got %d", expected, got)
		}
	})

	t.Run("overflow #3", func(t *testing.T) {
		num1 := 1
		num2 := 79999999

		l1 := numToReversedLL(num1)
		l2 := numToReversedLL(num2)

		res := addTwoNumbers(l1, l2)

		expected := num1 + num2
		got := sumLL(res)
		if expected != got {
			t.Fatalf("expected %d, got %d", expected, got)
		}
	})

}

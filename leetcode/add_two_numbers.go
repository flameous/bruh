package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/add-two-numbers/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := l1

	overflow := 0
	for {
		sum := l1.Val + overflow
		if l2 != nil {
			sum += l2.Val
		}
		l1.Val = sum % 10
		overflow = sum / 10 // 1 or 0 only

		if overflow > 0 {
			// overflow and first list (used as sum storage) ended
			if l1.Next == nil {
				// second list also ended
				if l2 == nil || l2.Next == nil {
					l1.Next = &ListNode{Val: 1}
					break
				}
				// first list ended, so link its last node to the second list's next node
				l1.Next = l2.Next
				l2 = nil
			}
		} else {
			// no overflow, keep iterating over list(s)
			if l1.Next == nil {
				if l2 != nil {
					// first list ended, link it to second list node
					l1.Next = l2.Next
				}
				break
			}
		}
		// shift cursor on the first list
		l1 = l1.Next
		if l2 != nil {
			// and also on the second list if its not nil
			l2 = l2.Next
		}
	}

	return head
}

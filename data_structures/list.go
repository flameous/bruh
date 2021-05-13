package structures

// nil <- head <-> ... <-> ... <-> tail -> nil
type LinkedList struct {
	head *Node
	tail *Node
}

type Node struct {
	list *LinkedList
	next *Node
	prev *Node
	val  int
}

// O(1)
func (ll *LinkedList) enqueue(val int) *Node {
	node := &Node{
		val:  val,
		list: ll,
	}

	if ll.tail == nil {
		ll.tail = node
		ll.head = node
		return node
	}

	ll.tail.next = node
	node.prev = ll.tail
	ll.tail = node
	return node
}

// O(1)
func (ll *LinkedList) dequeue() (int, bool) {
	if ll.head == nil {
		return 0, false
	}

	val := ll.head.val
	if ll.head == ll.tail {
		ll.head = nil
		ll.tail = nil
		return val, true
	}

	ll.head = ll.head.next
	ll.head.prev = nil
	return val, true
}

// O(1)
func (n *Node) remove() {
	list := n.list
	if list == nil {
		return
	}

	if n == list.tail || n == list.head {
		if n == list.tail {
			list.tail = list.tail.prev
		}
		if n == list.head {
			list.head = list.head.next
		}
		return
	}

	n.prev.next = n.next
	n.next.prev = n.prev
}

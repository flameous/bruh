package structures

type lruCache struct {
	m          map[string]*lruNode
	head, tail *lruNode
	size       int
}

type lruNode struct {
	key        string
	val        int
	prev, next *lruNode
}

func NewLRUCache(size int) *lruCache {
	if size < 2 {
		panic("invalid size, bruh")
	}
	return &lruCache{
		m:    make(map[string]*lruNode, size), // btw it can be ompimized a little bit
		size: size,
	}
}

func (l *lruCache) get(key string) (int, bool) {
	node, ok := l.m[key]
	if !ok {
		return 0, false
	}

	l.pushToTop(node)
	return node.val, true
}

func (l *lruCache) set(key string, val int) {
	if node, ok := l.m[key]; ok {
		node.val = val
		l.pushToTop(node)
		return
	}

	node := &lruNode{
		key: key,
		val: val,
	}
	l.m[key] = node

	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}

	l.head.prev = node
	node.next = l.head
	l.head = node

	if len(l.m) > l.size {
		// remove oldest value
		delete(l.m, l.tail.key)
		l.tail = l.tail.prev
		l.tail.next = nil
	}
}

func (l *lruCache) pushToTop(node *lruNode) {
	if node == l.head {
		return
	}

	if node == l.tail {
		l.tail = l.tail.prev
		l.tail.next = nil
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
	l.head.prev = node
	node.next = l.head
	l.head = node
}

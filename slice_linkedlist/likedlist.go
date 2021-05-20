package slice_linkedlist

import "sync"

type LikedList struct {
	head *Node
	tail *Node
	mu   sync.RWMutex
}

type Node struct {
	val  *T
	next *Node
	prev *Node
	mu   sync.RWMutex
}

func (l *LikedList) Append(point *T) {
	l.mu.Lock()
	defer l.mu.Unlock()

	newNode := &Node{
		val: point,
	}
	if l.tail == nil {
		// First insertion
		l.head = newNode
		l.tail = newNode
		return
	}
	newNode.prev = l.tail
	l.tail.next = newNode
	l.tail = newNode
}

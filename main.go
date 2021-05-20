package main

import "sync"

type T struct{}

// Slice
type Slice struct {
	s  []*T
	mu sync.Mutex
}

func (s *Slice) Append(t *T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.s = append(s.s, t)
}

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
		l.head = newNode
		l.tail = newNode
		return
	}
	newNode.prev = l.tail
	l.tail.next = newNode
	l.tail = newNode
}

func main() {
}

package main

type Value struct{}

// Slice
type Slice struct {
	s []*Value
}

func (s *Slice) Append(t *Value) {
	s.s = append(s.s, t)
}

// LinkedList
type LinkedList struct {
	head *Node
	tail *Node
}

type Node struct {
	val  *Value
	next *Node
	prev *Node
}

func (l *LinkedList) Append(v *Value) {
	newNode := &Node{val: v}
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

package main

import "testing"

func BenchmarkSliceAppend(b *testing.B) {
	s := []*Value{}
	for i := 0; i < b.N; i++ {
		s = append(s, &Value{})
	}
}

func BenchmarkLinkedListAppend(b *testing.B) {
	list := &LinkedList{}
	for i := 0; i < b.N; i++ {
		list.Append(&Value{})
	}
}

type Value struct{}

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

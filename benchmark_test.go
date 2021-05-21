package main

import (
	"testing"
)

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

func BenchmarkSliceIterate(b *testing.B) {
	s := []*Value{}
	for i := 0; i < 100000; i++ {
		s = append(s, &Value{id: i})
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := range s {
			v := s[i]
			_ = v
		}
	}
}

func BenchmarkLinkedListIterate(b *testing.B) {
	list := &LinkedList{}
	for i := 0; i < 100000; i++ {
		list.Append(&Value{id: i})
	}
	iterator := list.NewIterator()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for iterator.Next() {
			v := iterator.current.val
			_ = v
		}
	}
}

type Value struct {
	id int
}

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

func (l *LinkedList) NewIterator() *Iterator {
	head := l.head
	// Put a dummy node so that it positions the head on the first Next() call.
	dummy := &Node{
		next: head,
	}
	return &Iterator{
		current: dummy,
	}
}

type Iterator struct {
	current *Node
}

func (i *Iterator) Next() bool {
	if i.current == nil {
		return false
	}
	next := i.current.next
	i.current = next
	return i.current != nil
}

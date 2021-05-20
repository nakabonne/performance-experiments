package main

import "testing"

func BenchmarkSliceAppend(b *testing.B) {
	s := &Slice{s: []*Value{}}
	for i := 0; i < b.N; i++ {
		s.Append(&Value{})
	}
}

func BenchmarkLinkedListAppend(b *testing.B) {
	list := &LinkedList{}
	for i := 0; i < b.N; i++ {
		list.Append(&Value{})
	}
}

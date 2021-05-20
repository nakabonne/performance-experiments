package main

import "testing"

func BenchmarkSliceAppend(b *testing.B) {
	s := &Slice{s: []*T{}}
	for i := 0; i < b.N; i++ {
		s.Append(&T{})
	}
}

func BenchmarkLinkedListAppend(b *testing.B) {
	list := &LikedList{}
	for i := 0; i < b.N; i++ {
		list.Append(&T{})
	}
}

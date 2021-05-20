package slice_linkedlist

import "sync"

type T struct{}

type Slice struct {
	s  []*T
	mu sync.Mutex
}

func (s *Slice) Append(t *T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.s = append(s.s, t)
}

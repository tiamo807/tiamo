package set

import "sync"

type Set[T comparable] struct {
	mu  sync.RWMutex
	val map[T]struct{}
}

func New[T comparable](item ...T) *Set[T] {
	s := &Set[T]{
		val: make(map[T]struct{}),
	}

	s.Put(item...)

	return s
}

func (s *Set[T]) Add(item T) *Set[T] {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.val[item] = struct{}{}
	return s
}

func (s *Set[T]) Put(item ...T) *Set[T] {
	for _, el := range item {
		s.Add(el)
	}
	return s
}

func (s *Set[T]) Remove(item T) *Set[T] {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.val, item)
	return s
}

func (s *Set[T]) Contains(item T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, ok := s.val[item]
	return ok
}

func (s *Set[T]) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.val)
}

func (s *Set[T]) Clear() *Set[T] {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.val = make(map[T]struct{})
	return s
}

func (s *Set[T]) List() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	els := make([]T, 0, s.Len())
	for item := range s.val {
		els = append(els, item)
	}
	return els
}

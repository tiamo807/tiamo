package stack

import "sync"

type Stack[T any] struct {
	mu sync.RWMutex

	val []T
}

func New[T any]() *Stack[T] {
	s := &Stack[T]{
		val: make([]T, 0),
	}
	return s
}

func (s *Stack[T]) Push(val T) *Stack[T] {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.val = append(s.val, val)
	return s
}

func (s *Stack[T]) Pop() (T, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.val) == 0 {
		var zero T
		return zero, false
	}
	item := s.val[len(s.val)-1]
	s.val = s.val[:len(s.val)-1]
	return item, true
}

func (s *Stack[T]) Peek() (T, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if len(s.val) == 0 {
		var zero T
		return zero, false
	}
	return s.val[len(s.val)-1], true
}

func (s *Stack[T]) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.val)
}

func (s *Stack[T]) IsEmpty() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.val) == 0
}

func (s *Stack[T]) Clear() *Stack[T] {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.val = make([]T, 0)
	return s
}

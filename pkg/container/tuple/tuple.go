package tuple

import "reflect"

func equals(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

type Tuple[T any] struct {
	val []T
}

// New
// tp := New[string]("Canada","Mexico","Brazil")
// tp := New[any]("Tom", 18, "China")
func New[T any](ele ...T) *Tuple[T] {
	t := &Tuple[T]{
		val: make([]T, 0),
	}

	for _, el := range ele {
		t.val = append(t.val, el)
	}

	return t
}

func (t *Tuple[T]) Len() int {
	return len(t.val)
}

func (t *Tuple[T]) List() []T {
	return t.val
}

func (t *Tuple[T]) In(ele T) bool {
	for _, el := range t.val {
		if equals(ele, el) {
			return true
		}
	}
	return false
}

func (t *Tuple[T]) NotIn(ele T) bool {
	for _, el := range t.val {
		if equals(ele, el) {
			return false
		}
	}
	return true
}

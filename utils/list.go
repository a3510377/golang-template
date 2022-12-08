package utils

import (
	"sort"
)

type List[T comparable] []T

func New[T comparable](list ...T) List[T] {
	return List[T](list)
}

func (l List[T]) Len() int          { return len(l) }
func (l List[T]) Contains(e T) bool { return l.Index(e) >= 0 }

func (l List[T]) Reverse() List[T] {
	s := l.Clone()
	sort.SliceStable(s, func(i, j int) bool { return true })
	return s
}

func (l List[T]) Get(i int) *T {
	if i < len(l) {
		return &l[i]
	}
	return nil
}

func (l List[T]) Index(e T) int {
	for index, v := range l {
		if e == v {
			return index
		}
	}

	return -1
}

func (l List[T]) Clone() List[T] {
	return append(List[T]{}, l...)
}

// func (l List[T]) Pop() T {
// 	return
// }

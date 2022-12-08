package utils

import (
	"sort"
)

type List[T comparable] []T

func New[T comparable](list ...T) List[T] {
	return List[T](list)
}

func (l List[T]) Count() int                   { return len(l) }
func (l List[T]) Contains(e T) bool            { return l.Index(e) >= 0 }
func (l List[T]) Pop() T                       { return l.Remove(l.Count() - 1) }
func (l List[T]) Clone() List[T]               { return append(List[T]{}, l...) }
func (l List[T]) Slice(start, end int) List[T] { return l[start:end] }
func (l *List[T]) Append(i int, v ...T)        { *l = append(*l, v...) }
func (l *List[T]) Clear()                      { *l = []T{} }

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

func (l *List[T]) Insert(i int, v ...T) {
	if tot := len(*l) + len(v); tot <= cap(*l) {
		s2 := l.Slice(0, tot)
		copy(s2[i+len(v):], (*l)[i:])
		copy(s2[i:], v)
		*l = s2
	} else {
		s2 := make([]T, tot)
		copy(s2, l.Slice(0, i))
		copy(s2[i:], v)
		copy(s2[i+len(v):], (*l)[i:])
		*l = s2
	}
}

func (l *List[T]) Remove(s int) T {
	key := (*l)[s]
	*l = append(l.Slice(0, s), l.Slice(s+1, l.Count())...)

	return key
}

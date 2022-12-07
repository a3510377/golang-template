package utils

import (
	"sort"
)

type List[T comparable] []T

func (l List[T]) Len() int          { return len(l) }
func (l List[T]) FindFirst(e T) int { return FindFirst(l, e) }
func (l List[T]) FindLast(e T) int  { return FindLast(l, e) }
func (l List[T]) Contains(e T) bool { return Contains(l, e) }
func (l List[T]) Reverse() List[T]  { return List[T](Reverse(l)) }
func (l List[T]) Index(i int) T     { return l[i] }

func FindFirst[T comparable](s []T, e T) int {
	for index, v := range s {
		if e == v {
			return index
		}
	}

	return -1
}

func FindLast[T comparable](s []T, e T) int {
	for index, v := range Reverse(s) {
		if e == v {
			return index
		}
	}

	return -1
}

func Contains[T comparable](s []T, e T) bool {
	return FindFirst(s, e) >= 0
}

func Reverse[T comparable](l []T) []T {
	s := l
	sort.SliceStable(s, func(i, j int) bool { return true })
	return s
}

func RuneList(data string) []rune {
	return []rune(data)
}

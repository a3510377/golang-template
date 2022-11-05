package utils

func FindFirst[T comparable](s []T, e T) int {
	for index, v := range s {
		if e == v {
			return index
		}
	}

	return -1
}

func Contains[T comparable](s []T, e T) bool {
	return FindFirst(s, e) >= 0
}

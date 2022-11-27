package main

import (
	"fmt"
	"math"
)

const (
	escape     = "%"
	leftDelim  = "{{"
	rightDelim = "}}"
)

type String string

func (s String) String() string {
	return string(s)
}

func (s String) StringList() []int32 {
	return []rune(s)
}

func (s String) Index(index int) String {
	st := s.StringList()
	index = int(math.Abs(float64(index)))

	if len(st) > index {
		return String(st[index])
	}

	return ""
}

func (s String) Slice(start, end int) String {
	result := ""

	for _, v := range s.StringList()[start:end] {
		result += string(v)
	}

	return String(result)
}

func FormatString(sr string, args map[string]string) string {
	result := ""
	str := String(sr)
	isFormat := false
	// formatData := ""

	for i := 0; i < len(str.StringList()); i++ {
		s := str.Index(i)

		if s == escape {
			i += len(escape) + len(leftDelim)
			str.Slice(i, len(str))
			continue
		}
		if str.Slice(i, i+len(leftDelim)) == leftDelim {
			i += len(leftDelim) - 1
			isFormat = true
			continue
		}
		if str.Slice(i, i+len(rightDelim)) == rightDelim {
			i += len(rightDelim) - 1
			isFormat = false
			continue
		}
		if isFormat {
			fmt.Println("formatData", s)
		}
	}

	return result
}

// func main() {
// 	FormatString("你好 {{ awa }}")
// }

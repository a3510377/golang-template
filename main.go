package main

import (
	"test/log"
)

func main() {
	t := log.Template{
		StartDelimiter: "{",
		EndDelimiter:   "}",
		BreakDelimiter: "\\",
	}
	t.GetMatch("{{test")
}

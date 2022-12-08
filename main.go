package main

import (
	"fmt"

	"test/log"
	"test/utils"
)

func main() {
	test := utils.New("0", "1")
	fmt.Println(test)

	fmt.Println(test.Reverse())
	fmt.Println(test)
	t := log.Template{
		StartDelimiter: "{",
		EndDelimiter:   "}",
		BreakDelimiter: "\\",
	}
	t.GetMatch("{{test")
}

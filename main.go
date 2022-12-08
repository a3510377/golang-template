package main

import (
	"fmt"

	"test/log"
	"test/utils"
)

func main() {
	test := utils.New("0", "1", "2", "3", "4", "5", "6", "7", "8", "9")
	fmt.Println(test)

	test.Insert(0, "awa", "awa2")
	fmt.Println(test)
	t := log.Template{
		StartDelimiter: "{",
		EndDelimiter:   "}",
		BreakDelimiter: "\\",
	}
	t.GetMatch("{{test")
}

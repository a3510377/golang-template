package main

import (
	"fmt"

	"test/i18n"
	"test/log"
)

func main() {
	fmt.Println(i18n.Init("en"))

	i18n.Tr("test")
	i18n.Tr("test")
	i18n.Tr("test\ntest")

	// log.Test()

	a := log.WARN
	a |= log.INFO

	// ar.level &= ^level
	a &= log.INFO

	fmt.Println((a & log.DEBUG) == log.DEBUG)
	fmt.Printf("%#05b\n", a)
	fmt.Printf("%#05b\n", log.INFO)
}

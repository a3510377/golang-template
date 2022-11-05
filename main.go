package main

import (
	"fmt"

	"test/i18n"
)

func main() {
	fmt.Println(i18n.Init("en"))

	i18n.Tr("test")
	i18n.Tr("test")
	i18n.Tr("test")
}

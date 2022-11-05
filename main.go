package main

import (
	"fmt"

	"test/i18n"
	"test/i18n/po"
)

func main() {
	fmt.Println(i18n.Init("en"))
	po.GeneratePo()
	i18n.Tr("test")
}

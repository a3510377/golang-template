package main

import (
	"fmt"
	"os"

	"test/i18n"
	"test/i18n/po"
)

func main() {
	fmt.Println(i18n.Init("en"))
	po := po.GeneratePo()

	f, _ := os.Create("test.po")
	defer f.Close()
	po.Write(f)

	i18n.Tr("test")
}

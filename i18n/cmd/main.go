package main

import (
	"fmt"
	"os"

	"test/i18n/cmd/commands"
)

func main() {
	if err := commands.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

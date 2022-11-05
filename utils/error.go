package utils

import "fmt"

func Recover() {
	if r := recover(); r != nil {
		// TODO use logging
		fmt.Println(r)
	}
}

package po

import (
	"fmt"
	"os"
	"path/filepath"
)

type Message struct {
	FileLine string
	Value    string
}

func GeneratePo() {
	var pos []Message

	filepath.Walk(".", func(name string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || filepath.Ext(name) != ".go" {
			return nil
		}

		if msg := parse(name); msg != nil {
			pos = append(pos, *msg)
		}

		return nil
	})

	fmt.Println(pos)
}

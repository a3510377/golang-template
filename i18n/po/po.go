package po

import (
	"os"
	"path/filepath"
)

type (
	Message struct {
		FileLine []string
		Value    string
		Id       string
	}
	PoFile struct {
		Messages map[string]*Message
		Lang     string
	}
)

func GeneratePo() *PoFile {
	poFile := &PoFile{}

	filepath.Walk(".", func(name string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || filepath.Ext(name) != ".go" {
			return nil
		}

		if msg := parse(name); msg != nil {
			poFile.AddMessage(msg...)
		}

		return nil
	})

	return poFile
}

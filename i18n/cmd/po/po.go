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

func GeneratePo(path string, lang string) *PoFile {
	poFile := &PoFile{}
	poFile.Lang = lang

	filepath.Walk(path, func(name string, info os.FileInfo, err error) error {
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

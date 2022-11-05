package po

import (
	"fmt"
	"os"
	"strings"

	"test/utils"
)

func (po *PoFile) Add(id, value string, fileLine []string) {
	if po.Messages == nil {
		po.Messages = map[string]*Message{}
	}

	if po.Messages[id] == nil {
		po.Messages[id] = &Message{Id: id, Value: value, FileLine: fileLine}
	} else {
		for _, m := range fileLine {
			if !utils.Contains(po.Messages[id].FileLine, m) {
				po.Messages[id].FileLine = append(po.Messages[id].FileLine, m)
			}
		}
	}
}

func (po *PoFile) AddMessage(msg ...Message) {
	for _, m := range msg {
		po.Add(m.Id, m.Value, m.FileLine)
	}
}

func (po *PoFile) Write(w *os.File) {
	contents := []string{
		fmt.Sprintf(`"Language: %s"`, po.Lang),
	}

	for _, msg := range po.Messages {
		contents = append(contents, "\n")

		for _, file := range msg.FileLine {
			contents = append(contents, fmt.Sprintf("#: %s\n", file))
		}

		contents = append(
			contents,
			format(`msgid "%s"`, msg.Id),
			format(`msgstr "%s"`, msg.Value),
		)
	}

	for _, content := range contents {
		w.Write([]byte(content))
	}
}

func format(value string, arg string) string {
	return (fmt.Sprintf(value+"\n", escape(arg)))
}

func escape(value string) string {
	return strings.ReplaceAll(strings.ReplaceAll(value, `"`, `\"`), "\n", "\\n")
}

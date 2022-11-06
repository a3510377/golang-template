package log

type LogRecord struct {
	args      []any
	msg       string
	levelName string
}

type (
	Style        interface{}
	PercentStyle struct {
		Style
	}
	StrFormatStyle struct {
		Style
	}
	StringTemplateStyle struct {
		Style
	}
)

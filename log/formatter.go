package log

import (
	"text/template"
)

type Formatter struct {
	default_time_format string
	default_msec_format string
}

func NewFormatter() *Formatter {
	f := &Formatter{}

	f.default_msec_format = "%Y-%m-%d %H:%M:%S"
	f.default_time_format = "%s,%03d"

	// regexp.Compile("")
	t1 := template.New("")
	t1.Parse("")

	return f
}

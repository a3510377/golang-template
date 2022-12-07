package log

import "fmt"

type StyleObject interface {
	Format(record LogRecord) string
}

type Style struct {
	StyleObject
	formatString string
}

func (s *Style) SetFormat(format string) {
	s.formatString = format
}

func (s *Style) GetFormat() string {
	return s.formatString
}

func (s *Style) Format(record LogRecord) string {
	return record.String()
}

/* ---------------------------------------------------------------- */
type PercentStyle struct{ Style }

func (s *PercentStyle) Format(record LogRecord) string {
	return ""
}

func NewPercentStyle(fmt string) *PercentStyle {
	style := &PercentStyle{}
	style.SetFormat(fmt)
	return style
}

/* ---------------------------------------------------------------- */
type StrFormatStyle struct{ Style }

func (s *StrFormatStyle) Format(record LogRecord) string {
	return ""
}

func NewStrFormatStyle(fmt string) *StrFormatStyle {
	style := &StrFormatStyle{}
	style.SetFormat(fmt)
	return style
}

/* ---------------------------------------------------------------- */

type StringTemplateStyle struct{ Style }

func (s *StringTemplateStyle) Format(record LogRecord) string {
	return ""
}

func NewStringTemplateStyle(fmt string) *StringTemplateStyle {
	style := &StringTemplateStyle{}
	style.SetFormat(fmt)
	return style
}

/* ---------------------------------------------------------------- */
type Template struct {
	StartDelimiter string
	EndDelimiter   string
	BreakDelimiter string
}

func (t *Template) GetMatch(data string) map[string]string {
	str := []rune(data)
	result := map[string]string{}
	for i, ch := range str {
		switch string(ch) {
		case t.StartDelimiter:
		case t.EndDelimiter:
		case t.BreakDelimiter:
		}
		if t.StartDelimiter == string(ch) {
			fmt.Println(i, string(ch))
		}
	}
	return result
}

func (t *Template) String() string {
	return ""
}

package log

type StyleObject interface {
	Format(record LogRecord) string
}

/* ---------------------------------------------------------------- */
type PercentStyle struct {
	fmt string
	StyleObject
}

func (s *PercentStyle) Format(record LogRecord) string {
	return ""
}

func NewPercentStyle(fmt string) *PercentStyle {
	return &PercentStyle{fmt: fmt}
}

/* ---------------------------------------------------------------- */
type StrFormatStyle struct {
	fmt string
	StyleObject
}

func (s *StrFormatStyle) Format(record LogRecord) string {
	return ""
}

func NewStrFormatStyle(fmt string) *StrFormatStyle {
	return &StrFormatStyle{fmt: fmt}
}

/* ---------------------------------------------------------------- */

type StringTemplateStyle struct {
	fmt string
	StyleObject
}

func (s *StringTemplateStyle) Format(record LogRecord) string {
	return ""
}

func NewStringTemplateStyle(fmt string) *StringTemplateStyle {
	return &StringTemplateStyle{fmt: fmt}
}

/* ---------------------------------------------------------------- */
type Template struct{}

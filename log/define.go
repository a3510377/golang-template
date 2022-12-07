package log

const (
	NOTSET = 1 << iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL

	WARNING  = WARN
	CRITICAL = FATAL
)

var LevelToName = map[int]string{
	CRITICAL: "CRITICAL",
	ERROR:    "ERROR",
	WARN:     "WARN",
	INFO:     "INFO",
	DEBUG:    "DEBUG",
	NOTSET:   "NOTSET",
}

var NameToLevel = map[string]int{
	"CRITICAL": CRITICAL,
	"FATAL":    FATAL,
	"ERROR":    ERROR,
	"WARN":     WARNING,
	"WARNING":  WARNING,
	"INFO":     INFO,
	"DEBUG":    DEBUG,
	"NOTSET":   NOTSET,
}

var Styles = map[string]struct {
	style any
	fmt   string
}{
	"%": {PercentStyle{}, "%(levelName)s:%(name)s:%(message)s"},
	"{": {StrFormatStyle{}, "{levelName}:{name}:{message}"},
	"$": {StringTemplateStyle{}, "${levelName}:${name}:${message}"},
}

func GetLevelName(level int) string {
	if value, ok := LevelToName[level]; ok {
		return value
	}

	return "UNKNOWN"
}

func GetNameLevel(level string) int {
	if value, ok := NameToLevel[level]; ok {
		return value
	}

	return -1
}

type BaseRecord struct {
	level     int
	disabled  bool
	formatter string
}

func (r *BaseRecord) IsEnabledFor(level int) bool {
	if r.disabled {
		return false
	}

	return (r.level | level) == r.level
}

func (r *BaseRecord) GetLevel() int {
	return r.level
}

func (r *BaseRecord) SetLevel(level int) int {
	r.level = level

	return r.GetLevel()
}

func (r *BaseRecord) AddLevel(level int) int {
	r.level |= level

	return r.GetLevel()
}

func (r *BaseRecord) RemoveLevel(level int) int {
	r.level &^= level

	return r.GetLevel()
}

func (r *BaseRecord) SetDisabled(disabled bool) {
	r.disabled = disabled
}

func (r *BaseRecord) GetFormatter() string {
	return r.formatter
}

func (r *BaseRecord) SetFormatter(fmt string) {
	r.formatter = fmt
}

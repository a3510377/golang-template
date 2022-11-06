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

var LevelToName map[int]string = map[int]string{
	CRITICAL: "CRITICAL",
	ERROR:    "ERROR",
	WARN:     "WARN",
	INFO:     "INFO",
	DEBUG:    "DEBUG",
	NOTSET:   "NOTSET",
}

var NameToLevel map[string]int = map[string]int{
	"CRITICAL": CRITICAL,
	"FATAL":    FATAL,
	"ERROR":    ERROR,
	"WARN":     WARNING,
	"WARNING":  WARNING,
	"INFO":     INFO,
	"DEBUG":    DEBUG,
	"NOTSET":   NOTSET,
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

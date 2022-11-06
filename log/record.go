package log

type LogRecord struct {
	disabled bool
	level    int
}

func (r *LogRecord) Debug(msg string, args ...any) {
	r.log(DEBUG, msg, args...)
}

func (r *LogRecord) Info(msg string, args ...any) {
	r.log(INFO, msg, args...)
}

func (r *LogRecord) Warn(msg string, args ...any) {
	r.log(WARN, msg, args...)
}

func (r *LogRecord) Warning(msg string, args ...any) {
	r.log(WARNING, msg, args...)
}

func (r *LogRecord) Error(msg string, args ...any) {
	r.log(ERROR, msg, args...)
}

func (r *LogRecord) Critical(msg string, args ...any) {
	r.log(CRITICAL, msg, args...)
}

func (r *LogRecord) Fatal(msg string, args ...any) {
	r.log(FATAL, msg, args...)
}

func (r *LogRecord) Log(level int, msg string, args ...any) {
	if r.IsEnabledFor(level) {
		r.log(level, msg, args...)
	}
}

func (r *LogRecord) log(level int, msg string, args ...any) {
}

func (r *LogRecord) IsEnabledFor(level int) bool {
	if r.disabled {
		return false
	}

	return (r.level | level) == r.level
}

func (r *LogRecord) GetLevel() int {
	return r.level
}

func (r *LogRecord) SetLevel(level int) int {
	r.level = level

	return r.GetLevel()
}

func (r *LogRecord) AddLevel(level int) int {
	r.level |= level

	return r.GetLevel()
}

func (r *LogRecord) RemoveLevel(level int) int {
	r.level &^= level

	return r.GetLevel()
}

func (r *LogRecord) SetDisabled(disabled bool) {
	r.disabled = disabled
}

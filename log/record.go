package log

type Logger struct {
	BaseRecord
	handlers []Handler
}

func (r *Logger) Debug(msg string, args ...any) {
	r.log(DEBUG, msg, args...)
}

func (r *Logger) Info(msg string, args ...any) {
	r.log(INFO, msg, args...)
}

func (r *Logger) Warn(msg string, args ...any) {
	r.log(WARN, msg, args...)
}

func (r *Logger) Warning(msg string, args ...any) {
	r.log(WARNING, msg, args...)
}

func (r *Logger) Error(msg string, args ...any) {
	r.log(ERROR, msg, args...)
}

func (r *Logger) Critical(msg string, args ...any) {
	r.log(CRITICAL, msg, args...)
}

func (r *Logger) Fatal(msg string, args ...any) {
	r.log(FATAL, msg, args...)
}

func (r *Logger) Log(level int, msg string, args ...any) {
	if r.IsEnabledFor(level) {
		r.log(level, msg, args...)
	}
}

func (r *Logger) log(level int, msg string, args ...any) {
}

package logger

type Logger interface {
	LogLevel(level Level) error
	Trace(msg string, vars ...any)
	Debug(msg string, vars ...any)
	Info(msg string, vars ...any)
	Warn(msg string, vars ...any)
	Error(err error, msg string, vars ...any)
	Fatal(err error, msg string, vars ...any)
}

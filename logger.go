package logger

// Logger is an abstraction for a logger.
type Logger interface {

	// SetLevel sets log level.
	SetLevel(l Level) error

	// GetLevel returns a log level.
	GetLevel() Level

	// Trace writes log with trace level.
	Trace(msg string, vars ...any)

	// Debug writes log with debug level.
	Debug(msg string, vars ...any)

	// Info writes log with info level.
	Info(msg string, vars ...any)

	// Warn writes log with warn level.
	Warn(msg string, vars ...any)

	// Error writes log with error level.
	Error(err error, msg string, vars ...any)

	// Fatal writes log with fatal level.
	Fatal(err error, msg string, vars ...any)
}

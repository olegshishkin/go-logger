package logger1

import (
	"log"
	"os"
	"path/filepath"
)

type Logger interface {
	LogLevel(level Level) error
	Trace(msg string, vars ...any)
	Debug(msg string, vars ...any)
	Info(msg string, vars ...any)
	Warn(msg string, vars ...any)
	Error(err error, msg string, vars ...any)
	Fatal(err error, msg string, vars ...any)
}

func LogFile() *os.File {
	dir := os.Getenv("LOG_DIR")
	if dir == "" {
		dir = "/var/log/app"
	}

	fileName := os.Getenv("LOG_FILE_NAME")
	if fileName == "" {
		fileName = "app.log"
	}

	err := os.MkdirAll(dir, 0664)
	if err != nil {
		panic(err)
	}

	path := filepath.Join(dir, fileName)
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		panic(err)
	}

	log.Printf("Use log file: %v", path)
	return file
}

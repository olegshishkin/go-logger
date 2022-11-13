package logger

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	Trace Level = iota
	Debug
	Info
	Warn
	Error
	Fatal
)

var (
	DefaultLogLevel = Info
)

type Level int8

func LogLevel() Level {
	envLevel := os.Getenv("LOG_LEVEL")
	if envLevel == "" {
		log.Printf("Set default log level: %v", DefaultLogLevel)
		return DefaultLogLevel
	}

	level, err := ParseLevel(strings.ToLower(envLevel))
	if err != nil {
		panic(err)
	}

	log.Printf("Set log level: %v", level)
	return level
}

func ParseLevel(level string) (Level, error) {
	switch level {
	case Trace.String():
		return Trace, nil
	case Debug.String():
		return Debug, nil
	case Info.String():
		return Info, nil
	case Warn.String():
		return Warn, nil
	case Error.String():
		return Error, nil
	case Fatal.String():
		return Fatal, nil
	default:
		return -1, errors.New(fmt.Sprintf("Unknown log level: %v", level))
	}
}

func (l Level) String() string {
	switch l {
	case Trace:
		return "trace"
	case Debug:
		return "debug"
	case Info:
		return "info"
	case Warn:
		return "warn"
	case Error:
		return "error"
	case Fatal:
		return "fatal"
	default:
		panic(fmt.Sprintf("Unknown LogLevel %v", strconv.Itoa(int(l))))
	}
}

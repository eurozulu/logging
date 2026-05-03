package logging

import (
	"fmt"
	"io"
	"os"
	"time"
)

const timeformat = "2006-01-02 15:04:05.111"

// DefaultLogger is the logger used for the package.
var DefaultLogger Logger = &logger{
	level: LogError,
	out:   os.Stdout,
}

type Logger interface {
	Log(level LogLevel, msg string, a ...any)
	LogLevel() LogLevel
	SetLogLevel(level LogLevel)
	SetOutput(out io.Writer)
}

type logger struct {
	level LogLevel
	out   io.Writer
}

func (l logger) Log(level LogLevel, msg string, a ...any) {
	if level > l.level {
		return
	}
	if level > LogInfo {
		fmt.Fprintf(l.out, "[%s] [%s] %s\n", time.Now().Format(timeformat), level, fmt.Sprintf(msg, a...))
	} else {
		fmt.Fprintf(l.out, "%s\n", fmt.Sprintf(msg, a...))
	}
}

func (l logger) LogLevel() LogLevel {
	return l.level
}

func (l *logger) SetLogLevel(level LogLevel) {
	if level >= LogDebug {
		l.Log(l.level, "Changing log level from %s to %s", l.level, level)
	}
	l.level = level
}

func (l *logger) SetOutput(out io.Writer) {
	l.out = out
}

func Error(msg string, a ...any) {
	DefaultLogger.Log(LogError, msg, a...)
}

func Warning(msg string, a ...any) {
	DefaultLogger.Log(LogWarning, msg, a...)
}

func Info(msg string, a ...any) {
	DefaultLogger.Log(LogInfo, msg, a...)
}

func Debug(msg string, a ...any) {
	DefaultLogger.Log(LogDebug, msg, a...)
}

func Trace(msg string, a ...any) {
	DefaultLogger.Log(LogTrace, msg, a...)
}

func IsLogLevel(level LogLevel) bool {
	return DefaultLogger.LogLevel() >= level
}

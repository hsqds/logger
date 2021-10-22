package logger

import (
	"log"
	"os"
)

var (
	_ Logger = &StdLogger{}
)

// StdLogger represents
type StdLogger struct {
	level Level
}

func NewStdLogger(lvl Level) *StdLogger {
	return &StdLogger{
		level: lvl,
	}
}

// SetLevel
func (l *StdLogger) SetLevel(lvl Level) {
	l.level = lvl
}

// Log
func (l *StdLogger) log(lvl Level, msg string, args ...interface{}) {
	if lvl > l.level {
		return
	}

	log.Printf(msg, args...)
}

// Debug
func (l *StdLogger) Debug(msg string, args ...interface{}) {
	l.log(LevelDebug, msg, args...)
}

// Info
func (l *StdLogger) Info(msg string, args ...interface{}) {
	l.log(LevelInfo, msg, args...)
}

// Warn
func (l *StdLogger) Warn(msg string, args ...interface{}) {
	l.log(LevelWarn, msg, args...)
}

// Error
func (l *StdLogger) Error(msg string, args ...interface{}) {
	l.log(LevelError, msg, args...)
}

// Fatal
func (l *StdLogger) Fatal(msg string, args ...interface{}) {
	l.log(LevelCritical, msg, args...)
	os.Exit(1)
}

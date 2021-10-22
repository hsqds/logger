package logger

type Level int

const (
	LevelCritical Level = iota
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
)

// Logger
type Logger interface {
	SetLevel(lvl Level)

	Fatal(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Debug(msg string, args ...interface{})
}

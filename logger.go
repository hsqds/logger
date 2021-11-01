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

	// Fatal
	F(msg string, args ...interface{})
	// Error
	E(msg string, args ...interface{})
	// Warning
	W(msg string, args ...interface{})
	// Info
	I(msg string, args ...interface{})
	// Debug
	D(msg string, args ...interface{})
}

package logging

import (
	"fmt"
	"sails-sftp/backend/utils/date"
)

type LogLevel int

const (
	LevelDebug    LogLevel = iota
	LevelInfo     LogLevel = iota
	LevelWarn     LogLevel = iota
	LevelError    LogLevel = iota
	LevelCritical LogLevel = iota
)

type Logger struct {
	level            LogLevel
	outputPath       string
	isEnabled        bool
	includeTimestamp bool
}

func NewLogger(level LogLevel, outputPath string) *Logger {
	return &Logger{
		level:            validateLogLevel(level),
		outputPath:       outputPath,
		isEnabled:        true, // Default to enabled
		includeTimestamp: true, // Default to including timestamps
	}
}

func (l *Logger) SetLevel(level LogLevel) {
	l.level = validateLogLevel(level)
}

func (l *Logger) Disable() {
	l.isEnabled = false
}

func (l *Logger) Enable() {
	l.isEnabled = true
}

func (l *Logger) Log(level LogLevel, message string) {
	if !l.isEnabled {
		return
	}

	if level < l.level {
		return
	}

	var formattedMessage string

	switch level {
	case LevelDebug:
		formattedMessage = "[DEBUG] "
	case LevelInfo:
		formattedMessage = "[INFO] "
	case LevelWarn:
		formattedMessage = "[WARN] "
	case LevelError:
		formattedMessage = "[ERROR] "
	case LevelCritical:
		formattedMessage = "[CRITICAL] "
	default:
		formattedMessage = "[UNKNOWN] "
	}

	if l.includeTimestamp {
		formattedMessage += fmt.Sprintf("(%s) %s", date.GetCurrentTimestamp(), message)
	} else {
		formattedMessage += message
	}

	// TODO: File output logic here, for now we are just printing to stdout
	fmt.Println(formattedMessage)
}

func validateLogLevel(level LogLevel) LogLevel {
	if level < LevelDebug || level > LevelCritical {
		fmt.Println("Invalid log level. Setting to default (LevelInfo).")
		return LevelInfo
	}

	return level
}

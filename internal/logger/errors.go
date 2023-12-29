// Package logger provides an interface and implementation for logging functionality.
package logger

import (
	"fmt"
	"time"
)

const (
	LevelDebug = "DEBUG"
	LevelInfo  = "INFO"
	LevelWarn  = "WARN"
	LevelError = "ERROR"
	LevelFatal = "FATAL"
)

// Logger defines methods that can be used for various logging operations.
type Logger interface {
	String(key string) (string, error)
	Bool(key string) (bool, error)
	Int(key string) (int, error)
	Float(key string) (float64, error)
	Bytes(key string) ([]byte, error)
	JSON(key string) (interface{}, error)
	JSONLines(key string) ([]interface{}, error)
	Table(header map[string]string, rows []map[string]string) error
	LogError(err error) error
	LogInvalid(err error)
	LogInfo(message string)
}

// LoggerImpl is the concrete implementation of the Logger interface.
type LoggerImpl struct{}

// NewLogger returns a new instance of LoggerImpl.
func NewLogger() *LoggerImpl {
	return &LoggerImpl{}
}

// String converts the value associated with the key to a string.
func (l *LoggerImpl) String(key string) (string, error) {
	return "", nil
}

// Bool converts the value associated with the key to a boolean.
func (l *LoggerImpl) Bool(key string) (bool, error) {
	panic("implement me")
}

// Int converts the value associated with the key to an integer.
func (l *LoggerImpl) Int(key string) (int, error) {
	panic("implement me")
}

// Float converts the value associated with the key to a float64.
func (l *LoggerImpl) Float(key string) (float64, error) {
	panic("implement me")
}

// Bytes converts the value associated with the key to a byte slice.
func (l *LoggerImpl) Bytes(key string) ([]byte, error) {
	panic("implement me")
}

// JSON parses the value associated with the key as a JSON object.
func (l *LoggerImpl) JSON(key string) (interface{}, error) {
	panic("implement me")
}

// JSONLines parses the value associated with the key as a list of JSON objects.
func (l *LoggerImpl) JSONLines(key string) ([]interface{}, error) {
	panic("implement me")
}

// Table logs tabular data using a header and rows.
func (l *LoggerImpl) Table(header map[string]string, rows []map[string]string) error {
	panic("implement me")
}

// LogError logs the given error.
func (l *LoggerImpl) LogError(err error) error {
	panic("implement me")
}

// LogInvalid logs an invalid error type.
func (l *LoggerImpl) LogInvalid(err error) {
	panic("implement me")
}

func (l *LoggerImpl) LogInfo(message string) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	// Format the log message
	formattedMessage := fmt.Sprintf("[%s] [%s] [%s]", currentTime, LevelInfo, message)
	fmt.Println(formattedMessage)

}

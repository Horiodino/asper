// Package logger provides an interface and implementation for logging functionality.
package logger

import (
	"fmt"
	"reflect"
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
	Bytes(msg string, option interface{}) error
	JSON(msg string, option interface{}) error
	JSONLines(msg string, option interface{}) error
	Table(header map[string]string, rows []map[string]string) error
	LogError(msg string, option interface{}) error
	LogInvalid(msg string, option interface{}) error
	LogInfo(msg string, option interface{}) error
	LogDebug(msg string, option interface{}) error
	LogWarn(msg string, option interface{}) error
	LogFatal(msg string, option interface{}) error
}

// LoggerImpl is the concrete implementation of the Logger interface.
type LoggerImpl struct {
	now     time.Time
	warning bool
	// the warning is a bool based on that we set the color of logger like there is an warning previusly so after that if we use debugger
	// then the print will see is warning set to true , if yes then we will print the warning color
}

// NewLogger returns a new instance of LoggerImpl.
func NewLogger() *LoggerImpl {
	return &LoggerImpl{}
}

func (l *LoggerImpl) LogError(msg string, option interface{}) error {

	return nil
}

func (l *LoggerImpl) LogInvalid(msg string, option interface{}) error {
	fmt.Println(msg)
	return nil
}

func (l *LoggerImpl) LogInfo(msg string, option interface{}) error {
	fmt.Print(l.now.Format("2006-01-02 15:04:05"), " ", LevelInfo, " ", msg)
	return nil
}

func (l *LoggerImpl) LogDebug(msg string, option interface{}) error {
	fmt.Println(l.now.Format("2006-01-02 15:04:05"), " ", LevelDebug, " ", msg)
	if reflect.TypeOf(option).Kind() == reflect.Struct {
		val := reflect.ValueOf(option)
		for i := 0; i < val.NumField(); i++ {
			fmt.Println(val.Type().Field(i).Name, ":", val.Field(i).Interface())
		}
	}

	return nil
}

func (l *LoggerImpl) LogWarn(msg string, option interface{}) error {
	fmt.Println(msg)
	return nil
}

func (l *LoggerImpl) LogFatal(msg string, option interface{}) error {
	fmt.Println(msg)
	return nil
}

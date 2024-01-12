package logger

import (
	"fmt"
	"reflect"
	"time"

	"github.com/fatih/color"
)

const (
	LevelDebug = "DEBUG"
	LevelInfo  = "INFO"
	LevelWarn  = "WARN"
	LevelError = "ERROR"
	LevelFatal = "FATAL"
)

type Logger interface {
	// ... (rest remains the same)
}

type LoggerImpl struct {
	now     time.Time
	warning bool
}

func NewLogger() *LoggerImpl {
	return &LoggerImpl{}
}

func (l *LoggerImpl) LogError(msg string, option interface{}) error {
	color.Red(l.now.Format("2006-01-02 15:04:05") + " " + LevelError + " " + msg)
	return nil
}

func (l *LoggerImpl) LogInvalid(msg string, option interface{}) error {
	color.Yellow(msg)
	return nil
}

func (l *LoggerImpl) LogInfo(msg string, option interface{}) error {
	color.Blue(l.now.Format("2006-01-02 15:04:05") + " " + LevelInfo + " " + msg)
	return nil
}

func (l *LoggerImpl) LogDebug(msg string, option interface{}) error {
	color.Yellow(l.now.Format("2006-01-02 15:04:05") + " " + LevelDebug + " " + msg)
	if reflect.TypeOf(option).Kind() == reflect.Struct {
		val := reflect.ValueOf(option)
		for i := 0; i < val.NumField(); i++ {
			color.Blue(val.Type().Field(i).Name + ":" + fmt.Sprint(val.Field(i).Interface()))
		}
	}
	return nil
}

func (l *LoggerImpl) LogWarn(msg string, option interface{}) error {
	color.Yellow(msg)
	return nil
}

func (l *LoggerImpl) LogFatal(msg string, option interface{}) error {
	color.Red(msg)
	return nil
}

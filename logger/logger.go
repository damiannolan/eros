package logger

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	stdLogger *logrus.Logger
)

// Fields alias type for logging key value pairs
type Fields map[string]interface{}

// Entry wrapper type Entry of logrus.Entry
type Entry logrus.Entry

func init() {
	stdLogger = newLogger()
}

func newLogger() *logrus.Logger {
	return &logrus.Logger{
		Out:       os.Stdout,
		Formatter: newTextFormatter(),
		Level:     logrus.DebugLevel,
	}
}

func newJSONFormatter() *logrus.JSONFormatter {
	return &logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
	}
}

func newTextFormatter() *logrus.TextFormatter {
	return &logrus.TextFormatter{
		DisableColors:   true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339Nano,
	}
}

// Debug logs the provided args at the [DEBUG] level
func Debug(args ...interface{}) {
	stdLogger.Debug(args...)
}

// Error logs the args at the [ERROR] level
func Error(args ...interface{}) {
	stdLogger.Error(args...)
}

// Fatal logs the provided args at the [FATAL] level and calls os.Exit(1)
func Fatal(args ...interface{}) {
	stdLogger.Fatal(args...)
}

// Info logs the provided args at the [INFO] level
func Info(args ...interface{}) {
	stdLogger.Info(args...)
}

// Trace logs the provided args at the [TRACE] level
func Trace(args ...interface{}) {
	stdLogger.Trace(args...)
}

// Warn logs the provided message at the [WARN] level
func Warn(args ...interface{}) {
	stdLogger.Fatal(args...)
}

// WithFields returns an Entry for chaining with Info, Debug, Trace, etc...
// Allows for logging of key value pairs
// log.WithFields(log.Fields{"key": "value"}).Info("Logging key/value entry")
func WithFields(fields Fields) Entry {
	return Entry(*stdLogger.WithFields(logrus.Fields(fields)))
}

// WithError returns an Entry for chaining
func WithError(err error) Entry {
	return Entry(*stdLogger.WithError(err))
}

// Debug logs the Entry data fields at [DEBUG] level
func (e Entry) Debug(args ...interface{}) {
	stdLogger.WithFields(e.Data).Debug(args...)
}

// Error logs the Entry data fields at the [ERROR] level
func (e Entry) Error(args ...interface{}) {
	stdLogger.WithFields(e.Data).Error(args...)
}

// Info logs the Entry data fields at the [INFO] level
func (e Entry) Info(args ...interface{}) {
	stdLogger.WithFields(e.Data).Info(args...)
}

// Fatal logs the Entry data fields at the [FATAL] level
func (e Entry) Fatal(args ...interface{}) {
	stdLogger.WithFields(e.Data).Fatal(args...)
}

// Trace logs the Entry data fields at the [TRACE] level
func (e Entry) Trace(args ...interface{}) {
	stdLogger.WithFields(e.Data).Trace(args...)
}

// Warn logs the Entry data fields at the [WARN] level
func (e Entry) Warn(args ...interface{}) {
	stdLogger.WithFields(e.Data).Warn(args...)
}

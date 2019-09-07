package logger

import (
	"github.com/sirupsen/logrus"
)

type defaultLogger struct {
	log    *logrus.Logger
	fields map[string]interface{}
}

func NewDefaultLogger() *defaultLogger {
	dl := &defaultLogger{}
	dl.log = logrus.New()
	dl.log.SetFormatter(&logrus.TextFormatter{})
	return dl
}

func (dl *defaultLogger) SetLevel(level Level) {
	var lvl logrus.Level
	switch level {
	case Debug:
		lvl = logrus.DebugLevel
	case Info:
		lvl = logrus.InfoLevel
	case Error:
		lvl = logrus.ErrorLevel
	}
	dl.log.SetLevel(logrus.Level(lvl))
}

func (dl *defaultLogger) Debug(msg string) {
	dl.log.WithFields(logrus.Fields(dl.fields)).Debug(msg)
}

func (dl *defaultLogger) Info(msg string) {
	dl.log.WithFields(logrus.Fields(dl.fields)).Info(msg)
}
func (dl *defaultLogger) Error(msg string) {
	dl.log.WithFields(logrus.Fields(dl.fields)).Error(msg)
}

func (dl *defaultLogger) WithFields(fields map[string]interface{}) Logger {
	ndl := &defaultLogger{
		log:    dl.log,
		fields: make(map[string]interface{}),
	}
	for k, v := range fields {
		ndl.fields[k] = v
	}
	for k, v := range dl.fields {
		ndl.fields[k] = v
	}
	return ndl
}

type noopLogger struct{}

func NewNoopLogger() *noopLogger {
	return &noopLogger{}
}

func (*noopLogger) Debug(msg string) {}
func (*noopLogger) Info(msg string)  {}
func (*noopLogger) Error(msg string) {}
func (n *noopLogger) WithFields(fields map[string]interface{}) Logger {
	return n
}

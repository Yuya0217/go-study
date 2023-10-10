package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Fields logrus.Fields

//go:generate mockgen -source=$GOFILE -destination=./mocks/mock_$GOFILE
type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Fatal(args ...interface{})
	WithFields(fields Fields) Logger
}

type LogrusLogger struct {
	*logrus.Entry
}

func New() Logger {
	log := logrus.New()
	log.Out = os.Stdout
	log.Formatter = &logrus.JSONFormatter{}

	defaultFields := logrus.Fields{
		"serviceName": "go-layered-architecture-sample",
	}

	return &LogrusLogger{
		log.WithFields(defaultFields),
	}

}

func (l *LogrusLogger) WithFields(fields Fields) Logger {
	return &LogrusLogger{
		l.Entry.WithFields(logrus.Fields(fields)),
	}
}

func (l *LogrusLogger) Debug(args ...interface{}) {
	l.Entry.Debug(args...)
}

func (l *LogrusLogger) Info(args ...interface{}) {
	l.Entry.Info(args...)
}

func (l *LogrusLogger) Warn(args ...interface{}) {
	l.Entry.Warn(args...)
}

func (l *LogrusLogger) Error(args ...interface{}) {
	l.Entry.Error(args...)
}

func (l *LogrusLogger) Errorf(format string, args ...interface{}) {
	l.Entry.Errorf(format, args...)
}

func (l *LogrusLogger) Printf(format string, args ...interface{}) {
	l.Entry.Printf(format, args...)
}

func (l *LogrusLogger) Fatal(args ...interface{}) {
	l.Entry.Fatal(args...)
}

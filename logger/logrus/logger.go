package logrus

import (
	"github.com/sirupsen/logrus"
)

type Logger struct {
	logger *logrus.Logger
}

func NewLogger(logger *logrus.Logger) *Logger {
	return &Logger{logger: logger}
}

func (l *Logger) Info(args ...any) {
	l.logger.Info(args...)
}

func (l *Logger) Error(args ...any) {
	l.logger.Error(args...)
}

func (l *Logger) Fatal(args ...any) {
	l.logger.Fatal(args...)
}

package logrus

import (
	"github.com/sirupsen/logrus"
)

type Entry struct {
	logger *logrus.Entry
}

func NewEntry(logger *logrus.Entry) *Entry {
	return &Entry{logger: logger}
}

func (l *Entry) Info(args ...any) {
	l.logger.Info(args...)
}

func (l *Entry) Error(args ...any) {
	l.logger.Error(args...)
}

func (l *Entry) Fatal(args ...any) {
	l.logger.Fatal(args...)
}

package zap

import (
	"fmt"

	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.Logger
}

func NewLogger(logger *zap.Logger) *Logger {
	return &Logger{logger: logger}
}

func (l *Logger) Info(args ...any) {
	l.logger.Info(fmt.Sprint(args...))
}

func (l *Logger) Error(args ...any) {
	l.logger.Error(fmt.Sprint(args...))
}

func (l *Logger) Fatal(args ...any) {
	l.logger.Fatal(fmt.Sprint(args...))
}

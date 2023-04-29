package zap

import "go.uber.org/zap"

type Sugared struct {
	logger *zap.SugaredLogger
}

func NewSugared(logger *zap.SugaredLogger) *Sugared {
	return &Sugared{logger: logger}
}

func (s *Sugared) Warn(msg string, args ...any) {
	args = append([]any{msg}, args...)
	s.logger.Warn(args...)
}

func (s *Sugared) Error(msg string, args ...any) {
	args = append([]any{msg}, args...)
	s.logger.Error(args...)
}

func (s *Sugared) Fatal(msg string, args ...any) {
	args = append([]any{msg}, args...)
	s.logger.Fatal(args...)
}
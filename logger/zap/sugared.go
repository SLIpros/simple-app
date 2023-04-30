package zap

import "go.uber.org/zap"

type Sugared struct {
	logger *zap.SugaredLogger
}

func NewSugared(logger *zap.SugaredLogger) *Sugared {
	return &Sugared{logger: logger}
}

func (s *Sugared) Info(args ...any) {
	s.logger.Info(args...)
}

func (s *Sugared) Error(args ...any) {
	s.logger.Error(args...)
}

func (s *Sugared) Fatal(args ...any) {
	s.logger.Fatal(args...)
}

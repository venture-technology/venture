package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.Logger
}

func New(task string) (*Logger, error) {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return &Logger{
		logger: zapLogger,
	}, nil
}

func (l *Logger) Infof(format string, args ...zap.Field) {
	l.logger.Info(format, args...)
}

func (l *Logger) Errorf(format string, args ...zap.Field) {
	l.logger.Error(format, args...)
}

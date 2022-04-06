package logger

import "go.uber.org/zap"

type logger struct {
	*zap.Logger
}

func NewLogger() *logger {
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()

	return &logger{zapLogger}
}

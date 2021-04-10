package logger

import (
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func init() {
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()
	logger = zapLogger.Sugar()
}

func Logger() *zap.SugaredLogger {
	return logger
}

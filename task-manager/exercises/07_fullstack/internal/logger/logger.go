package logger

import (
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func Init(env string) {
	var logger *zap.Logger

	if env == "production" {
		logger, _ = zap.NewProduction()
	} else {
		logger, _ = zap.NewDevelopment()
	}

	log = logger.Sugar()
}

func Info(msg string, keysAndValues ...any) { log.Infow(msg, keysAndValues...) }

func Error(msg string, keysAndValues ...any) { log.Errorw(msg, keysAndValues...) }

func Fatal(msg string, keysAndValues ...any) { log.Fatalw(msg, keysAndValues...) }

func Debug(msg string, keysAndValues ...any) { log.Debugw(msg, keysAndValues...) }

func Warn(msg string, keysAndValues ...any) { log.Warnw(msg, keysAndValues...) }

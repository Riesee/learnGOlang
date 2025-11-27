package logger

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// DERS 3: Zap ile Logging
// ========================
//
// Zap özellikleri:
// - Uber tarafından geliştirildi
// - Çok hızlı (zero allocation)
// - Structured logging (JSON formatında)
// - Log seviyeleri: Debug, Info, Warn, Error, Fatal
//
// Production'da JSON, development'ta console formatı kullanılır

var log *zap.SugaredLogger

// Init logger'ı başlatır
func Init(env string) {
	var config zap.Config

	if env == "production" {
		// Production: JSON format, Info seviyesi
		config = zap.NewProductionConfig()
		config.EncoderConfig.TimeKey = "timestamp"
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	} else {
		// Development: Console format, Debug seviyesi
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}

	log = logger.Sugar()
}

// Wrapper fonksiyonlar - kolay kullanım için
func Debug(msg string, keysAndValues ...interface{}) {
	log.Debugw(msg, keysAndValues...)
}

func Info(msg string, keysAndValues ...interface{}) {
	log.Infow(msg, keysAndValues...)
}

func Warn(msg string, keysAndValues ...interface{}) {
	log.Warnw(msg, keysAndValues...)
}

func Error(msg string, keysAndValues ...interface{}) {
	log.Errorw(msg, keysAndValues...)
}

func Fatal(msg string, keysAndValues ...interface{}) {
	log.Fatalw(msg, keysAndValues...)
	os.Exit(1)
}

// GinLogger Gin için middleware
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// Request'i işle
		c.Next()

		// Response sonrası log
		latency := time.Since(start)
		status := c.Writer.Status()

		Info("HTTP Request",
			"status", status,
			"method", c.Request.Method,
			"path", path,
			"query", query,
			"ip", c.ClientIP(),
			"latency", latency.String(),
			"user-agent", c.Request.UserAgent(),
		)
	}
}

package main

// ============================================
// EGZERSİZ 4: Gin + Viper + Zap Birlikte
// ============================================
//
// Bu egzersizde:
// 1. Config'den port oku (Viper)
// 2. Logger kur (Zap)
// 3. API oluştur (Gin)
// 4. Her istekte log yaz
//
// ============================================

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Global logger - her yerden erişilebilir
var sugar *zap.SugaredLogger

func main() {
	viper.SetConfigFile("config.yaml")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	logger, _ := zap.NewDevelopment()
	sugar = logger.Sugar()

	defer logger.Sync()

	sugar.Infow("Uygulama baslatiliyor", "port", viper.GetString("server.port"))

	r := gin.Default()

	r.Use(LoggingMiddleware())

	r.GET("/health", healthHandler)
	r.GET("/info", infoHandler)
	r.POST("/echo", echoHandler)

	port := viper.GetString("server.port")
	r.Run(":" + port)
	// ============================================
	// BÖLÜM 1: CONFIG (Viper)
	// ============================================
	// GÖREV 1: Config dosyasını oku

	// ============================================
	// BÖLÜM 2: LOGGER (Zap)
	// ============================================
	// GÖREV 2: Logger oluştur ve global değişkene ata

	// GÖREV 3: Başlangıç logu yaz

	// ============================================
	// BÖLÜM 3: GIN ROUTER
	// ============================================
	// GÖREV 4: Gin router oluştur

	// GÖREV 5: Logging middleware ekle (aşağıda tanımlı)

	// ============================================
	// BÖLÜM 4: ROUTES
	// ============================================
	// GÖREV 6: Route'ları tanımla

	// ============================================
	// BÖLÜM 5: SERVER BAŞLAT
	// ============================================
	// GÖREV 7: Config'den port'u oku ve server'ı başlat
}

// ============================================
// HANDLERS
// ============================================

func healthHandler(c *gin.Context) {
	sugar.Debug("Health check yapıldı")
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
}

func infoHandler(c *gin.Context) {
	appName := viper.GetString("app.name")
	version := viper.GetString("app.version")

	sugar.Infow("Info endpoint çağrıldı",
		"app", appName,
		"version", version,
	)

	c.JSON(http.StatusOK, gin.H{
		"app":     appName,
		"version": version,
	})
}

func echoHandler(c *gin.Context) {
	var body map[string]any

	if err := c.ShouldBindJSON(&body); err != nil {
		sugar.Errorw("JSON parse hatası", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz JSON"})
		return
	}

	sugar.Infow("Echo isteği alındı", "body", body)
	c.JSON(http.StatusOK, gin.H{
		"echo": body,
	})
}

// ============================================
// MIDDLEWARE
// ============================================

// LoggingMiddleware her isteği loglar
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// İstek başlangıcı
		sugar.Infow("İstek alındı",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"ip", c.ClientIP(),
		)

		// Sonraki handler'a geç
		c.Next()

		// İstek tamamlandı
		sugar.Infow("İstek tamamlandı",
			"status", c.Writer.Status(),
		)
	}
}

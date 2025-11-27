package main

import (
	"log"

	"task-manager/internal/config"
	"task-manager/internal/handler"
	"task-manager/internal/logger"

	"github.com/gin-gonic/gin"
)

// DERS 1: Proje Yapısı ve Gin Temelleri
// =====================================
//
// Go projelerinde standart klasör yapısı:
//
// task-manager/
// ├── cmd/api/          → Uygulama entry point'leri
// ├── internal/         → Private uygulama kodu (dışarıdan import edilemez)
// │   ├── config/       → Konfigürasyon (Viper)
// │   ├── handler/      → HTTP handler'ları (Controller gibi)
// │   ├── model/        → Database modelleri
// │   ├── repository/   → Database işlemleri
// │   ├── service/      → Business logic
// │   └── middleware/   → Gin middleware'leri
// ├── pkg/              → Public, paylaşılabilir kod
// └── config.yaml       → Konfigürasyon dosyası

func main() {
	// 1. Config'i yükle (Viper)
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Config yüklenemedi:", err)
	}

	// 2. Logger'ı başlat (Zap)
	logger.Init(cfg.App.Env)
	logger.Info("Uygulama başlatılıyor",
		"env", cfg.App.Env,
		"port", cfg.Server.Port,
	)

	// 3. Gin router'ı oluştur
	if cfg.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()

	// 4. Global middleware'ler
	r.Use(gin.Recovery()) // Panic'leri yakala
	r.Use(logger.GinLogger()) // Request logging

	// 5. Route'ları tanımla
	handler.SetupRoutes(r)

	// 6. Server'ı başlat
	addr := ":" + cfg.Server.Port
	logger.Info("Server başlatıldı", "addr", addr)
	if err := r.Run(addr); err != nil {
		logger.Fatal("Server başlatılamadı", "error", err)
	}
}

package main

import (
	"go.uber.org/zap"
)

// ============================================
// EGZERSİZ 3: Zap ile Logging
// ============================================
//
// ZAP NEDİR?
// ----------
// Uber'in geliştirdiği çok hızlı logger
// fmt.Println yerine kullanılır
//
// NEDEN fmt.Println DEĞİL?
// ------------------------
// fmt.Println("User logged in")           → Sadece text
// logger.Info("User logged in", "id", 5)  → Structured (JSON)
//
// Production'da JSON log'lar:
// - Aranabilir (Elasticsearch, CloudWatch)
// - Filtrelenebilir (sadece ERROR'ları göster)
// - Metadata içerir (timestamp, level, fields)
//
// LOG SEVİYELERİ:
// ---------------
// Debug → Geliştirme detayları
// Info  → Normal bilgiler
// Warn  → Uyarılar
// Error → Hatalar
// Fatal → Kritik hata (program durur)
//
// ============================================

// GÖREV 1: Zap'ı import et
// İpucu: "go.uber.org/zap"

func main() {
	// GÖREV 2: Development logger oluştur
	logger, _ := zap.NewDevelopment()
	sugar := logger.Sugar()

	sugar.Debug("Debug mesajı")
	sugar.Info("Info mesajı")
	sugar.Warn("Uyarı mesajı")
	sugar.Error("Hata mesajı")

	sugar.Infow("Kullanıcı girişi",
		"user_id", 123,
		"email", "test@test.com",
	)

	logger.Sync()

	// GÖREV 3: Sugar logger al (kullanımı daha kolay)

	// GÖREV 4: Farklı seviyelerde log yaz

	// GÖREV 5: Structured logging (key-value)

	// GÖREV 6: Programı düzgün kapat
}

// ============================================
// ÇIKTI ÖRNEĞİ (Development):
// ============================================
// 2024-01-15T10:30:00.000+0300  DEBUG  Debug mesajı
// 2024-01-15T10:30:00.000+0300  INFO   Info mesajı
// 2024-01-15T10:30:00.000+0300  WARN   Uyarı mesajı
// 2024-01-15T10:30:00.000+0300  ERROR  Hata mesajı
// 2024-01-15T10:30:00.000+0300  INFO   Kullanıcı girişi  {"user_id": 123, "email": "test@test.com"}
// ============================================

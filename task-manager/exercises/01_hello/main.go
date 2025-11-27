package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// ============================================
// EGZERSİZ 1: İlk Gin Uygulaması
// ============================================
//
// HEDEF: Basit bir web server oluştur
//
// GIN NEDİR?
// ----------
// Gin = Go için web framework (Express.js gibi)
// HTTP isteklerini karşılar, route tanımlar, JSON döner
//
// TEMEL KAVRAMLAR:
// ----------------
// gin.Default()     → Yeni bir router oluşturur
// r.GET("/path", f) → GET isteği için handler tanımlar
// r.Run(":3000")    → Server'ı başlatır
// c.JSON(200, data) → JSON response döner
// gin.H{}           → map[string]any kısayolu (JSON için)
//
// ============================================

// GÖREV 1: Gin'i import et

func main() {
	fmt.Println("Gin Server başlatılıyor...")

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Merhaba Go!",
		})
	})

	r.GET("/selam/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"selam": "Merhaba " + name + "!",
		})
	})

	r.Run(":3001")

	// GÖREV 2: Gin router oluştur

	// GÖREV 3: Ana sayfa route'u ekle

	// GÖREV 4: /selam/:isim route'u ekle

	// GÖREV 5: Server'ı 3001 portunda başlat
}

// ============================================
// TEST ETMEK İÇİN:
// ============================================
// 1. Terminal'de: cd task-manager/exercises/01_hello
// 2. go run main.go
// 3. Tarayıcıda: http://localhost:3001
// 4. Tarayıcıda: http://localhost:3001/selam/Hilmi
// ============================================

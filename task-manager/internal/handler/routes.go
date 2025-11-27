package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DERS 4: Gin Routing ve Handler'lar
// ===================================
//
// Gin route tanımlama:
// - r.GET("/path", handler)
// - r.POST("/path", handler)
// - r.PUT("/path", handler)
// - r.DELETE("/path", handler)
//
// Route grupları:
// - api := r.Group("/api")
// - api.Use(middleware) // Gruba middleware ekle
//
// Path parametreleri:
// - /users/:id -> c.Param("id")
//
// Query parametreleri:
// - /users?page=1 -> c.Query("page")

// SetupRoutes tüm route'ları tanımlar
func SetupRoutes(r *gin.Engine) {
	// Health check - her API'de olmalı
	r.GET("/health", healthCheck)

	// API grubu
	api := r.Group("/api")
	{
		// Auth routes (public)
		auth := api.Group("/auth")
		{
			auth.POST("/register", register)
			auth.POST("/login", login)
		}

		// Task routes (protected - sonra middleware ekleyeceğiz)
		tasks := api.Group("/tasks")
		{
			tasks.GET("", getAllTasks)
			tasks.POST("", createTask)
			tasks.GET("/:id", getTaskByID)
			tasks.PUT("/:id", updateTask)
			tasks.DELETE("/:id", deleteTask)
		}
	}
}

// healthCheck sunucu durumunu döndürür
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Task Manager API is running",
	})
}

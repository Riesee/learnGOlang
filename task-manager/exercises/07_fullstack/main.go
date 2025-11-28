package main

import (
	"fullstack/internal/cache"
	"fullstack/internal/config"
	"fullstack/internal/database"
	"fullstack/internal/handler"
	"fullstack/internal/logger"
	"fullstack/internal/middleware"
	"fullstack/internal/model"
	"fullstack/internal/service"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	logger.Init(cfg.App.Env)
	logger.Info("Uygulama baslatiliyor",
		"env", cfg.App.Env,
		"port", cfg.Server.Port,
	)

	db, err := database.Connect(&cfg.Database)
	if err != nil {
		logger.Fatal("Database baglanilamadı", "error", err)
	}
	logger.Info("Database baglantısı yapıldı")

	if err := cache.Connect(cfg.Redis.Host, cfg.Redis.Port); err != nil {
		logger.Fatal("Redis baglanilamadı", "error", err)
	}
	logger.Info("Redis baglantısı yapıldı")

	db.AutoMigrate(&model.User{}, &model.Task{})
	logger.Info("Migration ile Tablolar olusturuldu")

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	logger.Info("Server başlatıldı ve Routers olusturuldu", "port", cfg.Server.Port)

	// Public Routes
	r.POST("/login", handler.Login(db, []byte(cfg.JWT.Secret)))
	r.POST("/register", handler.Register(db, []byte(cfg.JWT.Secret)))

	r.GET("/slow", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.JSON(200, gin.H{"message": "Yavaş cevap geldi!"})
	})
	// Protected Routes
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware(db, []byte(cfg.JWT.Secret)))
	protected.Use(middleware.RateLimitMiddleware(20, time.Minute))
	{
		protected.GET("/tasks", handler.GetTasks(db))
		protected.POST("/tasks", handler.PostTask(db))
		protected.GET("/tasks/:id", handler.GetTaskByID(db))
		protected.PUT("/tasks/:id", handler.UpdateTask(db))
		protected.DELETE("/tasks/:id", handler.DeleteTask(db))
	}

	service.SendNotificationToMany([]uint{1, 2, 3,4,5,6,7}, "Merhaba")
	r.Run(":" + cfg.Server.Port)
	

}

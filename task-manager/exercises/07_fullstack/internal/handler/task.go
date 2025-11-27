package handler

import (
	"encoding/json"
	"fmt"
	"fullstack/internal/cache"
	"fullstack/internal/logger"
	"fullstack/internal/model"
	"net/http"
	"time"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func GetTasks(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetUint("userID")
		cacheKey := fmt.Sprintf("user:%d:tasks", userID)

		if cached, err := cache.Get(cacheKey); err == nil {
			logger.Info("Tasklar redisten listelendi", "user_id", userID)
			c.Data(http.StatusOK, "application/json", []byte(cached))
			return
		}

		var tasks []model.Task

		if err := db.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		response := gin.H{"data": tasks, "message": "Tasklar listelendi"}
		jsonData, _ := json.Marshal(response)
		cache.Set(cacheKey, string(jsonData), 5*time.Minute)
		logger.Info("Tasklar listelendi", "user_id", userID)
		c.JSON(http.StatusOK, response)
	}
}

func PostTask(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.CreateTaskRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		userID := c.GetUint("userID")
		cacheKey := fmt.Sprintf("user:%d:tasks", userID)
		cache.Delete(cacheKey)

		task := model.Task{
			Title:       req.Title,
			Description: req.Description,
			Completed:   false,
			UserID:      userID,
		}

		if err := db.Create(&task).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		logger.Info("Task oluşturuldu", "task_id", task.ID, "user_id", userID)
		c.JSON(http.StatusCreated, gin.H{"data": task, "message": "Task oluşturuldu"})
	}
}

func GetTaskByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		userID := c.GetUint("userID")

		var task model.Task
		if err := db.Where("id = ? AND user_id = ?", id, userID).First(&task).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task bulunamadı"})
			return
		}
		logger.Info("Task çekildi", "task_id", task.ID, "user_id", userID)
		c.JSON(http.StatusOK, gin.H{"data": task, "message": "Task getirildi"})
	}
}

func UpdateTask(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		userID := c.GetUint("userID")
		cacheKey := fmt.Sprintf("user:%d:tasks", userID)
		cache.Delete(cacheKey)

		var req model.UpdateTaskRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var task model.Task
		if err := db.Where("id = ? AND user_id = ?", id, userID).First(&task).Error; err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "Bu task'a erişim yetkiniz yok"})
			return
		}

		if req.Title != nil {
			task.Title = *req.Title
		}
		if req.Description != nil {
			task.Description = *req.Description
		}
		if req.Completed != nil {
			task.Completed = *req.Completed
		}

		if err := db.Save(&task).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		logger.Info("Task güncellendi", "task_id", task.ID, "user_id", userID)
		c.JSON(http.StatusOK, gin.H{"data": task, "message": "Task güncellendi"})
	}
}

func DeleteTask(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		userID := c.GetUint("userID")
		cacheKey := fmt.Sprintf("user:%d:tasks", userID)
		cache.Delete(cacheKey)

		var task model.Task
		if err := db.Where("id = ? AND user_id = ?", id, userID).First(&task).Error; err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "Bu task'a erişim yetkiniz yok"})
			return
		}

		if err := db.Delete(&task).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		logger.Info("Task silindi", "task_id", task.ID, "user_id", userID)
		c.JSON(http.StatusOK, gin.H{"message": "Task silindi"})
	}
}

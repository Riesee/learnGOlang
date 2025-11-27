package handler

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// DERS 6: CRUD Operations
// ========================
//
// REST API standartları:
// - GET    /resources     -> List (200)
// - POST   /resources     -> Create (201)
// - GET    /resources/:id -> Read (200, 404)
// - PUT    /resources/:id -> Update (200, 404)
// - DELETE /resources/:id -> Delete (204, 404)
//
// Response kodları:
// - 200 OK
// - 201 Created
// - 204 No Content
// - 400 Bad Request
// - 401 Unauthorized
// - 404 Not Found
// - 500 Internal Server Error

// Task modeli (şimdilik burada, sonra model/ klasörüne taşıyacağız)
type Task struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	UserID      uint      `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateTaskRequest task oluşturma isteği
type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required,min=1,max=200"`
	Description string `json:"description" binding:"max=1000"`
}

// UpdateTaskRequest task güncelleme isteği
type UpdateTaskRequest struct {
	Title       *string `json:"title" binding:"omitempty,min=1,max=200"`
	Description *string `json:"description" binding:"omitempty,max=1000"`
	Completed   *bool   `json:"completed"`
}

// In-memory storage (geçici - sonra GORM ile değişecek)
var (
	tasks  = make(map[uint]*Task)
	nextID uint = 1
	mu     sync.RWMutex
)

// getAllTasks tüm taskları listeler
func getAllTasks(c *gin.Context) {
	mu.RLock()
	defer mu.RUnlock()

	// Map'i slice'a çevir
	result := make([]*Task, 0, len(tasks))
	for _, task := range tasks {
		result = append(result, task)
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  result,
		"count": len(result),
	})
}

// createTask yeni task oluşturur
func createTask(c *gin.Context) {
	var req CreateTaskRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	task := &Task{
		ID:          nextID,
		Title:       req.Title,
		Description: req.Description,
		Completed:   false,
		UserID:      1, // TODO: JWT'den al
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tasks[nextID] = task
	nextID++
	mu.Unlock()

	c.JSON(http.StatusCreated, gin.H{"data": task})
}

// getTaskByID tek task döndürür
func getTaskByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz ID"})
		return
	}

	mu.RLock()
	task, exists := tasks[uint(id)]
	mu.RUnlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task bulunamadı"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// updateTask task günceller
func updateTask(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz ID"})
		return
	}

	var req UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	task, exists := tasks[uint(id)]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task bulunamadı"})
		return
	}

	// Partial update - sadece gönderilen alanları güncelle
	if req.Title != nil {
		task.Title = *req.Title
	}
	if req.Description != nil {
		task.Description = *req.Description
	}
	if req.Completed != nil {
		task.Completed = *req.Completed
	}
	task.UpdatedAt = time.Now()

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// deleteTask task siler
func deleteTask(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz ID"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if _, exists := tasks[uint(id)]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task bulunamadı"})
		return
	}

	delete(tasks, uint(id))
	c.Status(http.StatusNoContent)
}

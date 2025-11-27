package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DERS 5: Request/Response Handling
// ==================================
//
// Request body okuma:
// - c.ShouldBindJSON(&struct) - JSON body'yi struct'a bind et
// - c.ShouldBind(&struct) - Content-Type'a göre otomatik bind
//
// Response gönderme:
// - c.JSON(status, data) - JSON response
// - c.String(status, text) - Plain text
// - c.HTML(status, template, data) - HTML template
//
// Validation:
// - Struct tag'leri: `binding:"required,min=3,max=100"`
// - ShouldBindJSON otomatik validate eder

// RegisterRequest kayıt isteği
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required,min=2"`
}

// LoginRequest giriş isteği
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// AuthResponse auth yanıtı
type AuthResponse struct {
	Token string `json:"token"`
	User  UserResponse `json:"user"`
}

// UserResponse kullanıcı bilgisi
type UserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

// register yeni kullanıcı kaydı
func register(c *gin.Context) {
	var req RegisterRequest

	// JSON body'yi bind et ve validate et
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Validation hatası",
			"details": err.Error(),
		})
		return
	}

	// TODO: Database'e kaydet
	// TODO: Password hash'le (bcrypt)
	// TODO: JWT token oluştur

	// Şimdilik mock response
	c.JSON(http.StatusCreated, AuthResponse{
		Token: "mock-jwt-token",
		User: UserResponse{
			ID:    1,
			Email: req.Email,
			Name:  req.Name,
		},
	})
}

// login kullanıcı girişi
func login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Validation hatası",
			"details": err.Error(),
		})
		return
	}

	// TODO: Database'den kullanıcıyı bul
	// TODO: Password kontrol et
	// TODO: JWT token oluştur

	// Şimdilik mock response
	c.JSON(http.StatusOK, AuthResponse{
		Token: "mock-jwt-token",
		User: UserResponse{
			ID:    1,
			Email: req.Email,
			Name:  "Test User",
		},
	})
}

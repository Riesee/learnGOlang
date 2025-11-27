package handler

import (
	"fullstack/internal/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

func Login(db *gorm.DB, jwtSecret []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var dbUser model.User
		if err := db.First(&dbUser, "email = ?", req.Email).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Kullanıcı bulunamadı"})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(req.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Şifre yanlış"})
			return
		}

		claims := &model.Claims{
			UserID: dbUser.ID,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtSecret)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Token oluşturulamadı"})
			return
		}
		db.Model(&dbUser).Update("token", tokenString)

		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	}
}

func Register(db *gorm.DB, jwtSecret []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var existingUser model.User
		if err := db.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Bu email zaten kayıtlı"})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Şifre oluşturulamadı"})
			return
		}

		user := model.User{
			Email:    req.Email,
			Password: string(hashedPassword),
			Name:     req.Name,
		}

		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Kullanıcı oluşturulamadı"})
			return
		}

		token, _ := generateToken(user, jwtSecret)

		c.JSON(http.StatusCreated, gin.H{
			"message": "Kullanıcı oluşturuldu",
			"token":   token,
			"user":    user,
		})
	}
}

func generateToken(user model.User, jwtSecret []byte) (string, error) {
	claims := &model.Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

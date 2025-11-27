package middleware

import (
	"fullstack/internal/model"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func AuthMiddleware(db *gorm.DB, jwtSecret []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Header'dan token al
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token gerekli"})
			c.Abort()
			return
		}

		// 2. "Bearer " prefix'ini kaldır
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// 3. Token'ı parse et ve doğrula
		claims := &model.Claims{}
		userExist := false
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		var user model.User
		if err := db.First(&user, claims.UserID).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Kullanıcı bulunamadı"})
			c.Abort()
			return
		}
		userExist = true

		if err != nil || !token.Valid || !userExist || *&user.Token != tokenString {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Gecersiz token veya oturum sonlandırılmış"})
			c.Abort()
			return
		} else {
			c.Set("userID", claims.UserID)
			c.Next()
		}
	}
}

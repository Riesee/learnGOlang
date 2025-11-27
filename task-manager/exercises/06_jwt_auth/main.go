package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// ============================================
// EGZERSİZ 6: JWT Authentication
// ============================================
//
// JWT NEDİR?
// ----------
// JSON Web Token - kullanıcı kimliğini doğrulamak için
// Login → Token al → Her istekte token gönder
//
// AKIŞ:
// -----
// 1. POST /register → Kullanıcı kaydı (şifre hash'lenir)
// 2. POST /login    → Email/şifre doğru → JWT token döner
// 3. GET /profile   → Token ile istek → Kullanıcı bilgisi
//
// TOKEN YAPISI:
// -------------
// Header.Payload.Signature
// eyJhbGc...(header).eyJzdWI...(payload).SflKxw...(signature)
//
// ============================================

// JWT secret key (production'da env variable olmalı!)
var jwtSecret = []byte("super-secret-key-change-me")

// User modeli (normalde database'de olur)
type User struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"` // JSON'da görünmez
	Name     string `json:"name"`
}

// In-memory user storage (normalde database)
var users = []User{}
var nextUserID uint = 1

// Request/Response struct'ları
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

// JWT Claims - token içinde saklanacak bilgiler
type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func main() {
	r := gin.Default()

	// Public routes (token gerekmez)
	r.POST("/register", register)
	r.POST("/login", login)

	// Protected routes (token gerekir)
	protected := r.Group("/")
	protected.Use(AuthMiddleware()) // Middleware ekle
	{
		protected.GET("/profile", getProfile)
		protected.GET("/users", getAllUsers)
	}

	r.Run(":3003")
}

// ============================================
// GÖREV 1: register handler
// ============================================
func register(c *gin.Context) {
	var req RegisterRequest

	// 1. Request'i bind et
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. Email zaten var mı kontrol et
	for _, u := range users {
		if u.Email == req.Email {
			c.JSON(http.StatusConflict, gin.H{"error": "Email zaten kayıtlı"})
			return
		}
	}

	// GÖREV: Şifreyi hash'le
	// İpucu: bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	// BURAYA KOD YAZ
	hashedPassword := "" // Bu satırı değiştir

	// 3. User oluştur ve kaydet
	user := User{
		ID:       nextUserID,
		Email:    req.Email,
		Password: hashedPassword,
		Name:     req.Name,
	}
	users = append(users, user)
	nextUserID++

	// 4. Token oluştur ve döndür
	token, _ := generateToken(user)

	c.JSON(http.StatusCreated, AuthResponse{
		Token: token,
		User:  user,
	})
}

// ============================================
// GÖREV 2: login handler
// ============================================
func login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// GÖREV: Kullanıcıyı bul ve şifreyi doğrula
	// 1. users içinde email'e göre ara
	// 2. bcrypt.CompareHashAndPassword ile şifre kontrolü
	// 3. Doğruysa token oluştur ve döndür
	// 4. Yanlışsa 401 Unauthorized döndür
	// BURAYA KOD YAZ

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Geçersiz email veya şifre"})
}

// ============================================
// GÖREV 3: generateToken fonksiyonu
// ============================================
func generateToken(user User) (string, error) {
	// GÖREV: JWT token oluştur
	// 1. Claims oluştur (UserID, Email, ExpiresAt)
	// 2. jwt.NewWithClaims ile token oluştur
	// 3. token.SignedString ile imzala
	// BURAYA KOD YAZ

	claims := Claims{
		UserID: user.ID,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 24 saat
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ============================================
// GÖREV 4: AuthMiddleware
// ============================================
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// GÖREV: Token'ı doğrula
		// 1. Authorization header'dan token'ı al
		// 2. "Bearer " prefix'ini kaldır
		// 3. Token'ı parse et ve doğrula
		// 4. Claims'i context'e ekle
		// 5. Hata varsa 401 döndür

		// 1. Header'dan token al
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token gerekli"})
			c.Abort()
			return
		}

		// 2. "Bearer " prefix'ini kaldır
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// GÖREV: Token'ı parse et ve doğrula
		// İpucu:
		// claims := &Claims{}
		// token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		//     return jwtSecret, nil
		// })
		// BURAYA KOD YAZ

		// Geçici olarak geç (bunu sil ve yukarıdaki kodu yaz)
		_ = tokenString
		c.Next()
	}
}

// ============================================
// Protected handlers
// ============================================
func getProfile(c *gin.Context) {
	// Context'ten user bilgisini al
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Kullanıcı bulunamadı"})
		return
	}

	// User'ı bul
	for _, u := range users {
		if u.ID == userID.(uint) {
			c.JSON(http.StatusOK, gin.H{"user": u})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Kullanıcı bulunamadı"})
}

func getAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"users": users})
}

// ============================================
// TEST KOMUTLARI:
// ============================================
// 1. Register:
// curl -X POST http://localhost:3003/register \
//   -H "Content-Type: application/json" \
//   -d '{"email":"test@test.com","password":"123456","name":"Test User"}'
//
// 2. Login:
// curl -X POST http://localhost:3003/login \
//   -H "Content-Type: application/json" \
//   -d '{"email":"test@test.com","password":"123456"}'
//
// 3. Profile (token ile):
// curl http://localhost:3003/profile \
//   -H "Authorization: Bearer <TOKEN>"
// ============================================

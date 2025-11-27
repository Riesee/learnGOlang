package model

import (
	"gorm.io/gorm"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"-"`
	Name     string `gorm:"not null" json:"name"`
	Token    string `gorm:"" json:"-"`
}

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required,min=2"`
}

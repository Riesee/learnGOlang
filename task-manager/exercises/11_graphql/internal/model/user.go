package dbmodel

import "gorm.io/gorm"

// User - veritabanı modeli (GORM)
type User struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Name     string `gorm:"not null"`
	Tasks    []Task `gorm:"foreignKey:UserID"`
}

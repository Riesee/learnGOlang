package dbmodel

import "gorm.io/gorm"

// Task - veritabanı modeli (GORM)
type Task struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string
	Completed   bool   `gorm:"default:false"`
	UserID      uint   `gorm:"not null"`
	User        User   `gorm:"foreignKey:UserID"`
}

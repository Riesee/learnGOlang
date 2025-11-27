package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title"`
	Description string `gorm:"not null" json:"description"`
	Completed   bool   `gorm:"not null" json:"completed"`
	UserID      uint   `gorm:"not null" json:"user_id"`
}

type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateTaskRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Completed   *bool   `json:"completed"`
}

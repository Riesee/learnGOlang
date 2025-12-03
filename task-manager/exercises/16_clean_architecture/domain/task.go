package domain

import "time"

type Task struct {
	ID   int
	Title string
	Description string
	Completed bool
	UserID uint
	CreatedAt time.Time
}

type TaskRepository interface {
	GetAll(userID uint) ([]Task, error)
	GetByID(id int) (*Task, error)
	Create(task *Task) error
	Update(task *Task) error
	Delete(id int) error
}
package service

import (
	"task-manager/exercises/16_clean_architecture/domain"
)

type TaskService struct {
	repo domain.TaskRepository
}

func NewTaskService(repo domain.TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}
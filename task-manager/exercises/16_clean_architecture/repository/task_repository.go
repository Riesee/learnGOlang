package repository

import "task-manager/exercises/16_clean_architecture/domain"

type GormTaskRepository struct {
	tasks []domain.Task
	nextID uint
}

func NewGormTaskRepository() *GormTaskRepository {
	return &GormTaskRepository{
		tasks: []domain.Task{},
		nextID: 1,
	}
}

func (r *GormTaskRepository) GetAll(userID uint) ([]domain.Task, error) {
	var result []domain.Task
	for _, task := range r.tasks {
		if task.UserID == userID {
			result = append(result, task)
		}
	}
	return result, nil
}

func (r *GormTaskRepository) GetByID(id int) (*domain.Task, error) {
	for _, task := range r.tasks {
		if task.ID == id {
			return &task, nil
		}
	}
	return nil, nil
}

func (r *GormTaskRepository) Create(task *domain.Task) error {
	task.ID = int(r.nextID)
	r.nextID++
	r.tasks = append(r.tasks, *task)
	return nil
}

func (r *GormTaskRepository) Update(task *domain.Task) error {
	for i, t := range r.tasks {
		if t.ID == task.ID {
			r.tasks[i] = *task
			return nil
		}
	}
	return nil
}

func (r *GormTaskRepository) Delete(id int) error {
	for i, task := range r.tasks {
		if task.ID == id {
			r.tasks = append(r.tasks[:i], r.tasks[i+1:]...)
			return nil
		}
	}
	return nil
}
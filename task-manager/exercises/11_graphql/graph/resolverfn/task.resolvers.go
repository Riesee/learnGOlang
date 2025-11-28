package resolverfn

import (
	"context"
	"strconv"

	"graphql-demo/graph/model"
	"graphql-demo/internal/middleware"
	dbmodel "graphql-demo/internal/model"
)

// ============================================================
// GÖREV 1: Tasks Resolver'ını Yaz
// ============================================================

func (r *queryResolver) Tasks(ctx context.Context) ([]*model.Task, error) {
	db := r.DB

	var items []dbmodel.Task
	if err := db.Find(&items).Error; err != nil {
		return nil, err
	}

	var tasks []*model.Task
	for _, item := range items {
		tasks = append(tasks, dbTaskToGraphQL(&item))
	}

	return tasks, nil
}

// ============================================================
// GÖREV 2: Task Resolver'ını Yaz (ID ile tek task getir)
// ============================================================

func (r *queryResolver) Task(ctx context.Context, id string) (*model.Task, error) {
	db := r.DB
	idStr, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, err
	}
	var item dbmodel.Task
	if err := db.First(&item, idStr).Error; err != nil {
		return nil, err
	}
	task := dbTaskToGraphQL(&item)

	return task, nil

}

// ============================================================
// GÖREV 3: CreateTask Mutation'ını Yaz
// ============================================================

func (r *mutationResolver) CreateTask(ctx context.Context, input model.CreateTaskInput) (*model.Task, error) {
	db := r.DB

	userId := middleware.GetUserIDFromContext(ctx)

	task := dbmodel.Task{
		Title:       input.Title,
		Completed: false,
		UserID:      userId,
	}
	if input.Description != nil {
		task.Description = *input.Description
	}

	if err := db.Create(&task).Error; err != nil {
		return nil, err
	}

	data := dbTaskToGraphQL(&task)

	return data, nil
}

// ============================================================
// GÖREV 4: UpdateTask Mutation'ını Yaz
// ============================================================

func (r *mutationResolver) UpdateTask(ctx context.Context, id string, input model.UpdateTaskInput) (*model.Task, error) {
	db := r.DB
	var task dbmodel.Task
	if err := db.First(&task, id).Error; err != nil {
		return nil, err
	}

	if input.Title != nil {
		task.Title = *input.Title
	}
	if input.Description != nil {
		task.Description = *input.Description
	}
	if input.Completed != nil {
		task.Completed = *input.Completed
	}

	if err := db.Save(&task).Error; err != nil {
		return nil, err
	}

	data := dbTaskToGraphQL(&task)

	return data, nil
}

// ============================================================
// GÖREV 5: DeleteTask Mutation'ını Yaz
// ============================================================
func (r *mutationResolver) DeleteTask(ctx context.Context, id string) (bool, error) {
	db := r.DB

	task := dbmodel.Task{}
	if err := db.First(&task, id).Error; err != nil {
		return false, err
	}
	
	result := db.Delete(&task)
	if result.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}

// ============================================================
// YARDIMCI FONKSİYON
// ============================================================
// DB Task modelini GraphQL Task modeline dönüştürür
//
// GÖREV 6: Bu fonksiyonu anla ve tamamla
// Şu anda eksik - sen doldur!
//
// İpucu: dbmodel.Task şöyle:
//   type Task struct {
//       gorm.Model          // ID, CreatedAt, UpdatedAt, DeletedAt içerir
//       Title       string
//       Description string
//       Completed   bool
//       UserID      uint
//   }
//
// GraphQL model.Task şöyle:
//   type Task struct {
//       ID          string
//       Title       string
//       Description *string   // pointer!
//       Completed   bool
//       UserID      string
//       CreatedAt   *time.Time
//       UpdatedAt   *time.Time
//   }
//


func dbTaskToGraphQL(t *dbmodel.Task) *model.Task {
	data := &model.Task{
		Title:       t.Title,
		ID:          strconv.FormatUint(uint64(t.ID), 10),
		Completed:   t.Completed,
		UpdatedAt:   &t.UpdatedAt,
		Description: &t.Description,
		UserID:      strconv.FormatUint(uint64(t.UserID), 10),
		CreatedAt:   &t.CreatedAt,
	}
	return data
}


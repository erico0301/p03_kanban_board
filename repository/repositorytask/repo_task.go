package repositorytask

import (
	"errors"
	"p03_kanban_board/constant"
	"p03_kanban_board/entity"

	"gorm.io/gorm"
)

type RepositoryTask interface {
	Create(task entity.Task) (entity.Task, error)
	IsCategoryExist(categoryID uint) error
	Gets() ([]entity.Task, error)
	Update(task entity.Task) (entity.Task, error)
	Delete(task entity.Task) error
}

type repository struct {
	db *gorm.DB
}

func (r *repository) Create(task entity.Task) (entity.Task, error) {
	err := r.db.Create(&task).Error
	if err != nil {
		return entity.Task{}, err
	}
	return task, nil
}

func (r *repository) IsCategoryExist(categoryID uint) error {
	category := new(entity.Category)
	err := r.db.Where("id = ?", categoryID).First(&category).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return constant.ErrorCategoryDoesNotExists
		}
		return err
	}
	return nil
}

func (r *repository) IsOwner(taskID uint) (entity.Task, error) {
	task := new(entity.Task)
	err := r.db.Where("id = ?", taskID).First(&task).Error
	if err != nil {
		return entity.Task{}, err
	}
	return *task, nil
}

func (r *repository) Gets() ([]entity.Task, error) {
	var tasks []entity.Task
	err := r.db.Preload("User").Preload("Category").Find(&tasks).Error
	if err != nil {
		return []entity.Task{}, err
	}
	return tasks, nil
}

func (r *repository) Update(task entity.Task) (entity.Task, error) {

	// validate ownership
	taskCheck, err := r.IsOwner(task.ID)
	if err != nil {
		return entity.Task{}, err
	}

	if taskCheck.UserID != task.UserID {
		return entity.Task{}, constant.ErrorOwnership
	}

	err = r.db.Updates(&task).First(&task).Error
	if err != nil {
		return entity.Task{}, err
	}
	return task, nil
}

func (r *repository) Delete(task entity.Task) error {
	checkTask, err := r.IsOwner(task.ID)
	if err != nil {
		return err
	}

	if checkTask.UserID != task.UserID {
		return constant.ErrorOwnership
	}

	err = r.db.First(&task).Where("id = ?", task.ID).Delete(&task).Error
	if err != nil {
		return err
	}
	return nil
}

func New(db *gorm.DB) RepositoryTask {
	return &repository{db: db}
}

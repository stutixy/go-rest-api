package repository

import (
	"go-rest-api/models"

	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task *models.Task) error
	GetAllTasks() ([]models.Task, error)
	GetTaskByID(id uint) (*models.Task, error)
	UpdateTask(task *models.Task) error
	DeleteTask(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) CreateTask(task *models.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepository) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) GetTaskByID(id uint) (*models.Task, error) {
	var task models.Task
	err := r.db.First(&task, id).Error
	return &task, err
}

func (r *taskRepository) UpdateTask(task *models.Task) error {
	return r.db.Save(task).Error
}

func (r *taskRepository) DeleteTask(id uint) error {
	return r.db.Delete(&models.Task{}, id).Error
}

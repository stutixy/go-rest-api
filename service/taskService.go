package service

import (
	"go-rest-api/models"
	"go-rest-api/repository"
)

type TaskService interface {
	CreateTask(task *models.Task) error
	GetAllTasks() ([]models.Task, error)
	GetTaskByID(id uint) (*models.Task, error)
	UpdateTask(task *models.Task) error
	DeleteTask(id uint) error
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{repo}
}

func (s *taskService) CreateTask(task *models.Task) error {
	return s.repo.CreateTask(task)
}

func (s *taskService) GetAllTasks() ([]models.Task, error) {
	return s.repo.GetAllTasks()
}

func (s *taskService) GetTaskByID(id uint) (*models.Task, error) {
	return s.repo.GetTaskByID(id)
}

func (s *taskService) UpdateTask(task *models.Task) error {
	return s.repo.UpdateTask(task)
}

func (s *taskService) DeleteTask(id uint) error {
	return s.repo.DeleteTask(id)
}

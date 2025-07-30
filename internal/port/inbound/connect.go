package inbound

import "task-api/internal/domain"

// Connect : Interface for the service
type Connect interface {
	CreateTask(task domain.Task) (string, error)
	GetByID(id string) (domain.Task, error)
	GetAll() ([]domain.Task, error)
	UpdateTask(task domain.Task) error
	Delete(id string) error
}

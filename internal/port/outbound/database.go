package outbound

import "task-api/internal/domain"

// Database : Interface for database structure in persistence
type Database interface {
	CreateTask(task domain.Task) (string, error)
	GetByID(id string) (domain.Task, error)
	GetAll() ([]domain.Task, error)
	UpdateTask(task domain.Task) error
	Delete(id string) error
}

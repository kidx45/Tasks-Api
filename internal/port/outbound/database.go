package outbound

import (
	"context"
	"task-api/internal/domain"
)

// Database : Interface for database structure in persistence
type Database interface {
	CreateTask(c context.Context, task domain.Task) (string, error)
	GetByID(c context.Context, id string) (domain.Task, error)
	GetAll(c context.Context) ([]domain.Task, error)
	UpdateTask(c context.Context, task domain.Task) error
	Delete(c context.Context, id string) error
}

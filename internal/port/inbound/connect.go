package inbound

import (
	"context"
	"task-api/internal/domain"
)

// Connect : Interface for the service
type Connect interface {
	CreateTask(c context.Context, task domain.Task) (string, error)
	GetByID(c context.Context, id string) (domain.Task, error)
	GetAll(c context.Context) ([]domain.Task, error)
	UpdateTask(c context.Context, task domain.Task) error
	Delete(c context.Context, id string) error
}

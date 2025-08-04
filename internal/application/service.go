package service

import (
	"context"
	"task-api/internal/domain"
	"task-api/internal/port/inbound"
	"task-api/internal/port/outbound"
)

// Connect : For collect data from persistence (data collection)
type Connect struct {
	repo outbound.Database
}

// NewConnect : To form connection
func NewConnect(repo outbound.Database) inbound.Connect {
	return Connect{repo: repo}
}

// CreateTask : To create a task
func (s Connect) CreateTask(c context.Context, task domain.Task) (string, error) {
	return s.repo.CreateTask(c, task)
}

// GetByID : To retrive task by id
func (s Connect) GetByID(c context.Context, id string) (domain.Task, error) {
	return s.repo.GetByID(c, id)
}

// GetAll : TO retrive all tasks
func (s Connect) GetAll(c context.Context) ([]domain.Task, error) {
	return s.repo.GetAll(c)
}

// UpdateTask : To update contents of a task by id
func (s Connect) UpdateTask(c context.Context, task domain.Task) error {
	return s.repo.UpdateTask(c, task)
}

// Delete : To delete task by id
func (s Connect) Delete(c context.Context, id string) error {
	return s.repo.Delete(c,id)
}

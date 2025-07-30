package service

import (
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
	return &Connect{repo: repo}
}

// CreateTask : To create a task
func (s *Connect) CreateTask(task domain.Task) (string, error) {
	return s.repo.CreateTask(task)
}

// GetByID : To retrive task by id
func (s *Connect) GetByID(id string) (domain.Task, error) {
	return s.repo.GetByID(id)
}

// GetAll : TO retrive all tasks
func (s *Connect) GetAll() ([]domain.Task, error) {
	return s.repo.GetAll()
}

// UpdateTask : To update contents of a task by id
func (s *Connect) UpdateTask(task domain.Task) error {
	return s.repo.UpdateTask(task)
}

// Delete : To delete task by id
func (s *Connect) Delete(id string) error {
	return s.repo.Delete(id)
}

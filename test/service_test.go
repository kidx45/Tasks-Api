package test

import (
	service "task-api/internal/application"
	"task-api/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock for outbound.Database
type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) CreateTask(task domain.Task) (string, error) {
	args := m.Called(task)
	return args.String(0), args.Error(1)
}
func (m *MockDatabase) GetByID(id string) (domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Task), args.Error(1)
}
func (m *MockDatabase) GetAll() ([]domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]domain.Task), args.Error(1)
}
func (m *MockDatabase) UpdateTask(task domain.Task) error {
	args := m.Called(task)
	return args.Error(0)
}
func (m *MockDatabase) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestServiceCreateTask(t *testing.T) {
	mock := new(MockDatabase)
	service := service.NewConnect(mock)
	task := domain.Task{Title: "test1", Description: "testing123"}
	mock.On("CreateTask", task).Return("1", nil)

	id, err := service.CreateTask(task)
	assert.NoError(t, err)
	assert.Equal(t, "1", id)
} 

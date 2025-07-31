package test

import (
	service "task-api/internal/application"
	"task-api/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock for outbound.Database
type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) CreateTask(task domain.Task) (string, error) {
	args := m.Called(task)
	return args.String(0), args.Error(1)
}
func (m *MockRepo) GetByID(id string) (domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Task), args.Error(1)
}
func (m *MockRepo) GetAll() ([]domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]domain.Task), args.Error(1)
}
func (m *MockRepo) UpdateTask(task domain.Task) error {
	args := m.Called(task)
	return args.Error(0)
}
func (m *MockRepo) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestServiceCreateTask(t *testing.T) {
	mockRepo := new(MockRepo)
	service := service.NewConnect(mockRepo)
	task := domain.Task{Title: "Test", Description: "Desc"}
	mockRepo.On("CreateTask", task).Return("123", nil)

	id, err := service.CreateTask(task)
	assert.NoError(t, err)
	assert.Equal(t, "123", id)
}

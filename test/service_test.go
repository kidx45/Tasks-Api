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
	task := domain.Task{
		Title: "test1", Description: "testing123",
	}
	mock.On("CreateTask", task).Return("1", nil)

	id, err := service.CreateTask(task)
	assert.NoError(t, err)
	assert.Equal(t, "1", id)
}

func TestServiceGetAllTask(t *testing.T) {
	mock := new(MockDatabase)
	service := service.NewConnect(mock)
	
	task := []domain.Task {
		{ID: "1", Title: "test1", Description: "testing123"},
		{ID: "2", Title: "test2", Description: "testing456"}, 
	}

	mock.On("GetAll").Return(task,nil)
	tasks,err := service.GetAll()

	assert.Equal(t, tasks,task)
	assert.NoError(t, err)
}

func TestServiceGetByID(t *testing.T) {
	mock := new(MockDatabase)
	service := service.NewConnect(mock)

	task := []domain.Task {
		{ID: "1", Title: "test1", Description: "testing123"},
		{ID: "2", Title: "test2", Description: "testing456"}, 
	}

	mock.On("GetByID", "1").Return(task[0],nil)
	test, err := service.GetByID("1")
	
	assert.Equal(t, test, task[0])
	assert.NoError(t, err)
}

func TestServiceUpdateTask(t *testing.T) {
	mock := new(MockDatabase)
	service := service.NewConnect(mock)

	task := domain.Task{
		ID: "1", Title: "test1", Description: "testing123",
	}
	mock.On("UpdateTask",task).Return(nil)
	err := service.UpdateTask(task)

	assert.NoError(t, err)
}

func TestServiceDelete(t *testing.T) {
	mock := new(MockDatabase)
	service := service.NewConnect(mock)

	ID := "1"
	mock.On("Delete",ID).Return(nil)
	err := service.Delete(ID)

	assert.NoError(t ,err)
}
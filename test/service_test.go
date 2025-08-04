package test

import (
	"context"
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

func (m *MockDatabase) CreateTask(c context.Context, task domain.UserInput) (string, error) {
	args := m.Called(c, task)
	return args.String(0), args.Error(1)
}
func (m *MockDatabase) GetByID(c context.Context, id string) (domain.Task, error) {
	args := m.Called(c, id)
	return args.Get(0).(domain.Task), args.Error(1)
}
func (m *MockDatabase) GetAll(c context.Context) ([]domain.Task, error) {
	args := m.Called(c)
	return args.Get(0).([]domain.Task), args.Error(1)
}
func (m *MockDatabase) UpdateTask(c context.Context, task domain.Task) error {
	args := m.Called(c, task)
	return args.Error(0)
}
func (m *MockDatabase) Delete(c context.Context, id string) error {
	args := m.Called(c, id)
	return args.Error(0)
}

func TestServiceCreateTask(t *testing.T) {
	mock := new(MockDatabase)
	service := service.NewConnect(mock)
	task := domain.UserInput{
		Title: "test1", Description: "testing123",
	}
	mock.On("CreateTask", context.Background(),task).Return("1", nil)

	id, err := service.CreateTask(context.Background(), task)
	assert.NoError(t, err)
	assert.Equal(t, "1", id)
}

func TestServiceGetAllTask(t *testing.T) {
	mock := new(MockDatabase)
	service := service.NewConnect(mock)

	task := []domain.Task{
		{ID: "1", Title: "test1", Description: "testing123"},
		{ID: "2", Title: "test2", Description: "testing456"},
	}

	mock.On("GetAll", context.Background()).Return(task, nil)
	tasks, err := service.GetAll(context.Background())

	assert.Equal(t, tasks, task)
	assert.NoError(t, err)
}

func TestServiceGetByID(t *testing.T) {
	mock := new(MockDatabase)
	service := service.NewConnect(mock)

	task := []domain.Task{
		{ID: "1", Title: "test1", Description: "testing123"},
		{ID: "2", Title: "test2", Description: "testing456"},
	}

	mock.On("GetByID",context.Background(), "1").Return(task[0], nil)
	test, err := service.GetByID(context.Background(), "1")

	assert.Equal(t, test, task[0])
	assert.NoError(t, err)
}

func TestServiceUpdateTask(t *testing.T) {
	mock := new(MockDatabase)
	service := service.NewConnect(mock)

	task := domain.Task{
		ID: "1", Title: "test1", Description: "testing123",
	}
	mock.On("UpdateTask",context.Background(), task).Return(nil)
	err := service.UpdateTask(context.Background(), task)

	assert.NoError(t, err)
}

func TestServiceDelete(t *testing.T) {
	mock := new(MockDatabase)
	service := service.NewConnect(mock)

	ID := "1"
	mock.On("Delete",context.Background(), ID).Return(nil)
	err := service.Delete(context.Background(), ID)

	assert.NoError(t, err)
}

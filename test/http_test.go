package test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	httpy "task-api/internal/adapter/inbound/http"
	"task-api/internal/domain"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock service
type MockConnect struct {
	mock.Mock
}

func (m *MockConnect) CreateTask(c context.Context, task domain.Task) (string, error) {
	args := m.Called(task)
	return args.String(0), args.Error(1)
}
func (m *MockConnect) GetByID(c context.Context, id string) (domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Task), args.Error(1)
}
func (m *MockConnect) GetAll(c context.Context) ([]domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]domain.Task), args.Error(1)
}
func (m *MockConnect) UpdateTask(c context.Context, task domain.Task) error {
	args := m.Called(task)
	return args.Error(0)
}
func (m *MockConnect) Delete(c context.Context, id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateTask(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockConnect)
	handler := httpy.HTTPHandler{Service: mockService}
	r := gin.New()
	r.POST("/tasks", handler.CreateTask)

	var test domain.Task
	task := domain.Task{Title: "Test", Description: "Desc"}
	mockService.On("CreateTask", context.Background(), task).Return("123", nil)

	body, _ := json.Marshal(task)
	request, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	code := httptest.NewRecorder()
	r.ServeHTTP(code, request)

	assert.Equal(t, http.StatusCreated, code.Code)
	err := json.Unmarshal(code.Body.Bytes(), &test)
	assert.NoError(t, err)
	assert.NotNil(t, test.ID)
	assert.Equal(t, test.Title, task.Title)
	assert.Equal(t, test.Description, task.Description)
}

func TestGetAll(t *testing.T) {
	gin.SetMode(gin.TestMode)
	MockConnect := new(MockConnect)
	handler := httpy.HTTPHandler{Service: MockConnect}
	router := gin.New()
	router.GET("/tasks", handler.GetAll)

	var tasks []domain.Task
	task := []domain.Task{
		{ID: "1", Title: "test1", Description: "testing123"},
		{ID: "2", Title: "test2", Description: "testing456"},
	}

	MockConnect.On("GetAll").Return(task, nil)
	request, _ := http.NewRequest("GET", "/tasks", nil)
	request.Header.Set("Content-Type", "application/json")
	code := httptest.NewRecorder()
	router.ServeHTTP(code, request)

	assert.Equal(t, http.StatusOK, code.Code)
	err := json.Unmarshal(code.Body.Bytes(), &tasks)
	assert.NoError(t, err)
	assert.Equal(t, tasks, task)
}

func TestGetByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	MockConnect := new(MockConnect)
	handler := httpy.HTTPHandler{Service: MockConnect}
	router := gin.New()
	router.GET("/tasks/:id", handler.GetByID)

	var test domain.Task
	task := []domain.Task{
		{ID: "1", Title: "test1", Description: "testing123"},
		{ID: "2", Title: "test2", Description: "testing456"},
	}

	MockConnect.On("GetByID", task[0].ID).Return(task[0], nil)
	request, _ := http.NewRequest("GET", "/tasks/1", nil)
	request.Header.Set("Content-Type", "application/json")
	code := httptest.NewRecorder()
	router.ServeHTTP(code, request)

	assert.Equal(t, http.StatusOK, code.Code)
	err := json.Unmarshal(code.Body.Bytes(), &test)
	assert.NoError(t, err)
	assert.Equal(t, test, task[0])
}

func TestUpdateTask(t *testing.T) {
	gin.SetMode(gin.TestMode)
	MockConnect := new(MockConnect)
	handler := httpy.HTTPHandler{Service: MockConnect}
	router := gin.New()
	router.PUT("/tasks", handler.UpdateTask)

	var update domain.Task
	task := domain.Task{
		ID: "1", Title: "test1", Description: "testing123",
	}

	MockConnect.On("UpdateTask", task).Return(nil)
	body, _ := json.Marshal(task)
	request, _ := http.NewRequest("PUT", "/tasks", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	code := httptest.NewRecorder()
	router.ServeHTTP(code, request)

	assert.Equal(t, http.StatusOK, code.Code)
	err := json.Unmarshal(code.Body.Bytes(), &update)
	assert.NoError(t, err)
	assert.Equal(t, update, task)
}

func TestDeleteTask(t *testing.T) {
	gin.SetMode(gin.TestMode)
	MockConnect := new(MockConnect)
	handler := httpy.HTTPHandler{Service: MockConnect}
	router := gin.New()
	router.DELETE("/tasks/:id", handler.Delete)

	MockConnect.On("Delete", "1").Return(nil)
	request, _ := http.NewRequest("DELETE", "/tasks/1", nil)
	request.Header.Set("Content-Type", "application/json")
	code := httptest.NewRecorder()
	router.ServeHTTP(code, request)

	assert.Equal(t, http.StatusOK, code.Code)
}

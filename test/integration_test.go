package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	httpy "task-api/internal/adapter/inbound/http"
	"task-api/internal/application"
	"task-api/internal/domain"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type InMemoryRepo struct {
	tasks map[string]domain.Task
}

func (r *InMemoryRepo) CreateTask(task domain.Task) (string, error) {
	id := "test-id"
	task.ID = id
	r.tasks[id] = task
	return id, nil
}
func (r *InMemoryRepo) GetByID(id string) (domain.Task, error) {
	task, ok := r.tasks[id]
	if !ok {
		return domain.Task{}, errors.New("not found")
	}
	return task, nil
}
func (r *InMemoryRepo) GetAll() ([]domain.Task, error) {
	var out []domain.Task
	for _, t := range r.tasks {
		out = append(out, t)
	}
	return out, nil
}
func (r *InMemoryRepo) UpdateTask(task domain.Task) error {
	r.tasks[task.ID] = task
	return nil
}
func (r *InMemoryRepo) Delete(id string) error {
	delete(r.tasks, id)
	return nil
}

func TestFullFlow_CreateGetDelete(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repo := &InMemoryRepo{tasks: make(map[string]domain.Task)}
	service := service.NewConnect(repo)
	handler := &httpy.HttpHandler{Service: service}
	r := gin.New()
	r.POST("/tasks", handler.CreateTask)
	r.GET("/tasks/:id", handler.GetByID)
	r.DELETE("/tasks/:id", handler.Delete)

	// Create
	task := domain.Task{Title: "Test", Description: "Desc"}
	body, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	// Get
	req, _ = http.NewRequest("GET", "/tasks/test-id", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Delete
	req, _ = http.NewRequest("DELETE", "/tasks/test-id", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

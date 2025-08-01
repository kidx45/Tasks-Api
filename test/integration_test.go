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


type TestingStruct struct {
    tasks map[string]domain.Task
}


func (r TestingStruct) CreateTask(task domain.Task) (string, error) {
    id := "1"
    task.ID = id
    r.tasks[id] = task
    return id, nil
}
func (r TestingStruct) GetByID(id string) (domain.Task, error) {
    task, ok := r.tasks[id]
    if !ok {
        return domain.Task{}, errors.New("not found")
    }
    return task, nil
}
func (r TestingStruct) GetAll() ([]domain.Task, error) {
    var out []domain.Task
    for _, t := range r.tasks {
        out = append(out, t)
    }
    return out, nil
}
func (r TestingStruct) UpdateTask(task domain.Task) error {
	task, ok := r.tasks[task.ID]
    if !ok {
        return errors.New("not found")
    }
    r.tasks[task.ID] = task
    return nil
}
func (r TestingStruct) Delete(id string) error {
	_, ok := r.tasks[id]
    if !ok {
        return errors.New("not found")
    }
    delete(r.tasks, id)
    return nil
}


func TestFullFlow(t *testing.T) {
    gin.SetMode(gin.TestMode)
    repo := TestingStruct{tasks: make(map[string]domain.Task)}
    service := service.NewConnect(repo)
    handler := httpy.HttpHandler{Service: service}
    r := gin.New()
    r.POST("/tasks", handler.CreateTask)
	r.GET("/tasks", handler.GetAll)
    r.GET("/tasks/:id", handler.GetByID)
	r.PUT("/tasks", handler.UpdateTask)
    r.DELETE("/tasks/:id", handler.Delete)

	var test domain.Task
	var tests []domain.Task
    task := domain.Task{Title: "Test1", Description: "Testing123"}
    body, _ := json.Marshal(task)
    request, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
    request.Header.Set("Content-Type", "application/json")
    code := httptest.NewRecorder()
    r.ServeHTTP(code, request)
    assert.Equal(t, http.StatusCreated, code.Code)
	err := json.Unmarshal(code.Body.Bytes(),&test)
	assert.NoError(t, err)
	assert.Equal(t,test.Title,task.Title)
	assert.Equal(t,test.Description,task.Description)

	request, _ = http.NewRequest("GET", "/tasks", nil)
    code = httptest.NewRecorder()
    r.ServeHTTP(code, request)
    assert.Equal(t, http.StatusOK, code.Code)

    request, _ = http.NewRequest("GET", "/tasks/2", nil)
    code = httptest.NewRecorder()
    r.ServeHTTP(code, request)
    assert.Equal(t, http.StatusNotFound, code.Code)

	request, _ = http.NewRequest("GET", "/tasks/1", nil)
    code = httptest.NewRecorder()
    r.ServeHTTP(code, request)
    assert.Equal(t, http.StatusOK, code.Code)
	err = json.Unmarshal(code.Body.Bytes(),&test)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, code.Code)

	task = domain.Task{ID: "2",Title: "Testnew",Description: "Testingnew"}
	body, _= json.Marshal(task)
	request, _ = http.NewRequest("PUT", "/tasks", bytes.NewBuffer(body))
    code = httptest.NewRecorder()
    r.ServeHTTP(code, request)
    assert.Equal(t, http.StatusInternalServerError, code.Code)

	task = domain.Task{ID: "1",Title: "Testnew",Description: "Testingnew"}
	body, _= json.Marshal(task)
	request, _ = http.NewRequest("PUT", "/tasks", bytes.NewBuffer(body))
    code = httptest.NewRecorder()
    r.ServeHTTP(code, request)
    assert.Equal(t, http.StatusOK, code.Code)
	request, _ = http.NewRequest("GET", "/tasks/1", nil)
    code = httptest.NewRecorder()
    r.ServeHTTP(code, request)
    assert.Equal(t, http.StatusOK, code.Code)
	err = json.Unmarshal(code.Body.Bytes(),&test)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, code.Code)

    request, _ = http.NewRequest("DELETE", "/tasks/2", nil)
    code = httptest.NewRecorder()
    r.ServeHTTP(code, request)
    assert.Equal(t, http.StatusInternalServerError, code.Code)

	request, _ = http.NewRequest("DELETE", "/tasks/1", nil)
    code = httptest.NewRecorder()
    r.ServeHTTP(code, request)
    assert.Equal(t, http.StatusOK, code.Code)
	request, _ = http.NewRequest("GET", "/tasks", nil)
    code = httptest.NewRecorder()
    r.ServeHTTP(code, request)
    assert.Equal(t, http.StatusOK, code.Code)
	err = json.Unmarshal(code.Body.Bytes(),&tests)
	assert.NoError(t,err)
	assert.Nil(t, tests)
}
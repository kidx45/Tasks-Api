// Package http Task API
//
// This is a simple Task API built with Gin and PostgreSQL.
//
//     Schemes: http
//     Host: localhost:8080
//     BasePath: /
//     Version: 1.0.0
//     Contact: Kidus Tessema<lumsk24@gmail.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta

package httpy

import (
	"log"
	"net/http"
	"task-api/internal/domain"
	"task-api/internal/port/inbound"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"
    
)

// Struct for connect to service in application
type httpHandler struct {
	service inbound.Connect
}

// Handler : For starting a server
func Handler(service inbound.Connect) {
	handler := &httpHandler{service: service}
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/tasks", handler.CreateTask)
	router.GET("/tasks", handler.GetAll)
	router.GET("/tasks/:id", handler.GetByID)
	router.PUT("/tasks", handler.UpdateTask)
	router.DELETE("/tasks/:id", handler.Delete)
	if err := router.Run("localhost:8080"); err != nil {
		log.Fatalf("Server can't start: %v", err)
	}

}

// @Summary Create a task
// @Description Create a task by accepting title and description from user
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body domain.Task true "Task to be created"
// @Success 201 {object} domain.Task "Created task"
// @Faliure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks [post]
func (h *httpHandler) CreateTask(c *gin.Context) {
	var task domain.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.service.CreateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Create_task task"})
		return
	}
	task.ID = id
	c.IndentedJSON(http.StatusCreated, task)
}

//@Summary Get all tasks
//@Description Retrive all tasks avaliable in the database
//@Tags tasks
//@Produce json
//@Success 200 {object} domain.Task "Tasks in database"
//@Failure 500 {object} map[string]string
//@Router /tasks [get]
func (h *httpHandler) GetAll(c *gin.Context) {
	tasks, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

//@Summary Get task by id
//@Description Retrive a task from the databse by accepting an id
//@Tags tasks
//@Accept json
//@Produce json
//@Param id path string true "Task ID"
//@Success 200 {object} domain.Task "Tasks in database by ID"
//@Failure 404 {object} map[string]string
//@Router /tasks/{id} [get]
func (h *httpHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	task, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

//@Summary Update task by ID
//@Description Update contents of a task whether the title or the description using the id
//@Tags tasks
//@Accept json
//@Produce json
//@Param task body domain.Task true "Task to update"
//@Success 200 {object} domain.Task "Tasks updated"
//@Failure 500 {object} map[string]string
//@Faliure 400 {object} map[string]string
//@Router /tasks [put]
func (h *httpHandler) UpdateTask(c *gin.Context) {
	var task domain.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.service.UpdateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to update the requested task"})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

//@Summary Delete task by ID
//@Description Delete a task for a database using the id
//@Tags tasks
//@Accept json
//@Produce json
//@Param task body domain.Task true "Task to delete"
//@Success 200 "OK"
//@Failure 500 {object} map[string]string
//@Router /tasks/{id} [delete]
func (h *httpHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete the requested task"})
		return
	}
	c.Status(http.StatusOK)
}

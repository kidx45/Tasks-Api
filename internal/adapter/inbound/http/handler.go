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
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// HTTPHandler : Struct for connect to service in application
type HTTPHandler struct {
	Service inbound.Connect
}

// Handler : For starting a server
func Handler(service inbound.Connect) {
	handler := HTTPHandler{Service: service}
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

// CreateTask : Function to create a task
// @Summary Create a task
// @Description Create a task by accepting title and description from user
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body domain.UserInput true "Task to be created"
// @Success 201 {object} domain.Task "Created task"
// @Faliure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks [post]
func (h HTTPHandler) CreateTask(c *gin.Context) {
	rqt := c.Request.Context()
	var task domain.Task
	var input domain.UserInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.Service.CreateTask(rqt, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Create task"})
		return
	}
	task.ID = id
	task.Title = input.Title
	task.Description = input.Description
	c.IndentedJSON(http.StatusCreated, task)
}

// GetAll : To get all tasks from the database
// @Summary Get all tasks
// @Description Retrive all tasks avaliable in the database
// @Tags tasks
// @Produce json
// @Success 200 {object} domain.Task "Tasks in database"
// @Failure 500 {object} map[string]string
// @Router /tasks [get]
func (h HTTPHandler) GetAll(c *gin.Context) {
	rqt := c.Request.Context()
	tasks, err := h.Service.GetAll(rqt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

// GetByID : To get tasks by id
// @Summary Get task by id
// @Description Retrive a task from the databse by accepting an id
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {object} domain.Task "Tasks in database by ID"
// @Failure 404 {object} map[string]string
// @Router /tasks/{id} [get]
func (h HTTPHandler) GetByID(c *gin.Context) {
	rqt := c.Request.Context()
	id := c.Param("id")
	task, err := h.Service.GetByID(rqt, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

// UpdateTask : To update task by Id
// @Summary Update task by ID
// @Description Update contents of a task whether the title or the description using the id
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body domain.Task true "Task to update"
// @Success 200 {object} domain.Task "Tasks updated"
// @Failure 500 {object} map[string]string
// @Faliure 400 {object} map[string]string
// @Router /tasks [put]
func (h HTTPHandler) UpdateTask(c *gin.Context) {
	rqt := c.Request.Context()
	var task domain.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.Service.UpdateTask(rqt, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to update the requested task"})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

// Delete : To Delete task by ID
// @Summary Delete task by ID
// @Description Delete a task for a database using the id
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "Task to delete"
// @Success 200 "OK"
// @Failure 500 {object} map[string]string
// @Router /tasks/{id} [delete]
func (h HTTPHandler) Delete(c *gin.Context) {
	rqt := c.Request.Context()
	id := c.Param("id")
	err := h.Service.Delete(rqt, id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete the requested task"})
		return
	}
	c.Status(http.StatusOK)
}

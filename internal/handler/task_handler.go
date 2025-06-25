package handler

import (
	"errors"
	"net/http"
	"taskctl/internal/model"
	"taskctl/internal/service"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskService service.TaskService
}

func NewTaskHandler(taskService service.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

var ErrCreateTask = errors.New("failed to create task")
var ErrTaskNotFound = errors.New("task not found")

func (t *TaskHandler) Create(c *gin.Context) {
	taskID, err := t.taskService.Create()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, ErrCreateTask)
		return
	}

	c.JSON(http.StatusCreated, map[string]model.TaskID{"task_id": taskID})
}

func (t *TaskHandler) Get(c *gin.Context) {
	taskID := model.TaskID(c.Param("task_id"))

	task, err := t.taskService.Get(taskID)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, ErrTaskNotFound)
		return
	}

	c.JSON(http.StatusOK, map[string]*model.Task{"data": task})
}

func (t *TaskHandler) Delete(c *gin.Context) {
	taskID := model.TaskID(c.Param("task_id"))

	if err := t.taskService.Delete(taskID); err != nil {
		c.AbortWithError(http.StatusNotFound, ErrTaskNotFound)
		return
	}

	c.AbortWithStatus(http.StatusOK)
}

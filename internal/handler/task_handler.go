package handler

import "taskctl/internal/service"

type TaskHandler struct {
	taskService service.TaskService
}

func NewTaskHandler(taskService service.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

func (t *TaskHandler) CreateTask() {
	
}
package service

import (
	"taskctl/internal/model"
	"taskctl/internal/repository"
	"time"
)

type TaskService interface {
	Create() (model.TaskID, error)
	Get(taskID model.TaskID) (*model.Task, error)
	Delete(taskID model.TaskID) error
}

type taskService struct {
	rep repository.TaskRepository
}

func NewTaskService(rep repository.TaskRepository) TaskService {
	return &taskService{
		rep: rep,
	}
}

func (t *taskService) Create() (model.TaskID, error) {
	task := model.Task{
		ID:        model.GenerateTaskID(),
		Status:    model.TaskStatusPending,
		CreatedAt: time.Now(),
	}

	if err := t.rep.CreateTask(task); err != nil {
		return "", err
	}

	go t.processTask(task)

	return task.ID, nil
}

func (t *taskService) Get(taskID model.TaskID) (*model.Task, error) {
	v, err := t.rep.GetTask(taskID)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (t *taskService) Delete(taskID model.TaskID) error {
	if err := t.rep.DeleteTask(taskID); err != nil {
		return err
	}

	return nil
}

func (t *taskService) processTask(task model.Task) {
	t.rep.UpdateTaskStatus(task.ID, string(model.TaskStatusInProgress))

	time.Sleep(10 * time.Second)

	updatedTask := task
	updatedTask.Status = model.TaskStatusCompleted
	updatedTask.Result = "created"
	updatedTask.Duration = time.Since(task.CreatedAt).Seconds()

	t.rep.UpdateTask(task.ID, updatedTask)
}

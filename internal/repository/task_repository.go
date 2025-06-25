package repository

import (
	"errors"
	"sync"
	"taskctl/internal/model"
	"taskctl/internal/storage"
)

type TaskRepository interface {
	CreateTask(task model.Task) error
	DeleteTask(taskID model.TaskID) error
	GetTask(taskID model.TaskID) (*model.Task, error)
	UpdateTaskStatus(taskID model.TaskID, status string)
	UpdateTask(taskID model.TaskID, task model.Task)
}

type taskRep struct {
	db *storage.Storage
	mu sync.Mutex
}

func NewTaskRepository(db *storage.Storage) TaskRepository {
	return &taskRep{db: db}
}

var ErrTaskAlreadyExists = errors.New("task already exists")

func (t *taskRep) CreateTask(task model.Task) error {
	if err := t.db.Set(task); err != nil {
		return err
	}

	return nil
}

func (t *taskRep) DeleteTask(taskID model.TaskID) error {
	if err := t.db.Delete(taskID); err != nil {
		return err
	}

	return nil
}

func (t *taskRep) GetTask(taskID model.TaskID) (*model.Task, error) {
	v, err := t.db.Get(taskID)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (t *taskRep) UpdateTaskStatus(taskID model.TaskID, status string) {
	t.db.UpdateTaskStatus(taskID, status)
}

func (t *taskRep) UpdateTask(taskID model.TaskID, task model.Task) {
	t.db.UpdateTask(taskID, task)
}
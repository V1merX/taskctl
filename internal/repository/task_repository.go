package repository

import (
	"errors"
	"sync"
	"taskctl/internal/model"
	"taskctl/internal/storage"
)

type TaskRepository interface {
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
	const op = "internal.repository.taskRep.CreateTask"

	if err := t.db.Set(task); err != nil {
		return ErrTaskAlreadyExists
	}

	return nil
}

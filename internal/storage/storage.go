package storage

import "taskctl/internal/model"

type StorageRepository interface {
}

type TaskID string

type Storage struct {
	data map[TaskID]model.Task
}

func NewStorage() StorageRepository {
	return &Storage{
		data: make(map[TaskID]model.Task, 0),
	}
}

package storage

import (
	"errors"
	"sync"
	"taskctl/internal/model"
)

type Storage struct {
	data map[model.TaskID]model.Task
	mu   sync.Mutex
}

func New() *Storage {
	return &Storage{
		data: make(map[model.TaskID]model.Task, 0),
	}
}

func (s *Storage) Set(task model.Task) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	if _, exists := s.data[task.ID]; exists {
		return errors.New("task already exists")
	}
	s.data[task.ID] = task

	return nil
}

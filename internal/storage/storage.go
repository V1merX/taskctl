package storage

import (
	"errors"
	"fmt"
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

var ErrTaskAlreadyExists = errors.New("task already exists")
var ErrTaskNotfound = errors.New("task not found")

func (s *Storage) Set(task model.Task) error {
	const op = "internal.storage.Storage.Set"

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.data[task.ID]; exists {
		return fmt.Errorf("%s: %w", op, ErrTaskAlreadyExists)
	}
	s.data[task.ID] = task

	return nil
}

func (s *Storage) Delete(taskID model.TaskID) error {
	const op = "internal.storage.Storage.Delete"

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.data[taskID]; !ok {
		return fmt.Errorf("%s: %w", op, ErrTaskNotfound)
	}

	delete(s.data, taskID)

	return nil
}

func (s *Storage) Get(taskID model.TaskID) (*model.Task, error) {
	const op = "internal.storage.Storage.Get"

	s.mu.Lock()
	defer s.mu.Unlock()

	v, ok := s.data[taskID]
	if !ok {
		return nil, fmt.Errorf("%s: %w", op, ErrTaskNotfound)
	}

	return &v, nil
}

func (s *Storage) UpdateTaskStatus(taskID model.TaskID, status string) {
	const op = "internal.storage.Storage.UpdateTaskStatus"

	s.mu.Lock()
	defer s.mu.Unlock()

	if task, exists := s.data[taskID]; exists {
		updatedTask := task
		updatedTask.Status = model.TaskStatus(status)
		s.data[taskID] = updatedTask
	}
}

func (s *Storage) UpdateTask(taskID model.TaskID, newTask model.Task) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.data[taskID]; exists {
		s.data[taskID] = newTask
	}
}
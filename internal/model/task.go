package model

import (
	"math/rand"
	"time"
)

type TaskStatus string
type TaskID string

const (
	TaskStatusPending    TaskStatus = "pending"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusCompleted  TaskStatus = "completed"
	TaskStatusFailed     TaskStatus = "failed"
)

type Task struct {
	ID        TaskID      `json:"id"`
	Status    TaskStatus  `json:"status"`
	Result    interface{} `json:"result"`
	Duration  float64     `json:"duration"`
	CreatedAt time.Time   `json:"created_at"`
}

var alph = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM0123456789"

func GenerateTaskID() TaskID {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	res := "task_"

	for i := 0; i < 9; i++ {
		res += string(alph[r.Intn(len(alph))])
	}

	return TaskID(res)
}

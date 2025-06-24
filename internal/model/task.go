package model

type TaskStatus string

const (
	TaskStatusPending    TaskStatus = "pending"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusCompleted  TaskStatus = "completed"
	TaskStatusFailed     TaskStatus = "failed"
)

type Task struct {
	ID        int64       `json:"id"`
	Status    TaskStatus  `json:"status"`
	Result    interface{} `json:"result"`
	Duration  int64       `json:"duration"`
	CreatedAt int64       `json:"created_at"`
}

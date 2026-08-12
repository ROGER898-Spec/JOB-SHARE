package domain

import (
	"context"
	"time"
)

type KanbanTask struct {
	ID          int       `json:"id"`
	JobID       int       `json:"job_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type KanbanTaskRepository interface {
	Create(ctx context.Context, task *KanbanTask) error
	FindByJobID(ctx context.Context, jobID int) ([]*KanbanTask, error)
	UpdateStatus(ctx context.Context, id int, status string) error
}

type KanbanTaskUsecase interface {
	CreateTask(ctx context.Context, task *KanbanTask) error
	GetTasksByJobID(ctx context.Context, jobID int) ([]*KanbanTask, error)
	UpdateTaskStatus(ctx context.Context, id int, status string) error
}

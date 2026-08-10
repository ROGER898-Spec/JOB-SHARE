package domain

import (
	"context"
	"time"
)

type Job struct {
	ID          string     `json:"id"`
	UmkmID      int        `json:"umkm_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Budget      float64    `json:"budget"`
	Status      string     `json:"status"` // open, in_progress, completed, closed
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

type JobRepository interface {
	Create(ctx context.Context, job *Job) error
	FindAll(ctx context.Context) ([]*Job, error)
	FindByID(ctx context.Context, id string) (*Job, error)
}

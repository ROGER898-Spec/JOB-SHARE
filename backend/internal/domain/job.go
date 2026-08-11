package domain

import (
	"context"
	"time"
)

type Job struct {
	ID           int        `json:"id"`
	UmkmID       int        `json:"umkm_id"`
	CategoryID   *int       `json:"category_id,omitempty"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	BudgetAmount float64    `json:"budget_amount"`
	StartDate    *time.Time `json:"start_date,omitempty"`
	EndDate      *time.Time `json:"end_date,omitempty"`
	Status       string     `json:"status"`
	CreatedAt    time.Time  `json:"created_at"`

	Skills []Skill `json:"skills,omitempty"`
}

type JobRepository interface {
	Create(ctx context.Context, job *Job) error
	AddJobSkills(ctx context.Context, jobID int, skillIDs []int) error
	FindAll(ctx context.Context) ([]*Job, error)
	FindByID(ctx context.Context, id int) (*Job, error)
}

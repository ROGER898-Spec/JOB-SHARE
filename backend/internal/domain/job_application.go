package domain

import (
	"context"
	"time"
)

type JobApplication struct {
	ID           int       `json:"id"`
	JobID        int       `json:"job_id"`
	FreelancerID int       `json:"freelancer_id"`
	PitchMessage string    `json:"pitch_message"`
	Status       string    `json:"status"` // pending, accepted, rejected, withdrawn
	AppliedAt    time.Time `json:"applied_at"`
}

type JobApplicationRepository interface {
	Create(ctx context.Context, app *JobApplication) error
	FindByJobID(ctx context.Context, jobID int) ([]*JobApplication, error)
	FindByFreelancerID(ctx context.Context, freelancerID int) ([]*JobApplication, error)
	UpdateStatus(ctx context.Context, id int, status string) error
}

type JobApplicationUsecase interface {
	ApplyJob(ctx context.Context, app *JobApplication) error
	GetByJobID(ctx context.Context, jobID int) ([]*JobApplication, error)
	GetByFreelancerID(ctx context.Context, freelancerID int) ([]*JobApplication, error)
	UpdateStatus(ctx context.Context, id int, status string) error
}

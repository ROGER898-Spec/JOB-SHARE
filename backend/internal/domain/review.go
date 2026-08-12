package domain

import (
	"context"
	"time"
)

type Review struct {
	ID           int       `json:"id"`
	JobID        int       `json:"job_id"`
	UmkmID       int       `json:"umkm_id"`
	FreelancerID int       `json:"freelancer_id"`
	Rating       int       `json:"rating"`
	Feedback     string    `json:"feedback"`
	CreatedAt    time.Time `json:"created_at"`
}

type ReviewRepository interface {
	Create(ctx context.Context, review *Review) error
	FindByJobID(ctx context.Context, jobID int) (*Review, error)
	FindByFreelancerID(ctx context.Context, freelancerID int) ([]*Review, error)
}

type ReviewUsecase interface {
	CreateReview(ctx context.Context, review *Review) error
	GetReviewByJobID(ctx context.Context, jobID int) (*Review, error)
	GetReviewsByFreelancerID(ctx context.Context, freelancerID int) ([]*Review, error)
}

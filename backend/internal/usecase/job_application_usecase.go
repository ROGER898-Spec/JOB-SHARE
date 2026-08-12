package usecase

import (
	"context"
	"time"

	"github.com/FyaEdu/JOB-SHARE/backend/internal/domain"
)

type jobApplicationUsecase struct {
	repo    domain.JobApplicationRepository
	timeout time.Duration
}

func NewJobApplicationUsecase(repo domain.JobApplicationRepository, t time.Duration) domain.JobApplicationUsecase {
	return &jobApplicationUsecase{repo, t}
}

func (u *jobApplicationUsecase) ApplyJob(ctx context.Context, app *domain.JobApplication) error {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()
	return u.repo.Create(ctx, app)
}

func (u *jobApplicationUsecase) GetByJobID(ctx context.Context, jobID int) ([]*domain.JobApplication, error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()
	return u.repo.FindByJobID(ctx, jobID)
}

func (u *jobApplicationUsecase) GetByFreelancerID(ctx context.Context, freelancerID int) ([]*domain.JobApplication, error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()
	return u.repo.FindByFreelancerID(ctx, freelancerID)
}

func (u *jobApplicationUsecase) UpdateStatus(ctx context.Context, id int, status string) error {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()
	return u.repo.UpdateStatus(ctx, id, status)
}

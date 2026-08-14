package usecase

import (
	"context"
	"time"

	"github.com/FyaEdu/JOB-SHARE/backend/internal/domain"
)

type JobUsecase interface {
	Create(ctx context.Context, job *domain.Job, skillIDs []int) (*domain.Job, error)
	GetAll(ctx context.Context) ([]*domain.Job, error)
	GetByID(ctx context.Context, id int) (*domain.Job, error)
	UpdateStatus(ctx context.Context, id int, status string) error
}

type jobUsecase struct {
	jobRepo        domain.JobRepository
	contextTimeout time.Duration
}

func NewJobUsecase(repo domain.JobRepository, timeout time.Duration) JobUsecase {
	return &jobUsecase{
		jobRepo:        repo,
		contextTimeout: timeout,
	}
}

func (u *jobUsecase) Create(ctx context.Context, job *domain.Job, skillIDs []int) (*domain.Job, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	err := u.jobRepo.Create(ctx, job)
	if err != nil {
		return nil, err
	}

	err = u.jobRepo.AddJobSkills(ctx, job.ID, skillIDs)
	if err != nil {
		return nil, err
	}

	return job, nil
}

func (u *jobUsecase) GetAll(ctx context.Context) ([]*domain.Job, error) {
	return u.jobRepo.FindAll(ctx)
}

func (u *jobUsecase) GetByID(ctx context.Context, id int) (*domain.Job, error) {
	return u.jobRepo.FindByID(ctx, id)
}

func (u *jobUsecase) UpdateStatus(ctx context.Context, id int, status string) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.jobRepo.UpdateStatus(ctx, id, status)
}
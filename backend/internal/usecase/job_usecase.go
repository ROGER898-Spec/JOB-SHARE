package usecase

import (
	"context"
	"time"

	"github.com/FyaEdu/JOB-SHARE/backend/internal/domain"
)

type JobUsecase interface {
	Create(ctx context.Context, umkmID int, title, description string, budget float64, status string) (*domain.Job, error)
	GetAll(ctx context.Context) ([]*domain.Job, error)
	GetByID(ctx context.Context, id string) (*domain.Job, error)
}

type jobUsecase struct {
	jobRepository  domain.JobRepository
	contextTimeout time.Duration
}

func NewJobUsecase(jobRepo domain.JobRepository, timeout time.Duration) JobUsecase {
	return &jobUsecase{
		jobRepository:  jobRepo,
		contextTimeout: timeout,
	}
}

func (u *jobUsecase) Create(ctx context.Context, umkmID int, title, description string, budget float64, status string) (*domain.Job, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	if status == "" {
		status = "open"
	}

	newJob := &domain.Job{
		UmkmID:      umkmID,
		Title:       title,
		Description: description,
		Budget:      budget,
		Status:      status,
	}

	err := u.jobRepository.Create(ctx, newJob)
	if err != nil {
		return nil, err
	}

	return newJob, nil
}

func (u *jobUsecase) GetAll(ctx context.Context) ([]*domain.Job, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	jobs, err := u.jobRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return jobs, nil
}

func (u *jobUsecase) GetByID(ctx context.Context, id string) (*domain.Job, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	job, err := u.jobRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return job, nil
}

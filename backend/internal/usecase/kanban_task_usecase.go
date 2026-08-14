package usecase

import (
	"context"
	"time"

	"github.com/FyaEdu/JOB-SHARE/backend/internal/domain"
)

type kanbanTaskUsecase struct {
	kanbanRepo     domain.KanbanTaskRepository
	contextTimeout time.Duration
}

func NewKanbanTaskUsecase(repo domain.KanbanTaskRepository, t time.Duration) domain.KanbanTaskUsecase {
	return &kanbanTaskUsecase{kanbanRepo: repo, contextTimeout: t}
}

func (u *kanbanTaskUsecase) CreateTask(ctx context.Context, task *domain.KanbanTask) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.kanbanRepo.Create(ctx, task)
}

func (u *kanbanTaskUsecase) GetTasksByJobID(ctx context.Context, jobID int) ([]*domain.KanbanTask, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.kanbanRepo.FindByJobID(ctx, jobID)
}

func (u *kanbanTaskUsecase) UpdateTaskStatus(ctx context.Context, id int, status string) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.kanbanRepo.UpdateStatus(ctx, id, status)
}

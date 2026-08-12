package usecase

import (
	"context"
	"time"

	"github.com/FyaEdu/JOB-SHARE/backend/internal/domain"
)

type CategoryUsecase interface {
	Create(ctx context.Context, name, description string) (*domain.ProjectCategory, error)
	GetAll(ctx context.Context) ([]*domain.ProjectCategory, error)
}

type categoryUsecase struct {
	categoryRepo   domain.CategoryRepository
	contextTimeout time.Duration
}

func NewCategoryUsecase(repo domain.CategoryRepository, timeout time.Duration) CategoryUsecase {
	return &categoryUsecase{
		categoryRepo:   repo,
		contextTimeout: timeout,
	}
}

func (u *categoryUsecase) Create(ctx context.Context, name, description string) (*domain.ProjectCategory, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	category := &domain.ProjectCategory{
		Name:        name,
		Description: description,
	}

	err := u.categoryRepo.Create(ctx, category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (u *categoryUsecase) GetAll(ctx context.Context) ([]*domain.ProjectCategory, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.categoryRepo.FindAll(ctx)
}

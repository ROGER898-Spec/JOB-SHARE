package usecase

import (
	"context"
	"time"

	"github.com/FyaEdu/JOB-SHARE/backend/internal/domain"
)

type SkillUsecase interface {
	Create(ctx context.Context, categoryID int, name string) (*domain.Skill, error)
	GetByCategory(ctx context.Context, categoryID int) ([]*domain.Skill, error)
}

type skillUsecase struct {
	skillRepo      domain.SkillRepository
	contextTimeout time.Duration
}

func NewSkillUsecase(repo domain.SkillRepository, timeout time.Duration) SkillUsecase {
	return &skillUsecase{
		skillRepo:      repo,
		contextTimeout: timeout,
	}
}

func (u *skillUsecase) Create(ctx context.Context, categoryID int, name string) (*domain.Skill, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	skill := &domain.Skill{
		CategoryID: categoryID,
		Name:       name,
	}

	err := u.skillRepo.Create(ctx, skill)
	if err != nil {
		return nil, err
	}

	return skill, nil
}

func (u *skillUsecase) GetByCategory(ctx context.Context, categoryID int) ([]*domain.Skill, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.skillRepo.FindByCategory(ctx, categoryID)
}

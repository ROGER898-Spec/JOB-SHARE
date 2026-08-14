package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/FyaEdu/JOB-SHARE/backend/internal/domain"
)

type UmkmProfileUsecase interface {
	CreateProfile(ctx context.Context, req *domain.UmkmProfile) (*domain.UmkmProfile, error)
	GetProfileByUserID(ctx context.Context, userID string) (*domain.UmkmProfile, error)
}

type umkmProfileUsecase struct {
	profileRepo    domain.UmkmProfileRepository
	contextTimeout time.Duration
}

func NewUmkmProfileUsecase(repo domain.UmkmProfileRepository, timeout time.Duration) UmkmProfileUsecase {
	return &umkmProfileUsecase{
		profileRepo:    repo,
		contextTimeout: timeout,
	}
}

func (u *umkmProfileUsecase) CreateProfile(ctx context.Context, req *domain.UmkmProfile) (*domain.UmkmProfile, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	existingProfile, err := u.profileRepo.FindByUserID(ctx, req.UserID)
	if err != nil {
		return nil, err
	}
	if existingProfile != nil {
		return nil, errors.New("profile for this user already exists")
	}

	err = u.profileRepo.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (u *umkmProfileUsecase) GetProfileByUserID(ctx context.Context, userID string) (*domain.UmkmProfile, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	profile, err := u.profileRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

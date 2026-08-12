package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/FyaEdu/JOB-SHARE/backend/internal/domain"
)

type FreelancerProfileUsecase interface {
	CreateProfile(ctx context.Context, req *domain.FreelancerProfile) (*domain.FreelancerProfile, error)
	GetProfileByUserID(ctx context.Context, userID string) (*domain.FreelancerProfile, error)
}

type freelancerProfileUsecase struct {
	profileRepo    domain.FreelancerProfileRepository
	contextTimeout time.Duration
}

func NewFreelancerProfileUsecase(repo domain.FreelancerProfileRepository, timeout time.Duration) FreelancerProfileUsecase {
	return &freelancerProfileUsecase{
		profileRepo:    repo,
		contextTimeout: timeout,
	}
}

func (u *freelancerProfileUsecase) CreateProfile(ctx context.Context, req *domain.FreelancerProfile) (*domain.FreelancerProfile, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	existingProfile, err := u.profileRepo.FindByUserID(ctx, req.UserID)
	if err != nil {
		return nil, err
	}
	if existingProfile != nil {
		return nil, errors.New("profile for this freelancer already exists")
	}

	err = u.profileRepo.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (u *freelancerProfileUsecase) GetProfileByUserID(ctx context.Context, userID string) (*domain.FreelancerProfile, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	profile, err := u.profileRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

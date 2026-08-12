package usecase

import (
	"context"
	"time"

	"github.com/FyaEdu/JOB-SHARE/backend/internal/domain"
)

type reviewUsecase struct {
	reviewRepo     domain.ReviewRepository
	contextTimeout time.Duration
}

func NewReviewUsecase(repo domain.ReviewRepository, t time.Duration) domain.ReviewUsecase {
	return &reviewUsecase{reviewRepo: repo, contextTimeout: t}
}

func (u *reviewUsecase) CreateReview(ctx context.Context, review *domain.Review) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.reviewRepo.Create(ctx, review)
}

func (u *reviewUsecase) GetReviewByJobID(ctx context.Context, jobID int) (*domain.Review, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.reviewRepo.FindByJobID(ctx, jobID)
}

func (u *reviewUsecase) GetReviewsByFreelancerID(ctx context.Context, freelancerID int) ([]*domain.Review, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.reviewRepo.FindByFreelancerID(ctx, freelancerID)
}

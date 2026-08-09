package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/FyaEdu/JOB-SHARE/backend/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Register(ctx context.Context, email, password, role string) (*domain.User, error)
	Login(ctx context.Context, email, password string) (*domain.User, error)
}

type authUsecase struct {
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

func NewAuthUsecase(userRepo domain.UserRepository, timeout time.Duration) AuthUsecase {
	return &authUsecase{
		userRepo:       userRepo,
		contextTimeout: timeout,
	}
}

func (u *authUsecase) Register(ctx context.Context, email, password, role string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	existingUser, err := u.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := &domain.User{
		Email:        email,
		PasswordHash: string(hashedPassword),
		Role:         role,
		IsActive:     true,
	}

	err = u.userRepo.Create(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (u *authUsecase) Login(ctx context.Context, email, password string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	user, err := u.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}

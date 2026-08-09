package domain

import (
	"context"
	"time"
)

type User struct {
	ID           string
	Email        string
	PasswordHash string
	Role         string
	IsActive     bool
	CreatedAt    time.Time
	UpdatedAt    *time.Time
}

type PaginationParam struct {
	Page   int
	Limit  int
	Offset int
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
}

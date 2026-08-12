package domain

import "context"

type ProjectCategory struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CategoryRepository interface {
	Create(ctx context.Context, category *ProjectCategory) error
	FindAll(ctx context.Context) ([]*ProjectCategory, error)
}

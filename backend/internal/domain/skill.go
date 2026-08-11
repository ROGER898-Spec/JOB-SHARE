package domain

import "context"

type Skill struct {
	ID         int    `json:"id"`
	CategoryID int    `json:"category_id"`
	Name       string `json:"name"`
}

type SkillRepository interface {
	Create(ctx context.Context, skill *Skill) error
	FindByCategory(ctx context.Context, categoryID int) ([]*Skill, error)
}

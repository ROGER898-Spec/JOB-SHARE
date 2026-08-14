package postgres

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"

	"github.com/FyaEdu/JOB-SHARE/backend/internal/domain"
)

type skillRepository struct {
	db *sql.DB
}

func NewSkillRepository(db *sql.DB) domain.SkillRepository {
	return &skillRepository{db: db}
}

func (r *skillRepository) Create(ctx context.Context, skill *domain.Skill) error {
	queryPath := filepath.Join("internal", "repository", "postgres", "queries", "skill_insert.sql")
	queryBytes, err := os.ReadFile(queryPath)
	if err != nil {
		return err
	}

	err = r.db.QueryRowContext(ctx, string(queryBytes), skill.CategoryID, skill.Name).Scan(&skill.ID)
	return err
}

func (r *skillRepository) FindByCategory(ctx context.Context, categoryID int) ([]*domain.Skill, error) {
	queryPath := filepath.Join("internal", "repository", "postgres", "queries", "skill_find_by_category.sql")
	queryBytes, err := os.ReadFile(queryPath)
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, string(queryBytes), categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var skills []*domain.Skill
	for rows.Next() {
		var s domain.Skill
		if err := rows.Scan(&s.ID, &s.CategoryID, &s.Name); err != nil {
			return nil, err
		}
		skills = append(skills, &s)
	}
	return skills, nil
}

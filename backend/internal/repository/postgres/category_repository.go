package postgres

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"

	"github.com/FyaEdu/JOB-SHARE/backend/internal/domain"
)

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) domain.CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(ctx context.Context, category *domain.ProjectCategory) error {
	queryPath := filepath.Join("internal", "repository", "postgres", "queries", "category_insert.sql")
	queryBytes, err := os.ReadFile(queryPath)
	if err != nil {
		return err
	}

	err = r.db.QueryRowContext(ctx, string(queryBytes), category.Name, category.Description).Scan(&category.ID)
	return err
}

func (r *categoryRepository) FindAll(ctx context.Context) ([]*domain.ProjectCategory, error) {
	queryPath := filepath.Join("internal", "repository", "postgres", "queries", "category_find_all.sql")
	queryBytes, err := os.ReadFile(queryPath)
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, string(queryBytes))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*domain.ProjectCategory
	for rows.Next() {
		var c domain.ProjectCategory
		if err := rows.Scan(&c.ID, &c.Name, &c.Description); err != nil {
			return nil, err
		}
		categories = append(categories, &c)
	}
	return categories, nil
}

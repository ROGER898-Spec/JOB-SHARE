package postgres

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"path/filepath"

	"github.com/FyaEdu/JOB-SHARE/backend/internal/domain"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	queryPath := filepath.Join("internal", "repository", "postgres", "queries", "user_insert.sql")
	queryBytes, err := os.ReadFile(queryPath)
	if err != nil {
		return err
	}

	query := string(queryBytes)

	err = r.db.QueryRowContext(
		ctx,
		query,
		user.Email,
		user.PasswordHash,
		user.Role,
		user.IsActive,
	).Scan(&user.ID, &user.Email, &user.Role, &user.IsActive, &user.CreatedAt, &user.UpdatedAt)

	return err
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	queryPath := filepath.Join("internal", "repository", "postgres", "queries", "user_select.sql")
	queryBytes, err := os.ReadFile(queryPath)
	if err != nil {
		return nil, err
	}

	query := string(queryBytes)

	user := &domain.User{}
	err = r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}

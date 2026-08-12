package postgres

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"path/filepath"

	"github.com/FyaEdu/JOB-SHARE/backend/internal/domain"
)

type freelancerProfileRepository struct {
	db *sql.DB
}

func NewFreelancerProfileRepository(db *sql.DB) domain.FreelancerProfileRepository {
	return &freelancerProfileRepository{db: db}
}

func (r *freelancerProfileRepository) Create(ctx context.Context, profile *domain.FreelancerProfile) error {
	queryPath := filepath.Join("internal", "repository", "postgres", "queries", "freelancer_profile_insert.sql")
	queryBytes, err := os.ReadFile(queryPath)
	if err != nil {
		return err
	}

	query := string(queryBytes)

	err = r.db.QueryRowContext(
		ctx,
		query,
		profile.UserID,
		profile.FullName,
		profile.BioSummary,
		profile.PhoneNumber,
		profile.City,
		profile.PortfolioLink,
		profile.CvURL,
	).Scan(
		&profile.ID,
		&profile.CreatedAt,
		&profile.UpdatedAt,
	)

	return err
}

func (r *freelancerProfileRepository) FindByUserID(ctx context.Context, userID string) (*domain.FreelancerProfile, error) {
	queryPath := filepath.Join("internal", "repository", "postgres", "queries", "freelancer_profile_select_by_user_id.sql")
	queryBytes, err := os.ReadFile(queryPath)
	if err != nil {
		return nil, err
	}

	query := string(queryBytes)

	profile := &domain.FreelancerProfile{}
	err = r.db.QueryRowContext(ctx, query, userID).Scan(
		&profile.ID,
		&profile.UserID,
		&profile.FullName,
		&profile.BioSummary,
		&profile.PhoneNumber,
		&profile.City,
		&profile.PortfolioLink,
		&profile.CvURL,
		&profile.CreatedAt,
		&profile.UpdatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return profile, nil
}

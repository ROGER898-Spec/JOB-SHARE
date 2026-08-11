package postgres

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"path/filepath"

	"github.com/FyaEdu/JOB-SHARE/backend/internal/domain"
)

type umkmProfileRepository struct {
	db *sql.DB
}

func NewUmkmProfileRepository(db *sql.DB) domain.UmkmProfileRepository {
	return &umkmProfileRepository{db: db}
}

func (r *umkmProfileRepository) Create(ctx context.Context, profile *domain.UmkmProfile) error {
	queryPath := filepath.Join("internal", "repository", "postgres", "queries", "umkm_profile_insert.sql")
	queryBytes, err := os.ReadFile(queryPath)
	if err != nil {
		return err
	}

	query := string(queryBytes)

	err = r.db.QueryRowContext(
		ctx,
		query,
		profile.UserID,
		profile.BusinessName,
		profile.OwnerName,
		profile.PhoneNumber,
		profile.City,
		profile.FullAddress,
	).Scan(
		&profile.ID,
		&profile.VerificationStatus,
		&profile.CreatedAt,
		&profile.UpdatedAt,
	)

	return err
}

func (r *umkmProfileRepository) FindByUserID(ctx context.Context, userID string) (*domain.UmkmProfile, error) {
	queryPath := filepath.Join("internal", "repository", "postgres", "queries", "umkm_profile_select_by_user_id.sql")
	queryBytes, err := os.ReadFile(queryPath)
	if err != nil {
		return nil, err
	}

	query := string(queryBytes)

	profile := &domain.UmkmProfile{}
	err = r.db.QueryRowContext(ctx, query, userID).Scan(
		&profile.ID,
		&profile.UserID,
		&profile.BusinessName,
		&profile.OwnerName,
		&profile.PhoneNumber,
		&profile.City,
		&profile.FullAddress,
		&profile.IdentityDocumentURL,
		&profile.VerificationStatus,
		&profile.VerifiedByAdminID,
		&profile.VerifiedAt,
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

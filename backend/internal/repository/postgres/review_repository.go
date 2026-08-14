package postgres

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"

	"github.com/FyaEdu/JOB-SHARE/backend/internal/domain"
)

type reviewRepository struct {
	db *sql.DB
}

func NewReviewRepository(db *sql.DB) domain.ReviewRepository {
	return &reviewRepository{db: db}
}

func (r *reviewRepository) Create(ctx context.Context, review *domain.Review) error {
	queryPath := filepath.Join("internal", "repository", "postgres", "queries", "review_insert.sql")
	queryBytes, err := os.ReadFile(queryPath)
	if err != nil {
		return err
	}

	err = r.db.QueryRowContext(
		ctx,
		string(queryBytes),
		review.JobID,
		review.UmkmID,
		review.FreelancerID,
		review.Rating,
		review.Feedback,
	).Scan(
		&review.ID,
		&review.CreatedAt,
	)

	return err
}

func (r *reviewRepository) FindByJobID(ctx context.Context, jobID int) (*domain.Review, error) {
	queryPath := filepath.Join("internal", "repository", "postgres", "queries", "review_select_by_job.sql")
	queryBytes, err := os.ReadFile(queryPath)
	if err != nil {
		return nil, err
	}

	var rev domain.Review
	err = r.db.QueryRowContext(ctx, string(queryBytes), jobID).Scan(
		&rev.ID, &rev.JobID, &rev.UmkmID, &rev.FreelancerID, &rev.Rating, &rev.Feedback, &rev.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &rev, nil
}

func (r *reviewRepository) FindByFreelancerID(ctx context.Context, freelancerID int) ([]*domain.Review, error) {
	queryPath := filepath.Join("internal", "repository", "postgres", "queries", "review_select_by_freelancer.sql")
	queryBytes, err := os.ReadFile(queryPath)
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, string(queryBytes), freelancerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []*domain.Review
	for rows.Next() {
		var rev domain.Review
		err := rows.Scan(
			&rev.ID, &rev.JobID, &rev.UmkmID, &rev.FreelancerID, &rev.Rating, &rev.Feedback, &rev.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, &rev)
	}

	return reviews, nil
}

package postgres

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"path/filepath"

	"github.com/FyaEdu/JOB-SHARE/backend/internal/domain"
)

type jobRepository struct {
	db *sql.DB
}

func NewJobRepository(db *sql.DB) domain.JobRepository {
	return &jobRepository{db: db}
}

func (r *jobRepository) Create(ctx context.Context, job *domain.Job) error {
	queryPath := filepath.Join("internal", "repository", "postgres", "queries", "job_create.sql")
	queryBytes, err := os.ReadFile(queryPath)
	if err != nil {
		return err
	}

	query := string(queryBytes)

	err = r.db.QueryRowContext(
		ctx,
		query,
		job.UmkmID,
		job.Title,
		job.Description,
		job.Budget,
		job.Status,
	).Scan(&job.ID, &job.CreatedAt, &job.UpdatedAt)

	return err
}

func (r *jobRepository) FindAll(ctx context.Context) ([]*domain.Job, error) {
	queryPath := filepath.Join("internal", "repository", "postgres", "queries", "job_find_all.sql")
	queryBytes, err := os.ReadFile(queryPath)
	if err != nil {
		return nil, err
	}

	query := string(queryBytes)

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobs []*domain.Job
	for rows.Next() {
		var j domain.Job
		err := rows.Scan(
			&j.ID,
			&j.UmkmID,
			&j.Title,
			&j.Description,
			&j.Budget,
			&j.Status,
			&j.CreatedAt,
			&j.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, &j)
	}

	return jobs, nil
}

func (r *jobRepository) FindByID(ctx context.Context, id string) (*domain.Job, error) {
	queryPath := filepath.Join("internal", "repository", "postgres", "queries", "job_find_by_id.sql")
	queryBytes, err := os.ReadFile(queryPath)
	if err != nil {
		return nil, err
	}

	query := string(queryBytes)

	row := r.db.QueryRowContext(ctx, query, id)
	var j domain.Job
	err = row.Scan(
		&j.ID,
		&j.UmkmID,
		&j.Title,
		&j.Description,
		&j.Budget,
		&j.Status,
		&j.CreatedAt,
		&j.UpdatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &j, nil
}

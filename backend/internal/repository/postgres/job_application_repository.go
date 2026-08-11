package postgres

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"

	"github.com/FyaEdu/JOB-SHARE/backend/internal/domain"
)

type jobApplicationRepository struct{ db *sql.DB }

func NewJobApplicationRepository(db *sql.DB) domain.JobApplicationRepository {
	return &jobApplicationRepository{db: db}
}

func (r *jobApplicationRepository) Create(ctx context.Context, app *domain.JobApplication) error {
	path := filepath.Join("internal", "repository", "postgres", "queries", "job_app_insert.sql")
	query, _ := os.ReadFile(path)
	return r.db.QueryRowContext(ctx, string(query), app.JobID, app.FreelancerID, app.PitchMessage).
		Scan(&app.ID, &app.AppliedAt)
}

func (r *jobApplicationRepository) FindByJobID(ctx context.Context, jobID int) ([]*domain.JobApplication, error) {
	path := filepath.Join("internal", "repository", "postgres", "queries", "job_app_select_by_job.sql")
	query, _ := os.ReadFile(path)
	rows, err := r.db.QueryContext(ctx, string(query), jobID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var apps []*domain.JobApplication
	for rows.Next() {
		var a domain.JobApplication
		rows.Scan(&a.ID, &a.JobID, &a.FreelancerID, &a.PitchMessage, &a.Status, &a.AppliedAt)
		apps = append(apps, &a)
	}
	return apps, nil
}

func (r *jobApplicationRepository) FindByFreelancerID(ctx context.Context, freelancerID int) ([]*domain.JobApplication, error) {
	path := filepath.Join("internal", "repository", "postgres", "queries", "job_app_select_by_freelancer.sql")
	query, _ := os.ReadFile(path)
	rows, err := r.db.QueryContext(ctx, string(query), freelancerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var apps []*domain.JobApplication
	for rows.Next() {
		var a domain.JobApplication
		rows.Scan(&a.ID, &a.JobID, &a.FreelancerID, &a.PitchMessage, &a.Status, &a.AppliedAt)
		apps = append(apps, &a)
	}
	return apps, nil
}

func (r *jobApplicationRepository) UpdateStatus(ctx context.Context, id int, status string) error {
	path := filepath.Join("internal", "repository", "postgres", "queries", "job_app_update_status.sql")
	query, _ := os.ReadFile(path)
	_, err := r.db.ExecContext(ctx, string(query), status, id)
	return err
}

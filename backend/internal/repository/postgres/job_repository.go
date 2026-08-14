package postgres

import (
	"context"
	"database/sql"
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
	queryPath := filepath.Join("internal", "repository", "postgres", "queries", "job_insert.sql")
	queryBytes, err := os.ReadFile(queryPath)
	if err != nil {
		return err
	}

	err = r.db.QueryRowContext(
		ctx,
		string(queryBytes),
		job.UmkmID,
		job.CategoryID,
		job.Title,
		job.Description,
		job.BudgetAmount,
		job.Location,
		job.RadiusKm,
		job.DurationLabel,
		job.StartDate,
		job.EndDate,
	).Scan(
		&job.ID,
		&job.Status,
		&job.CreatedAt,
	)

	return err
}

func (r *jobRepository) AddJobSkills(ctx context.Context, jobID int, skillIDs []int) error {
	queryPath := filepath.Join("internal", "repository", "postgres", "queries", "job_skill_insert.sql")
	queryBytes, err := os.ReadFile(queryPath)
	if err != nil {
		return err
	}
	query := string(queryBytes)

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, skillID := range skillIDs {
		_, err := tx.ExecContext(ctx, query, jobID, skillID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (r *jobRepository) FindAll(ctx context.Context) ([]*domain.Job, error) {
	queryPath := filepath.Join("internal", "repository", "postgres", "queries", "job_find_all.sql")
	queryBytes, err := os.ReadFile(queryPath)
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, string(queryBytes))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobs []*domain.Job
	for rows.Next() {
		var j domain.Job
		err := rows.Scan(
			&j.ID, &j.UmkmID, &j.CategoryID, &j.Title, &j.Description,
			&j.BudgetAmount, &j.Location, &j.RadiusKm, &j.DurationLabel,
			&j.StartDate, &j.EndDate, &j.Status, &j.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, &j)
	}
	return jobs, nil
}

func (r *jobRepository) UpdateStatus(ctx context.Context, id int, status string) error {
	queryPath := filepath.Join("internal", "repository", "postgres", "queries", "job_update_status.sql")
	queryBytes, err := os.ReadFile(queryPath)
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, string(queryBytes), status, id)
	return err
}

func (r *jobRepository) FindByID(ctx context.Context, id int) (*domain.Job, error) {
	queryPath := filepath.Join("internal", "repository", "postgres", "queries", "job_find_by_id.sql")
	queryBytes, err := os.ReadFile(queryPath)
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, string(queryBytes), id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var job *domain.Job
	for rows.Next() {
		if job == nil {
			job = &domain.Job{}
		}

		var skillID sql.NullInt32
		var skillCatID sql.NullInt32
		var skillName sql.NullString

		err := rows.Scan(
			&job.ID, &job.UmkmID, &job.CategoryID, &job.Title, &job.Description,
			&job.BudgetAmount, &job.Location, &job.RadiusKm, &job.DurationLabel,
			&job.StartDate, &job.EndDate, &job.Status, &job.CreatedAt,
			&skillID, &skillCatID, &skillName,
		)
		if err != nil {
			return nil, err
		}

		if skillID.Valid {
			job.Skills = append(job.Skills, domain.Skill{
				ID:         int(skillID.Int32),
				CategoryID: int(skillCatID.Int32),
				Name:       skillName.String,
			})
		}
	}

	if job == nil {
		return nil, sql.ErrNoRows
	}

	return job, nil
}
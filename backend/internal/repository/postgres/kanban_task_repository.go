package postgres

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"

	"github.com/FyaEdu/JOB-SHARE/backend/internal/domain"
)

type kanbanTaskRepository struct {
	db *sql.DB
}

func NewKanbanTaskRepository(db *sql.DB) domain.KanbanTaskRepository {
	return &kanbanTaskRepository{db: db}
}

func (r *kanbanTaskRepository) Create(ctx context.Context, task *domain.KanbanTask) error {
	queryPath := filepath.Join("internal", "repository", "postgres", "queries", "kanban_task_insert.sql")
	queryBytes, err := os.ReadFile(queryPath)
	if err != nil {
		return err
	}

	err = r.db.QueryRowContext(
		ctx,
		string(queryBytes),
		task.JobID,
		task.Title,
		task.Description,
	).Scan(
		&task.ID,
		&task.CreatedAt,
	)

	return err
}

func (r *kanbanTaskRepository) FindByJobID(ctx context.Context, jobID int) ([]*domain.KanbanTask, error) {
	queryPath := filepath.Join("internal", "repository", "postgres", "queries", "kanban_task_select_by_job.sql")
	queryBytes, err := os.ReadFile(queryPath)
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, string(queryBytes), jobID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*domain.KanbanTask
	for rows.Next() {
		var t domain.KanbanTask
		err := rows.Scan(
			&t.ID,
			&t.JobID,
			&t.Title,
			&t.Description,
			&t.Status,
			&t.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &t)
	}

	return tasks, nil
}

func (r *kanbanTaskRepository) UpdateStatus(ctx context.Context, id int, status string) error {
	queryPath := filepath.Join("internal", "repository", "postgres", "queries", "kanban_task_update_status.sql")
	queryBytes, err := os.ReadFile(queryPath)
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, string(queryBytes), status, id)
	return err
}

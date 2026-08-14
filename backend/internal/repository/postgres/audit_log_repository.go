package postgres

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"

	"github.com/FyaEdu/JOB-SHARE/backend/internal/domain"
)

type auditLogRepository struct {
	db *sql.DB
}

func NewAuditLogRepository(db *sql.DB) domain.AuditLogRepository {
	return &auditLogRepository{db: db}
}

func (r *auditLogRepository) Create(ctx context.Context, log *domain.AuditLog) error {
	queryPath := filepath.Join("internal", "repository", "postgres", "queries", "audit_log_insert.sql")
	queryBytes, err := os.ReadFile(queryPath)
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(
		ctx,
		string(queryBytes),
		log.UserID,
		log.Action,
		log.EntityType,
		log.EntityID,
		log.IPAddress,
		log.Details,
	)
	return err
}

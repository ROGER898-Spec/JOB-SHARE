package postgres

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"

	"github.com/FyaEdu/JOB-SHARE/backend/internal/domain"
)

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) domain.TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) Create(ctx context.Context, trx *domain.Transaction) error {
	queryPath := filepath.Join("internal", "repository", "postgres", "queries", "transaction_insert.sql")
	queryBytes, err := os.ReadFile(queryPath)
	if err != nil {
		return err
	}

	err = r.db.QueryRowContext(
		ctx,
		string(queryBytes),
		trx.JobID,
		trx.UmkmID,
		trx.FreelancerID,
		trx.Amount,
		trx.PaymentGatewayOrderID,
		trx.PaymentMethod,
	).Scan(
		&trx.ID,
		&trx.PaidAt,
	)

	return err
}

func (r *transactionRepository) FindByJobID(ctx context.Context, jobID int) (*domain.Transaction, error) {
	queryPath := filepath.Join("internal", "repository", "postgres", "queries", "transaction_select_by_job.sql")
	queryBytes, err := os.ReadFile(queryPath)
	if err != nil {
		return nil, err
	}

	var trx domain.Transaction
	var pgOrderID sql.NullString
	var pMethod sql.NullString

	err = r.db.QueryRowContext(ctx, string(queryBytes), jobID).Scan(
		&trx.ID, &trx.JobID, &trx.UmkmID, &trx.FreelancerID, &trx.Amount,
		&trx.EscrowStatus, &pgOrderID, &pMethod, &trx.PaidAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	trx.PaymentGatewayOrderID = pgOrderID.String
	trx.PaymentMethod = pMethod.String

	return &trx, nil
}

func (r *transactionRepository) FindByUserID(ctx context.Context, userID int, role string) ([]*domain.Transaction, error) {
	queryPath := filepath.Join("internal", "repository", "postgres", "queries", "transaction_select_by_user.sql")
	queryBytes, err := os.ReadFile(queryPath)
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, string(queryBytes), userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trxs []*domain.Transaction
	for rows.Next() {
		var trx domain.Transaction
		var pgOrderID sql.NullString
		var pMethod sql.NullString

		err := rows.Scan(
			&trx.ID, &trx.JobID, &trx.UmkmID, &trx.FreelancerID, &trx.Amount,
			&trx.EscrowStatus, &pgOrderID, &pMethod, &trx.PaidAt,
		)
		if err != nil {
			return nil, err
		}

		trx.PaymentGatewayOrderID = pgOrderID.String
		trx.PaymentMethod = pMethod.String
		trxs = append(trxs, &trx)
	}

	return trxs, nil
}

func (r *transactionRepository) UpdateEscrowStatus(ctx context.Context, id int, status string) error {
	queryPath := filepath.Join("internal", "repository", "postgres", "queries", "transaction_update_status.sql")
	queryBytes, err := os.ReadFile(queryPath)
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, string(queryBytes), status, id)
	return err
}

package domain

import (
	"context"
	"time"
)

type Transaction struct {
	ID                    int       `json:"id"`
	JobID                 int       `json:"job_id"`
	UmkmID                int       `json:"umkm_id"`
	FreelancerID          int       `json:"freelancer_id"`
	Amount                float64   `json:"amount"`
	EscrowStatus          string    `json:"escrow_status"`
	PaymentGatewayOrderID string    `json:"payment_gateway_order_id,omitempty"`
	PaymentMethod         string    `json:"payment_method,omitempty"`
	PaidAt                time.Time `json:"paid_at"`
}

type TransactionRepository interface {
	Create(ctx context.Context, trx *Transaction) error
	FindByJobID(ctx context.Context, jobID int) (*Transaction, error)
	FindByUserID(ctx context.Context, userID int, role string) ([]*Transaction, error)
	UpdateEscrowStatus(ctx context.Context, id int, status string) error
}

type TransactionUsecase interface {
	CreateTransaction(ctx context.Context, trx *Transaction) error
	GetTransactionByJobID(ctx context.Context, jobID int) (*Transaction, error)
	GetTransactionsByUser(ctx context.Context, userID int, role string) ([]*Transaction, error)
	ReleaseEscrow(ctx context.Context, id int) error
}

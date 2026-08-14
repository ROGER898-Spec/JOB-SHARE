package usecase

import (
	"context"
	"time"

	"github.com/FyaEdu/JOB-SHARE/backend/internal/domain"
)

type transactionUsecase struct {
	repo    domain.TransactionRepository
	timeout time.Duration
}

func NewTransactionUsecase(repo domain.TransactionRepository, t time.Duration) domain.TransactionUsecase {
	return &transactionUsecase{repo: repo, timeout: t}
}

func (u *transactionUsecase) CreateTransaction(ctx context.Context, trx *domain.Transaction) error {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	trx.EscrowStatus = "held"
	return u.repo.Create(ctx, trx)
}

func (u *transactionUsecase) GetTransactionByJobID(ctx context.Context, jobID int) (*domain.Transaction, error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()
	return u.repo.FindByJobID(ctx, jobID)
}

func (u *transactionUsecase) GetTransactionsByUser(ctx context.Context, userID int, role string) ([]*domain.Transaction, error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()
	return u.repo.FindByUserID(ctx, userID, role)
}

func (u *transactionUsecase) ReleaseEscrow(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()
	return u.repo.UpdateEscrowStatus(ctx, id, "released")
}

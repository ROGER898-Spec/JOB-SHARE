package handler

import (
	"strconv"

	dtoRequest "github.com/FyaEdu/JOB-SHARE/backend/internal/delivery/http/dto/request"
	dtoResponse "github.com/FyaEdu/JOB-SHARE/backend/internal/delivery/http/dto/response"
	"github.com/FyaEdu/JOB-SHARE/backend/internal/domain"
	pkgResponse "github.com/FyaEdu/JOB-SHARE/backend/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	trxUsecase domain.TransactionUsecase
}

func NewTransactionHandler(uc domain.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{trxUsecase: uc}
}

// Create godoc
// @Summary Create a Transaction (Escrow)
// @Description UMKM makes a payment for a job, funds are held in escrow.
// @Tags Transactions
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body request.CreateTransactionRequest true "Transaction Payload"
// @Success 201 {object} response.TransactionResponse
// @Router /transactions [post]
func (h *TransactionHandler) Create(c *fiber.Ctx) error {
	var req dtoRequest.CreateTransactionRequest
	if err := c.BodyParser(&req); err != nil {
		return pkgResponse.Error(c, fiber.StatusBadRequest, "Invalid request payload")
	}

	trx := &domain.Transaction{
		JobID:                 req.JobID,
		UmkmID:                req.UmkmID,
		FreelancerID:          req.FreelancerID,
		Amount:                req.Amount,
		PaymentGatewayOrderID: req.PaymentGatewayOrderID,
		PaymentMethod:         req.PaymentMethod,
	}

	if err := h.trxUsecase.CreateTransaction(c.Context(), trx); err != nil {
		return pkgResponse.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	res := dtoResponse.TransactionResponse{
		ID:                    trx.ID,
		JobID:                 trx.JobID,
		UmkmID:                trx.UmkmID,
		FreelancerID:          trx.FreelancerID,
		Amount:                trx.Amount,
		EscrowStatus:          trx.EscrowStatus,
		PaymentGatewayOrderID: trx.PaymentGatewayOrderID,
		PaymentMethod:         trx.PaymentMethod,
		PaidAt:                trx.PaidAt,
	}

	return pkgResponse.Success(c, fiber.StatusCreated, "Transaction created and funds held in escrow", res)
}

// GetByJobID godoc
// @Summary Get transaction by Job ID
// @Tags Transactions
// @Security BearerAuth
// @Param job_id path int true "Job ID"
// @Produce json
// @Success 200 {object} response.TransactionResponse
// @Router /transactions/job/{job_id} [get]
func (h *TransactionHandler) GetByJobID(c *fiber.Ctx) error {
	jobID, _ := strconv.Atoi(c.Params("job_id"))

	trx, err := h.trxUsecase.GetTransactionByJobID(c.Context(), jobID)
	if err != nil {
		return pkgResponse.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	if trx == nil {
		return pkgResponse.Error(c, fiber.StatusNotFound, "Transaction not found")
	}

	res := dtoResponse.TransactionResponse{
		ID: trx.ID, JobID: trx.JobID, UmkmID: trx.UmkmID, FreelancerID: trx.FreelancerID,
		Amount: trx.Amount, EscrowStatus: trx.EscrowStatus, PaidAt: trx.PaidAt,
		PaymentGatewayOrderID: trx.PaymentGatewayOrderID, PaymentMethod: trx.PaymentMethod,
	}

	return pkgResponse.Success(c, fiber.StatusOK, "Transaction retrieved", res)
}

// ReleaseEscrow godoc
// @Summary Release funds to Freelancer
// @Description UMKM releases the held funds when job is done.
// @Tags Transactions
// @Security BearerAuth
// @Param id path int true "Transaction ID"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /transactions/{id}/release [patch]
func (h *TransactionHandler) ReleaseEscrow(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	if err := h.trxUsecase.ReleaseEscrow(c.Context(), id); err != nil {
		return pkgResponse.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkgResponse.Success(c, fiber.StatusOK, "Escrow funds released successfully", nil)
}

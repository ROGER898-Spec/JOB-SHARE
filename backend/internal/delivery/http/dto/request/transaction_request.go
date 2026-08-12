package request

type CreateTransactionRequest struct {
	JobID                 int     `json:"job_id" validate:"required"`
	UmkmID                int     `json:"umkm_id" validate:"required"`
	FreelancerID          int     `json:"freelancer_id" validate:"required"`
	Amount                float64 `json:"amount" validate:"required"`
	PaymentGatewayOrderID string  `json:"payment_gateway_order_id"`
	PaymentMethod         string  `json:"payment_method"`
}

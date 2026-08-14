package response

import "time"

type TransactionResponse struct {
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

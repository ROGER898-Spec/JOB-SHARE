INSERT INTO transactions (job_id, umkm_id, freelancer_id, amount, payment_gateway_order_id, payment_method, escrow_status)
VALUES ($1, $2, $3, $4, $5, $6, 'held')
RETURNING id, paid_at;
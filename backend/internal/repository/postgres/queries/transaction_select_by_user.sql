SELECT id, job_id, umkm_id, freelancer_id, amount, escrow_status, payment_gateway_order_id, payment_method, paid_at
FROM transactions
WHERE (umkm_id = $1 OR freelancer_id = $1)
ORDER BY paid_at DESC;
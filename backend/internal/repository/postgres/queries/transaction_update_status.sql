UPDATE transactions
SET escrow_status = $1
WHERE id = $2;
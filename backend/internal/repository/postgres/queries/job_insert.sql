INSERT INTO jobs (umkm_id, category_id, title, description, budget_amount, start_date, end_date, status)
VALUES ($1, $2, $3, $4, $5, $6, $7, 'open')
RETURNING id, status, created_at;
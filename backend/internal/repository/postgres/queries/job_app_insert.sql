INSERT INTO jobs (umkm_id, category_id, title, description, budget_amount, location, radius_km, duration_label, start_date, end_date, status)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, 'open')
RETURNING id, status, created_at;
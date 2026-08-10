INSERT INTO jobs (umkm_id, title, description, budget, status)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, created_at, updated_at;
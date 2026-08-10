SELECT id, umkm_id, title, description, budget, status, created_at, updated_at 
FROM jobs 
WHERE id = $1;
SELECT id, job_id, umkm_id, freelancer_id, rating, feedback, created_at
FROM reviews
WHERE job_id = $1 LIMIT 1;
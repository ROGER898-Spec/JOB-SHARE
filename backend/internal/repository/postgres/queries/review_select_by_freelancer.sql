SELECT id, job_id, umkm_id, freelancer_id, rating, feedback, created_at
FROM reviews
WHERE freelancer_id = $1
ORDER BY created_at DESC;
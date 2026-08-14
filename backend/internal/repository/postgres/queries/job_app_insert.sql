INSERT INTO job_applications (job_id, freelancer_id, pitch_message, status)
VALUES ($1, $2, $3, 'pending')
RETURNING id, applied_at;
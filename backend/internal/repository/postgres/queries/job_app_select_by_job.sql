SELECT id, job_id, freelancer_id, pitch_message, status, applied_at 
FROM job_applications 
WHERE job_id = $1 ORDER BY applied_at DESC;
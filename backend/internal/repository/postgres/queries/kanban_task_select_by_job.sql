SELECT id, job_id, title, description, status, created_at
FROM kanban_tasks
WHERE job_id = $1
ORDER BY created_at ASC;
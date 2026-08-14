INSERT INTO kanban_tasks (job_id, title, description, status)
VALUES ($1, $2, $3, 'todo')
RETURNING id, created_at;
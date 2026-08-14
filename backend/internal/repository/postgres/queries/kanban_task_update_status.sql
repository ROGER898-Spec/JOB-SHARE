UPDATE kanban_tasks
SET status = $1
WHERE id = $2;
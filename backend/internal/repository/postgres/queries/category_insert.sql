INSERT INTO project_categories (name, description)
VALUES ($1, $2)
RETURNING id;
INSERT INTO skills (category_id, name)
VALUES ($1, $2)
RETURNING id;
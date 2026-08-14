SELECT id, category_id, name 
FROM skills 
WHERE category_id = $1 
ORDER BY name ASC;
INSERT INTO users (id, email, password_hash, role, is_active)
VALUES (gen_random_uuid(), $1, $2, $3, $4)
RETURNING id, email, role, is_active, created_at, updated_at;
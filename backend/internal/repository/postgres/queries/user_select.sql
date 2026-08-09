SELECT id, email, password_hash, role, is_active, created_at, updated_at
FROM users
WHERE email = $1;
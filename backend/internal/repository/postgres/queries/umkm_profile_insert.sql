INSERT INTO umkm_profiles (user_id, business_name, owner_name, phone_number, city, full_address, verification_status)
VALUES ($1, $2, $3, $4, $5, $6, 'pending')
RETURNING id, verification_status, created_at, updated_at;
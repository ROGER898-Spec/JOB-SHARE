SELECT id, user_id, business_name, owner_name, phone_number, city, full_address, identity_document_url, verification_status, verified_by_admin_id, verified_at, created_at, updated_at
FROM umkm_profiles
WHERE user_id = $1;
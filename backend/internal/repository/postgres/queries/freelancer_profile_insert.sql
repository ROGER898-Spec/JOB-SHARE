INSERT INTO freelancer_profiles (user_id, full_name, bio_summary, phone_number, city, portfolio_link, cv_url)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, created_at, updated_at;
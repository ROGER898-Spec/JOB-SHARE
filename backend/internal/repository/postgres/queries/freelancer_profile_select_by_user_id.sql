SELECT id, user_id, full_name, bio_summary, phone_number, city, portfolio_link, cv_url, created_at, updated_at
FROM freelancer_profiles
WHERE user_id = $1;
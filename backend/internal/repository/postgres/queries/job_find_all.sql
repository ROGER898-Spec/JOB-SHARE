SELECT id, umkm_id, category_id, title, description, budget_amount, location, radius_km, duration_label, start_date, end_date, status, created_at
FROM jobs
ORDER BY created_at DESC;
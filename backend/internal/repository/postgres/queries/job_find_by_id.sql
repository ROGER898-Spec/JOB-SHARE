SELECT 
    j.id, j.umkm_id, j.category_id, j.title, j.description, j.budget_amount,
    j.location, j.radius_km, j.duration_label, j.start_date, j.end_date, j.status, j.created_at,
    s.id AS skill_id, s.category_id AS skill_category_id, s.name AS skill_name
FROM jobs j
LEFT JOIN job_skills js ON j.id = js.job_id
LEFT JOIN skills s ON js.skill_id = s.id
WHERE j.id = $1;
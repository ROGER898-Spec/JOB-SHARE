CREATE TABLE jobs (
    id SERIAL PRIMARY KEY,
    umkm_id INT NOT NULL REFERENCES umkm_profiles(id) ON DELETE CASCADE,
    category_id INT REFERENCES project_categories(id) ON DELETE SET NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    budget_amount NUMERIC(15, 2) NOT NULL,
    start_date DATE,
    end_date DATE,
    status VARCHAR(50) DEFAULT 'open' CHECK (status IN ('open', 'in_progress', 'completed', 'cancelled')),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE job_skills (
    job_id INT NOT NULL REFERENCES jobs(id) ON DELETE CASCADE,
    skill_id INT NOT NULL REFERENCES skills(id) ON DELETE CASCADE,
    PRIMARY KEY (job_id, skill_id)
);
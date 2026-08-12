CREATE TABLE reviews (
    id SERIAL PRIMARY KEY,
    job_id INT NOT NULL REFERENCES jobs(id) ON DELETE CASCADE,
    umkm_id INT NOT NULL REFERENCES umkm_profiles(id) ON DELETE CASCADE,
    freelancer_id INT NOT NULL REFERENCES freelancer_profiles(id) ON DELETE CASCADE,
    rating INT NOT NULL CHECK (rating >= 1 AND rating <= 5),
    feedback TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(job_id) -- Mencegah UMKM memberi lebih dari 1 review untuk job yang sama
);
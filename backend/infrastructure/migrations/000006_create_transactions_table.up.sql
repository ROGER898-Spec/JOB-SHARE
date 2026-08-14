CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    job_id INT NOT NULL REFERENCES jobs(id) ON DELETE CASCADE,
    umkm_id INT NOT NULL REFERENCES umkm_profiles(id) ON DELETE CASCADE,
    freelancer_id INT NOT NULL REFERENCES freelancer_profiles(id) ON DELETE CASCADE,
    amount NUMERIC(15, 2) NOT NULL,
    escrow_status VARCHAR(50) DEFAULT 'held' CHECK (escrow_status IN ('held', 'released', 'refunded')),
    payment_gateway_order_id VARCHAR(255),
    payment_method VARCHAR(100),
    paid_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
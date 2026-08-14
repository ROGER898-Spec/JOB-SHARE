ALTER TABLE jobs
    DROP COLUMN IF EXISTS location,
    DROP COLUMN IF EXISTS radius_km,
    DROP COLUMN IF EXISTS duration_label;
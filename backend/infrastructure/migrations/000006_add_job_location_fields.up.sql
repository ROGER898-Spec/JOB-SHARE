-- Kolom ini dibutuhkan supaya fitur "hyperlocal matching" & mode SOS bisa jalan --
-- frontend (posting.html) sudah punya form untuk field-field ini, tapi tabel jobs
-- sebelumnya belum punya tempat buat nyimpennya.
ALTER TABLE jobs
    ADD COLUMN location VARCHAR(255),
    ADD COLUMN radius_km INT,
    ADD COLUMN duration_label VARCHAR(50);
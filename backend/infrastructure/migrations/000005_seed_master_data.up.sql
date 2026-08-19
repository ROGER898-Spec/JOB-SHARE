-- Seed kategori sesuai daftar yang sudah dipakai di frontend (posting.html / register.html)
-- supaya dropdown kategori & skill tidak kosong begitu aplikasi pertama kali dijalankan.
-- Setiap kategori diberi 1 skill default dengan nama yang sama, supaya field skill_ids
-- (wajib diisi saat membuat job) bisa langsung dipakai tanpa perlu UI pemilihan skill dulu.

INSERT INTO project_categories (name, description) VALUES
    ('Fotografer Produk', 'Pemotretan produk untuk katalog atau promosi'),
    ('Videografer', 'Pengambilan dan penyuntingan video'),
    ('Admin Media Sosial', 'Pengelolaan konten dan interaksi media sosial'),
    ('Bantuan Event / Bazar', 'Tenaga bantu untuk acara atau bazar'),
    ('Stok Opname', 'Penghitungan dan pencatatan stok barang'),
    ('Packing & Gudang', 'Pengemasan dan pengelolaan gudang'),
    ('Kasir Pengganti', 'Pengganti sementara untuk posisi kasir'),
    ('Desainer Konten', 'Pembuatan desain visual untuk konten promosi'),
    ('Lainnya', 'Kebutuhan di luar kategori yang tersedia');

INSERT INTO skills (category_id, name)
SELECT id, name FROM project_categories;
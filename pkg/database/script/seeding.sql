
INSERT INTO users (user_id, username, email, password_hash, profile_picture, registration_date, role, created_at, updated_at)
VALUES
('USR-001', 'admin', 'admin@library.test', '$2b$12$adminhash', decode('iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mP8/x8AAwMCAO5XWfQAAAAASUVORK5CYII=', 'base64'), now(), 'admin', now(), now()),
('USR-002', 'budi',  'budi@library.test',  '$2b$12$userhash1', decode('iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mP8/x8AAwMCAO5XWfQAAAAASUVORK5CYII=', 'base64'), now(), 'user',  now(), now()),
('USR-003', 'sari',  'sari@library.test',  '$2b$12$userhash2', decode('iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mP8/x8AAwMCAO5XWfQAAAAASUVORK5CYII=', 'base64'), now(), 'user',  now(), now());

INSERT INTO categories (category_id, name, description, created_at, updated_at)
VALUES
('CAT-001', 'Programming', 'Buku tentang pemrograman dan teknologi', now(), now()),
('CAT-002', 'Productivity', 'Buku pengembangan diri dan produktivitas', now(), now()),
('CAT-003', 'Fantasy', 'Buku fiksi fantasi', now(), now());

INSERT INTO books (book_id, author, book_picture, title, description, status, created_at, updated_at)
VALUES
('BOK-001', 'Robert C. Martin', decode('iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mP8/x8AAwMCAO5XWfQAAAAASUVORK5CYII=', 'base64'), 'Clean Code', 'Panduan menulis kode yang rapi dan mudah dipelihara', 'borrowed', now(), now()),
('BOK-002', 'James Clear',      decode('iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mP8/x8AAwMCAO5XWfQAAAAASUVORK5CYII=', 'base64'), 'Atomic Habits', 'Buku tentang kebiasaan kecil yang berdampak besar', 'available', now(), now()),
('BOK-003', 'Frank Herbert',    decode('iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mP8/x8AAwMCAO5XWfQAAAAASUVORK5CYII=', 'base64'), 'Dune', 'Novel fiksi ilmiah klasik', 'available', now(), now());

INSERT INTO book_categories (book_category_id, book_id, category_id, created_at)
VALUES
('BCA-001', 'BOK-001', 'CAT-001', now()),
('BCA-002', 'BOK-002', 'CAT-002', now()),
('BCA-003', 'BOK-003', 'CAT-003', now()),
('BCA-004', 'BOK-003', 'CAT-001', now());

INSERT INTO book_borrowings (borrowing_id, user_id, book_id, borrowed_at, returned_at)
VALUES
('BOR-001', 'USR-002', 'BOK-001', now() - interval '3 day', NULL),
('BOR-002', 'USR-003', 'BOK-002', now() - interval '10 day', now() - interval '2 day');
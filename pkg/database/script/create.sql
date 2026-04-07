CREATE DATABASE library_db;

CREATE TABLE users (
    user_id           TEXT PRIMARY KEY,
    username          TEXT NOT NULL UNIQUE,
    email             TEXT NOT NULL UNIQUE,
    password_hash     TEXT NOT NULL,
    profile_picture   BYTEA,
    registration_date TIMESTAMPTZ NOT NULL DEFAULT now(),
    role              TEXT NOT NULL DEFAULT 'user',
    created_at        TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at        TIMESTAMPTZ
);

CREATE TABLE categories (
    category_id TEXT PRIMARY KEY,
    name        TEXT NOT NULL UNIQUE,
    description TEXT,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at  TIMESTAMPTZ
);

CREATE TABLE books (
    book_id      TEXT PRIMARY KEY,
    author       TEXT NOT NULL,
    book_picture BYTEA,
    title        TEXT NOT NULL,
    description  TEXT,
    status       TEXT NOT NULL DEFAULT 'available',
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at   TIMESTAMPTZ
);

CREATE TABLE book_categories (
    book_category_id TEXT PRIMARY KEY,
    book_id          TEXT NOT NULL REFERENCES books(book_id) ON DELETE CASCADE,
    category_id      TEXT NOT NULL REFERENCES categories(category_id) ON DELETE CASCADE,
    created_at       TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (book_id, category_id)
);

CREATE TABLE book_borrowings (
    borrowing_id TEXT PRIMARY KEY,
    user_id      TEXT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    book_id      TEXT NOT NULL REFERENCES books(book_id) ON DELETE CASCADE,
    borrowed_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    returned_at  TIMESTAMPTZ
);

CREATE INDEX idx_books_title ON books(title);
CREATE INDEX idx_books_status ON books(status);
CREATE INDEX idx_book_categories_book_id ON book_categories(book_id);
CREATE INDEX idx_book_categories_category_id ON book_categories(category_id);
CREATE INDEX idx_book_borrowings_user_id ON book_borrowings(user_id);
CREATE INDEX idx_book_borrowings_book_id ON book_borrowings(book_id);
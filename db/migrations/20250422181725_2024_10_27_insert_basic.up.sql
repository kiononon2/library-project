-- Users
INSERT INTO users (id, name, email, password, role, created_at, updated_at)
VALUES
    (1, 'Alice Jackson', 'aliceJ@example.com', 'hashed_password_1', 'user', now(), now()),
    (2, 'Bob Johnson', 'bobJ@example.com', 'hashed_password_2', 'admin', now(), now());

-- Authors
INSERT INTO authors (id, name, bio, photo_url, created_at, updated_at)
VALUES
    (1, 'George Orwell', 'Author of 1984 and Animal Farm', 'https://example.com/orwell.jpg', now(), now()),
    (2, 'J.K. Rowling', 'Author of the Harry Potter series', 'https://example.com/rowling.jpg', now(), now());

-- Categories
INSERT INTO categories (id, name, created_at, updated_at)
VALUES
    (1, 'Dystopian', now(), now()),
    (2, 'Fantasy', now(), now());

-- Books
INSERT INTO books (id, title, description, isbn, year, cover_url, status, author_id, category_id, created_at, updated_at)
VALUES
    (1, '1984', 'Dystopian novel', '1234567890', 1949, 'https://example.com/1984.jpg', 'available', 1, 1, now(), now()),
    (2, 'Harry Potter and the Sorcerer''s Stone', 'Fantasy novel', '0987654321', 1997, 'https://example.com/hp1.jpg', 'available', 2, 2, now(), now());

-- Loans
INSERT INTO loans (id, user_id, book_id, borrowed_at, due_date, returned_at, created_at, updated_at)
VALUES
    (1, 1, 1, now(), now() + interval '14 days', NULL, now(), now()),
    (2, 2, 2, now(), now() + interval '14 days', NULL, now(), now());

-- Reviews
INSERT INTO reviews (id, user_id, book_id, rating, comment, created_at, updated_at)
VALUES
    (1, 1, 1, 5, 'Amazing read!', now(), now()),
    (2, 2, 2, 4, 'Great book, highly recommended.', now(), now());

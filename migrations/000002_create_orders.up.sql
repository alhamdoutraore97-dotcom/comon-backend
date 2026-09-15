CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    product TEXT NOT NULL,
    quantity INT NOT NULL
);
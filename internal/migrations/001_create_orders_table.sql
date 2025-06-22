CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    customer_name VARCHAR(100),
    total NUMERIC(10,2),
    created_at TIMESTAMP DEFAULT NOW()
);

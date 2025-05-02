CREATE TABLE products (
                          id SERIAL PRIMARY KEY,
                          name TEXT NOT NULL,
                          price NUMERIC NOT NULL,
                          description TEXT,
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

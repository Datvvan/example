CREATE TABLE modules (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    path text,
    type VARCHAR,
    updated_at TIMESTAMP DEFAULT NOW(),
    created_at TIMESTAMP DEFAULT NOW()
);
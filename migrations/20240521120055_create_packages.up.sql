CREATE TABLE IF NOT EXISTS packages(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    act VARCHAR(50) NOT NULL,
    price BIGINT,
    valid_months BIGINT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);
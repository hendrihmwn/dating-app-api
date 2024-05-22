CREATE TABLE IF NOT EXISTS orders(
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    package_name VARCHAR(50) NOT NULL,
    package_price BIGINT NOT NULL,
    status SMALLINT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP,

    CONSTRAINT orders_user_id_fk FOREIGN KEY (user_id) REFERENCES users(id)
);
COMMENT ON COLUMN orders.status IS '0: INITIAL_ORDER, 1: FINISH'
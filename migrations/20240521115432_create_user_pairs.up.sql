CREATE TABLE IF NOT EXISTS user_pairs(
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    pair_user_id BIGINT NOT NULL,
    status SMALLINT DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP,

    CONSTRAINT user_pairs_user_id_fk FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT user_pairs_pair_user_id_fk FOREIGN KEY (pair_user_id) REFERENCES users(id)
);
COMMENT ON COLUMN user_pairs.status IS '0: SWIPE, 1: LIKE'
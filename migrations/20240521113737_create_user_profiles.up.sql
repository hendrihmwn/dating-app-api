CREATE TABLE IF NOT EXISTS user_profiles(
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    image VARCHAR(100),
    birthdate DATE,
    gender SMALLINT DEFAULT 0,
    location VARCHAR(50),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP,

    CONSTRAINT user_profiles_user_id_fk FOREIGN KEY (user_id) REFERENCES users(id)
);
COMMENT ON COLUMN user_profiles.gender IS '0: MEN, 1: WOMEN'
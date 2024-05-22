CREATE TABLE IF NOT EXISTS user_packages(
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    package_id BIGINT NOT NULL,
    package_act VARCHAR(50) NOT NULL,
    expired_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP,

    CONSTRAINT user_packages_user_id_fk FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT user_packages_package_id_fk FOREIGN KEY (package_id) REFERENCES packages(id)
);
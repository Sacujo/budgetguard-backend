-- +goose Up
CREATE EXTENSION IF NOT EXISTS citext;
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), 
    email CITEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    revision_ctr BIGINT NOT NULL DEFAULT 0,
    created_at	TIMESTAMPTZ	NOT NULL DEFAULT now(),
    updated_at	TIMESTAMPTZ	NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE users;

CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE movies (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    title TEXT NOT NULL UNIQUE,
    description TEXT,
    duration_mins INT NOT NULL CHECK (duration_mins > 0),

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_movies_title ON movies(title);
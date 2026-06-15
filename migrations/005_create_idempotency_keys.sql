CREATE TABLE idempotency_keys (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    idempotency_key TEXT NOT NULL UNIQUE,

    response JSONB,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_idempotency_key
ON idempotency_keys(idempotency_key);
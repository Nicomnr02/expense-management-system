CREATE TABLE IF NOT EXISTS payments (
    id BIGSERIAL PRIMARY KEY,
    external_id UUID NOT NULL REFERENCES expenses(id),
    status VARCHAR NOT NULL DEFAULT 'Pending',
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);
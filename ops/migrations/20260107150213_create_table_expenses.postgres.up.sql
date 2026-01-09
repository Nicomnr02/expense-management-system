CREATE TABLE IF NOT EXISTS expenses (
    id UUID PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id),
    amount BIGINT NOT NULL,
    description TEXT NOT NULL,
    status expense_status NOT NULL,
    receipt_url TEXT NOT NULL,
    submitted_at TIMESTAMPTZ NOT NULL,
    processed_at TIMESTAMPTZ NOT NULL
);
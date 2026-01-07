CREATE TABLE IF NOT EXISTS approvals (
    id BIGSERIAL PRIMARY KEY,
    expense_id UUID NOT NULL REFERENCES expenses(id),
    approver_id BIGINT NOT NULL REFERENCES users(id),
    status expense_status NOT NULL,
    notes TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL
);
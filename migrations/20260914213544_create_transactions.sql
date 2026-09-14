-- +goose Up
CREATE TABLE transactions (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    account_id UUID NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    category_id UUID NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
    amount NUMERIC(19, 4) NOT NULL,
    transaction_date TIMESTAMPTZ NOT NULL,
    comment TEXT,
    revision BIGINT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ
);

CREATE INDEX idx_transactions_user_revision ON transactions (user_id, revision);
CREATE INDEX idx_transactions_account_date ON transactions (user_id, account_id, transaction_date);

-- +goose Down
DROP TABLE transactions;

-- +goose Up
-- +goose StatementBegin
CREATE TABLE subscriptions (
    id SERIAL PRIMARY KEY,
    service_name TEXT,
    price INTEGER,
    user_id UUID,
    start_date DATE,
    end_date DATE
);

CREATE INDEX idx_service_name ON subscriptions(service_name);

CREATE INDEX idx_subscription_dates ON subscriptions(start_date, end_date);

CREATE INDEX idx_user_id ON subscriptions(user_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS subscriptions

-- +goose StatementEnd

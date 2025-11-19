-- +goose Up
AlTER TABLE feeds ADD COLUMN last_fetched_at TIMESTAMPTZ;

-- +goose Down
AlTER TABLE feeds DROP COLUMN last_fetched_at;
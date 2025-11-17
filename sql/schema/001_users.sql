-- +goose Up

CREATE TABLE users(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),  
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    Name TEXT NOT NULL

);
-- +goose Down
DROP TABLE users;
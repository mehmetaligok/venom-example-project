-- +migrate Up
CREATE TABLE IF NOT EXISTS users
(
    "id"         uuid PRIMARY KEY,
    "first_name" varchar     NOT NULL,
    "last_name"  varchar     NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT now()
);

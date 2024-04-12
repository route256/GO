-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE note
(
    id         BIGSERIAL PRIMARY KEY,
    user_id    BIGSERIAL NOT NULL,
    title      TEXT      NOT NULL,
    content    TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE note;

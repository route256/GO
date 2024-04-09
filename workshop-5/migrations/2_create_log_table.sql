-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE logs
(
    id         BIGSERIAL PRIMARY KEY,
    note_id    BIGSERIAL NOT NULL,
    content    TEXT      NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROPT TABLE logs;

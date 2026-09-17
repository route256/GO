-- +goose Up
-- +goose StatementBegin
CREATE TABLE animals (
    id BIGSERIAL PRIMARY KEY,
    nickname TEXT NOT NULL,
    birthday TIMESTAMP NOT NULL,
    weight INTEGER NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE animals;
-- +goose StatementEnd

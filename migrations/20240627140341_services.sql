-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS services
(
    name     TEXT PRIMARY KEY,
    image    TEXT,
    state    INT DEFAULT 0,
    replicas INT DEFAULT 1
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS services;
-- +goose StatementEnd

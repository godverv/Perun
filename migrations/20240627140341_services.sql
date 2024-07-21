-- +goose Up
-- +goose StatementBegin

-- Table "services" contains information about services registered in system
-- this is virtual services information.
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

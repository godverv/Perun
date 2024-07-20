-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS resources
(
    name         TEXT PRIMARY KEY,
    service_name TEXT REFERENCES services (name),
    image        TEXT,
    state        INT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE resources;
-- +goose StatementEnd

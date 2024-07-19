-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS resources (
    resource_name TEXT,
    service_name TEXT REFERENCES services(name)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE resources;
-- +goose StatementEnd

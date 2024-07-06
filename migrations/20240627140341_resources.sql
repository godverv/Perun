-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS resources
(
    resource_full_name TEXT PRIMARY KEY,
    node_name          TEXT REFERENCES nodes (node_name),
    state              INT DEFAULT 0,
    port               INT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS resources;
-- +goose StatementEnd

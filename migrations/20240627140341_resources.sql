-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS resources
(
    resource_full_name TEXT,
    node_name          TEXT REFERENCES nodes (node_name),
    state              INT DEFAULT 0,
    port               INT,

    PRIMARY KEY (resource_full_name, node_name)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS resources;
-- +goose StatementEnd

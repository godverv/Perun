-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS instances (
    service_name TEXT,
    node_name    TEXT REFERENCES nodes(name),
    addr         TEXT,
    state        INT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS instances;
-- +goose StatementEnd

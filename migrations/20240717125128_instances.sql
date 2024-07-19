-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS instances (
    service_name TEXT REFERENCES services(name),
    node_name    TEXT REFERENCES nodes(name),
    addr         TEXT,
    state        INT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin

-- Table "nodes" contains information about working nodes - connections|ssh|settings
CREATE TABLE IF NOT EXISTS nodes
(
    name                  TEXT PRIMARY KEY,
    addr                  TEXT,

    velez_port            INT,
    custom_velez_key_path TEXT,
    is_insecure           BOOLEAN,

    ssh_key               TEXT,
    ssh_port              INT,
    ssh_user_name         TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS nodes;
-- +goose StatementEnd
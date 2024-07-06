-- +goose Up
CREATE TABLE IF NOT EXISTS nodes
(
    node_name             TEXT PRIMARY KEY,
    addr                  TEXT,
    velez_port            INT,
    custom_velez_key_path TEXT,
    is_insecure           BOOLEAN,

    ssh_key               TEXT,
    ssh_port              INT,
    ssh_user_name         TEXT
);
-- +goose Down

DROP TABLE IF EXISTS nodes;
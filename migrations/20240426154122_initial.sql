-- +goose Up
CREATE TABLE IF NOT EXISTS nodes
(
    node_name             TEXT PRIMARY KEY,
    ssh_key               TEXT,
    ssh_addr              TEXT,
    ssh_user_name         TEXT,
    velez_addr            TEXT,
    custom_velez_key_path TEXT,
    insecure              BOOLEAN
);
-- +goose Down

DROP TABLE IF EXISTS nodes;
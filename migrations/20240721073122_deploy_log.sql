-- +goose Up
-- +goose StatementBegin

-- Table "deploy_log" is used to log deployment sequences
CREATE TABLE IF NOT EXISTS deploy_log
(
    id     INTEGER PRIMARY KEY AUTOINCREMENT,
    name   TEXT,
    state  INT,
    reason TEXT,
    created_at TIMESTAMP DEFAULT current_timestamp
);

CREATE INDEX deploy_log_name_idx ON deploy_log(name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS deploy_log;
-- +goose StatementEnd

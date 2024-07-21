-- +goose Up
-- +goose StatementBegin

-- Table "resources" - contains information about virtual resources
CREATE TABLE IF NOT EXISTS resources
(
    name          TEXT PRIMARY KEY,
    service_name  TEXT REFERENCES services (name),
    image         TEXT REFERENCES deploy_templates(name),
    state         INT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE resources;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin

-- Table "resource_constructors" - contains presets for different images.
-- When resource is deployed - this table is used to define settings on working node
CREATE TABLE IF NOT EXISTS deploy_templates
(
    name              TEXT PRIMARY KEY,
    base_image        TEXT,

    deploy_settings   json -- velez_api.CreateSmerd_Request
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE deploy_templates;
-- +goose StatementEnd

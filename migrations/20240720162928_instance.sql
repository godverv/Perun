-- +goose Up
-- +goose StatementBegin

-- Table "instance" - is used to map virtual resources|services
-- to actually running applications on working nodes
CREATE TABLE IF NOT EXISTS instances
(
    name       TEXT,
    node_name  TEXT REFERENCES nodes (name),
    port       INT,
    state      INT,
    image_name TEXT REFERENCES deploy_templates(name)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS instances;
-- +goose StatementEnd

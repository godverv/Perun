-- +goose Up
-- +goose StatementBegin
INSERT INTO deploy_templates
        (         name,    base_image)
VALUES  ('postgres:16', 'postgres:16');

UPDATE deploy_templates
SET deploy_settings ='{
  "settings": {
    "ports": [
      {
        "container": 5432,
        "protoc": 1
      }
    ],
    "volumes": [
      {
        "container_path": "/var/lib/postgresql/data"
      }
    ]
  },
  "healthcheck": {
    "command": "pg_isready -U postgres",
    "interval_second": 5,
    "timeout_second": 5,
    "retries": 3
  },
  "env": {
    "POSTGRES_HOST_AUTH_METHOD": "trust"
  }
}'
WHERE name = 'postgres:16';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE
FROM constructors
WHERE name = '';

DROP TABLE constructors;
-- +goose StatementEnd

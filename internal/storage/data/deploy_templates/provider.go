package deploy_templates

import (
	"database/sql"
)

type Provider struct {
	conn *sql.DB
}

func New(conn *sql.DB) *Provider {
	return &Provider{
		conn: conn,
	}
}

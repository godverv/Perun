package resource_patterns

import (
	"github.com/godverv/matreshka"
	"github.com/godverv/matreshka/resources"

	"github.com/Red-Sock/Perun/internal/domain"
)

type ResourceConstructor func(resources matreshka.DataSources, resourceName string, serviceName string) (domain.Dependencies, error)

type storage map[string]ResourceConstructor

var s = storage{
	resources.PostgresResourceName: Postgres,
	resources.SqliteResourceName:   Sqlite,
}

func GetConstructor(name string) ResourceConstructor {
	return s[name]
}

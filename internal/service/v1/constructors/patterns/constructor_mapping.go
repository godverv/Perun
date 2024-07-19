package patterns

import (
	"github.com/godverv/matreshka"
	"github.com/godverv/matreshka/resources"

	"github.com/Red-Sock/Perun/internal/domain"
)

type InstanceConstructor func(resources matreshka.DataSources, resourceName string, serviceName string) (domain.Dependencies, error)

type instanceStorage map[string]InstanceConstructor

var is = instanceStorage{
	resources.PostgresResourceName: Postgres,
	resources.SqliteResourceName:   Sqlite,
}

func GetConstructor(name string) InstanceConstructor {
	return is[name]
}

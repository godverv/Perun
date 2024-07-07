package resource_patterns

import (
	"github.com/godverv/Velez/pkg/velez_api"
	"github.com/godverv/matreshka"

	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/internal/utils/pass_gen"
)

func Postgres(resources matreshka.DataSources, resourceName, serviceName string) (deps domain.Dependencies, err error) {
	pg, err := resources.Postgres(resourceName)
	if err != nil {
		return domain.Dependencies{}, err
	}

	if pg.Pwd == "" {
		pg.Pwd = pass_gen.GenPass()
	}

	if pg.Host == "" {
		pg.Host = pg.GetName()
	}
	var dependency domain.DataSource
	dependency.Name = serviceName + "_pg"

	dependency.SmerdReq = &velez_api.CreateSmerd_Request{
		Name:      dependency.Name,
		ImageName: "postgres:13.6",
		Hardware:  &velez_api.Container_Hardware{},
		Env: map[string]string{
			"POSTGRES_USER":     pg.User,
			"POSTGRES_PASSWORD": pg.Pwd,
			"POSTGRES_DB":       pg.DbName,
		},
		Healthcheck: &velez_api.Container_Healthcheck{
			Command:        "pg_isready -U postgres",
			IntervalSecond: 5,
			TimeoutSecond:  5,
			Retries:        3,
		},
	}

	deps.Resources = append(deps.Resources, dependency)
	return deps, nil
}

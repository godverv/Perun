package resource_patterns

import (
	"github.com/godverv/matreshka"

	"github.com/Red-Sock/Perun/internal/domain"
)

func Sqlite(resources matreshka.DataSources, resourceName, serviceName string) (deps domain.Dependencies, err error) {
	//res, err := resources.Sqlite(resourceName)
	//if err != nil {
	//	return deps, errors.Wrap(err, "error extracting resource cfg")
	//}

	//volDep := &velez_api.VolumeBindings{
	//	Volume:        serviceName + "_" + resourceName,
	//	ContainerPath: path.Dir(res.Path),
	//}

	// TODO
	//deps.Volumes = append(deps.Volumes, volDep)

	return deps, nil
}

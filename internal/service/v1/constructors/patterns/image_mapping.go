package patterns

import (
	"github.com/godverv/matreshka/resources"
)

const (
	Postgres16Image = "postgres:16"
)

var typeToImageMapping = map[string]string{
	resources.PostgresResourceName: Postgres16Image,
}

func GetImageNameByType(resName string) string {
	return typeToImageMapping[resName]
}

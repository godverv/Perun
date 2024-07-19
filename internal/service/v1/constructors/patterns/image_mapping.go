package patterns

import (
	"github.com/godverv/matreshka/resources"
)

const (
	PostgresImage = "postgres:16"
)

var typeToImageMapping = map[string]string{
	resources.PostgresResourceName: PostgresImage,
}

func GetImageNameByType(resName string) string {
	return typeToImageMapping[resName]
}

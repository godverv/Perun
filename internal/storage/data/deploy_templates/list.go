package deploy_templates

import (
	"context"
	"encoding/json"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/Velez/pkg/velez_api"

	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/internal/storage"
)

func (p *Provider) GetByResourceName(ctx context.Context, resourceName string) (domain.ResourceConstructor, error) {
	row := p.conn.QueryRowContext(ctx, `
SELECT 
		dt.name,
		dt.base_image,
		dt.deploy_settings

FROM 		deploy_templates  AS dt
INNER JOIN  resources		  AS res 
ON  		res.image 		   = dt.name
AND 		res.name 		   = $1
`, resourceName)

	deployTemplate, err := toDomain(row)
	if err != nil {
		return domain.ResourceConstructor{}, errors.Wrap(err, "error reading from db")
	}

	deployTemplate.DeployRequest.Name = resourceName

	return deployTemplate, nil
}

func toDomain(row storage.Row) (domain.ResourceConstructor, error) {
	var out domain.ResourceConstructor

	var containerSettings []byte

	out.DeployRequest = &velez_api.CreateSmerd_Request{
		Healthcheck: &velez_api.Container_Healthcheck{},
		Hardware:    &velez_api.Container_Hardware{},
	}
	err := row.Scan(
		&out.Name,
		&out.BaseImage,

		&containerSettings,
	)
	if err != nil {
		return out, errors.Wrap(err, "error scanning")
	}

	err = json.Unmarshal(containerSettings, &out.DeployRequest)
	if err != nil {
		return out, errors.Wrap(err, "error unmarshalling settings")
	}

	out.DeployRequest.ImageName = out.BaseImage

	return out, nil
}

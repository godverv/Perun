package deploy_log

import (
	"github.com/Red-Sock/Perun/internal/storage"
)

type DeployLogService struct {
	deployLogData storage.DeployLogs
}

func New(data storage.Data) *DeployLogService {
	return &DeployLogService{
		deployLogData: data.DeployLogs(),
	}
}

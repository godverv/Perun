package domain

type InitServiceReq struct {
	ServiceName       string
	ImageName         string
	ReplicationFactor int
}

type RefreshService struct {
	ServiceName string
}

package domain

type InitServiceReq struct {
	ServiceName       string
	ImageName         string
	ReplicationFactor int
}

type SyncServiceInfo struct {
	Service Service
}

package domain

type Velez struct {
	Name               string
	Addr               string
	Port               int
	CustomVelezKeyPath string
	IsInsecure         bool
}

type Ssh struct {
	Key      []byte
	Port     uint64
	Username string
}

type VelezConnection struct {
	Node Velez
	Ssh  Ssh
}

type ListVelezNodes struct {
	SearchPattern string
	Paging
}

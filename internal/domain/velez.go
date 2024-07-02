package domain

type Velez struct {
	Name               string
	Addr               string
	CustomVelezKeyPath string
	IsInsecure         bool
}

type Ssh struct {
	Key      []byte
	Addr     string
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

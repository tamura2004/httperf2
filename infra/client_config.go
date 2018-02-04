package infra

type clientConfig struct {
	InsecureSkipVerify  bool
	MaxIdleConnsPerHost int
	Proxy               string
	Bps                 int
}

var ClientConfig clientConfig

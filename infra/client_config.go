package infra

type clientConfig struct {
	InsecureSkipVerify  bool
	MaxIdleConnsPerHost int
	Proxy               string
	Bps                 int
	UserAgent           string
}

var ClientConfig clientConfig

package infra

type clientConfig struct {
	InsecureSkipVerify  bool
	MaxIdleConnsPerHost int
	Proxy               string
	Bps                 int
	UserAgent           string
	UserName            string
	Password            string
}

var ClientConfig clientConfig

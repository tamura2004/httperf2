package domain

type Client struct {
	InsecureSkipVerify  bool
	MaxIdleConnsPerHost int
	Proxy               string
	Bps                 int
}

package infra

import (
	"crypto/tls"
	"github.com/tamura2004/httperf2/usecase"
	"io"
	"log"
	"net/http"
	"net/url"
)

type client struct {
	http.Client
}

func (c *client) Get(url string) io.ReadCloser {
	res, err := c.Client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if ClientConfig.Bps != 0 {
		return NewSlowReader(res.Body, ClientConfig.Bps)
	}
	return res.Body
}

func InitClient() {
	tr := newTransport()
	cl := http.Client{Transport: tr}
	usecase.InitClient(
		&client{Client: cl},
	)
	log.Println("init client")

}

func newTransport() *http.Transport {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: ClientConfig.InsecureSkipVerify,
		},
		MaxIdleConnsPerHost: ClientConfig.MaxIdleConnsPerHost,
	}
	tr = addProxy(tr)
	return tr
}

func addProxy(tr *http.Transport) *http.Transport {
	if ClientConfig.Proxy == "" {
		return tr
	}

	proxyURL, err := url.Parse(ClientConfig.Proxy)
	if err != nil {
		log.Fatal(err)
	}

	tr.Proxy = http.ProxyURL(proxyURL)
	return tr
}

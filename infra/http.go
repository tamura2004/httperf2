package infra

import (
	"crypto/tls"
	"github.com/tamura2004/httperf2/domain"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Client struct {
	domain domain.Client
	http   http.Client
}

func (c *Client) Get(url string) io.ReadCloser {
	res, err := c.http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if c.domain.Bps != 0 {
		return NewSlowReader(res.Body, c.domain.Bps)
	}
	return res.Body
}

func NewClient(c domain.Client) *Client {
	tr := newTransport(c)
	cl := http.Client{Transport: tr}
	return &Client{
		http:   cl,
		domain: c,
	}
}

func newTransport(c domain.Client) *http.Transport {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: c.InsecureSkipVerify,
		},
		MaxIdleConnsPerHost: c.MaxIdleConnsPerHost,
	}
	tr = addProxy(tr, c)
	return tr
}

func addProxy(tr *http.Transport, c domain.Client) *http.Transport {
	if c.Proxy == "" {
		return tr
	}

	proxyURL, err := url.Parse(c.Proxy)
	if err != nil {
		log.Fatal(err)
	}

	tr.Proxy = http.ProxyURL(proxyURL)
	return tr
}

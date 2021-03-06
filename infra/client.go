package infra

import (
	"crypto/tls"
	"encoding/base64"
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
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	if ClientConfig.UserAgent != "" {
		req.Header.Set("User-Agent", ClientConfig.UserAgent)
	}

	if ClientConfig.UserName != "" {
		req.Header.Set("Proxy-Authorization", "Basic "+basicAuth(
			ClientConfig.UserName,
			ClientConfig.Password,
		))
	}

	res, err := c.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if ClientConfig.Bps != 0 {
		return NewSlowReader(res.Body, ClientConfig.Bps)
	}
	return res.Body
}

type ClientFactory func() usecase.Client

func (f ClientFactory) NewClient() usecase.Client {
	return f()
}

func InitClient() {
	usecase.ClientFactory = ClientFactory(NewClient)
	log.Println("init client")
}

func NewClient() usecase.Client {
	return &client{
		Client: http.Client{
			Transport: newTransport(),
		},
	}
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

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

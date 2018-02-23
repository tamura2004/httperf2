package usecase

import (
	"io"
)

type Client interface {
	Get(string) io.ReadCloser
}

type clientFactory interface {
	NewClient() Client
}

var ClientFactory clientFactory

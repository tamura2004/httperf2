package usecase

import (
	"io"
)

type Client interface {
	Get(string) io.ReadCloser
}

var client Client

func InitClient(c Client) {
	client = c
}

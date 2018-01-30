package domain

import (
	"io"
	"sync"
)

type Manager struct {
	Config
	Result chan *Result
	WaitGroup
	Client client
}

type WaitGroup interface {
	Add(int)
	Done()
	Wait()
}

type client interface {
	Get(url string) io.ReadCloser
}

func NewManager(cf Config, cl client) *Manager {
	return &Manager{
		Config:    cf,
		Client:    cl,
		Result:    make(chan *Result),
		WaitGroup: &sync.WaitGroup{},
	}
}

package domain

import (
	"sync"
)

type Worker struct {
	Id int
	Config
	Result chan *Result
	WaitGroup
	Client client
}

var workerMutex sync.Mutex
var workerId int

func getWorkerId() int {
	workerMutex.Lock()
	workerId++
	workerMutex.Unlock()
	return workerId
}

func NewWorker(m *Manager) *Worker {
	return &Worker{
		Id:        getWorkerId(),
		Config:    m.Config,
		Result:    m.Result,
		WaitGroup: m.WaitGroup,
		Client:    m.Client,
	}
}

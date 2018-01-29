package usecase

import (
	"github.com/tamura2004/httperf2/domain"
)

type JobRuner interface {
	Run()
}

type Manager struct {
	Ch chan domain.Result
}

func NewManager() *Manager {
	return &Manager{
		Ch: make(chan domain.Result),
	}
}

func (m *Manager) Run() {

}

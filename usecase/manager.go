package usecase

import (
	"github.com/tamura2004/httperf2/domain"
)

type Manager struct {
	*domain.Manager
	*Collector
}

func NewManager(m *domain.Manager, c *Collector) *Manager {
	return &Manager{
		Manager:   m,
		Collector: c,
	}
}

func (m *Manager) Run() {
	for i := 0; i < m.Scinario.Worker; i++ {
		m.Scinario.RampUp.Sleep()
		worker := NewWorker(m)
		m.WaitGroup.Add(1)
		go worker.Run(m.WaitGroup)
	}

	go func() {
		m.WaitGroup.Wait()
		close(m.Manager.Result)
	}()

	m.Collector.Run()
}

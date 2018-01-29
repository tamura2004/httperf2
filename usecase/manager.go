package usecase

import (
	"github.com/tamura2004/httperf2/domain"
	"sync"
	"time"
)

type Client interface {
	Get(url string) io.ReadCloser
}

type JobRuner interface {
	Run(chan *domain.Result)
}

type JobRunerFactory interface {
	New(Client, domain.Target) JobRuner
}

type Manager struct {
	Result   chan domain.Result
	Scinario *domain.Scinario
	Target   *domain.Target
	Client   Client
	WG       *sync.WaitGroup
	Factory  JobRunerFactory
}

func NewManager(s *Scinario, t *Target, cl Client, jf JobRunerFactory) *Manager {
	return &Manager{
		Result:   make(chan domain.Result, 2048),
		Scinario: s,
		Target:   t,
		WG:       &sync.WaitGroup{},
		Client:   cl,
		Factory:  jf,
	}
}

func (m *Manager) Run() {
	for i := 0; i < m.Scinario.Worker; i++ {
		time.Sleep(100 * time.Millisecond)
		m.WG.Add(1)

		runner := m.Factory.New(m.Client, m.Target)
		go runner.Run(m.Result)
	}
}

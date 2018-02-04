package usecase

type Manager struct {
	*Collector
}

func NewManager(c *Collector) *Manager {
	return &Manager{
		Collector: c,
	}
}

func (m *Manager) Run() {
	for i := 0; i < Scinario.WorkerNum; i++ {
		worker := NewWorker(i)
		waitGroup.Add(1)
		go worker.Run()
	}

	go func() {
		waitGroup.Wait()
		close(resultChan)
	}()

	m.Collector.Run()
}

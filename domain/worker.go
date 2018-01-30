package domain

type Worker struct {
	Id int
	Config
	Result chan *Result
	WaitGroup
	Client client
}

func NewWorker(m *Manager, id int) *Worker {
	return &Worker{
		Id:        id,
		Config:    m.Config,
		Result:    m.Result,
		WaitGroup: m.WaitGroup,
		Client:    m.Client,
	}
}

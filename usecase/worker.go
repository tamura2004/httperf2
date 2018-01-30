package usecase

import (
	"github.com/tamura2004/httperf2/domain"
	"io"
	"io/ioutil"
)

type Worker struct {
	*domain.Worker
}

type Client interface {
	Get(string) io.ReadCloser
}

func NewWorker(m *Manager) *Worker {
	return &Worker{
		Worker: domain.NewWorker(m.Manager),
	}
}

func (w *Worker) Run(wg domain.WaitGroup) {
	defer wg.Done()

	for i := 0; i < w.Scinario.Count; i++ {
		job := domain.NewJob()

		w.Result <- job.TimeStart()

		r := w.Connect()
		w.Result <- job.TimeConnect()

		w.Write(r)
		w.Result <- job.TimeFinish()

		w.Scinario.Interval.Sleep()
	}
}

func (w *Worker) Connect() io.ReadCloser {
	return w.Client.Get(w.Target.Url)
}

func (w *Worker) Write(r io.ReadCloser) {
	defer r.Close()
	io.Copy(ioutil.Discard, r)
}

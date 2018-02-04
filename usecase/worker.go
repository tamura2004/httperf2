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

func NewWorker(m *Manager, id int) *Worker {
	return &Worker{
		Worker: domain.NewWorker(m.Manager, id),
	}
}

func (w *Worker) Run(wg domain.WaitGroup) {
	defer wg.Done()

	w.Scinario.RampUp.Times(w.Id).Sleep()

	for i := 0; i < w.Scinario.Count; i++ {

		job := domain.NewJob(i)
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

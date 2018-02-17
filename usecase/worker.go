package usecase

import (
	"github.com/tamura2004/httperf2/domain"
)

type Worker struct {
	Id     int
	client Client
}

func NewWorker(id int) *Worker {
	return &Worker{
		Id: id,
	}
}

func (w *Worker) Run() {
	defer waitGroup.Done()

	Scinario.RampUp.Times(w.Id).Sleep()

	for i := 0; i < Scinario.Count; i++ {

		job := domain.NewJob(w.Id, i)
		resultChan <- job.TimeStart()

		r := w.client.Get(Target.Url)
		resultChan <- job.TimeConnect()

		HttpBodyPrinter.Print(r)
		resultChan <- job.TimeFinish()

		Scinario.Interval.Sleep()
	}
}

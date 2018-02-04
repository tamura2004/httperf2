package usecase

import (
	"github.com/tamura2004/httperf2/domain"
	"io"
	"io/ioutil"
)

type Worker struct {
	Id int
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

		r := client.Get(Target.Url)
		resultChan <- job.TimeConnect()

		w.Write(r)
		resultChan <- job.TimeFinish()

		Scinario.Interval.Sleep()
	}
}

func (w *Worker) Write(r io.ReadCloser) {
	defer r.Close()
	io.Copy(ioutil.Discard, r)
}

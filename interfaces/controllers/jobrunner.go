package controllers

import (
	"github.com/tamura2004/httperf2/domain"
	"io"
	"os"
)

type Client interface {
	Get(url string) io.ReadCloser
}

type JobRunner struct {
	Client
	*domain.Job
}

type JobRunnerFactory struct{}

func (jf *JobRunnerFactory) New(cl Client, t *domain.Target) *JobRunner {
	return NewJobRunner(cl, domain.NewJob(t))
}

func NewJobRunner(c Client, j *domain.Job) *JobRunner {
	return &JobRunner{
		Client: c,
		Job:    j,
	}
}

func (j *JobRunner) Run(result chan *domain.Result) {
	result <- j.Job.TimeStart()
	r := j.Connect()
	defer r.Close()
	result <- j.Job.TimeConnect()
	j.Write(r)
	result <- j.Job.TimeFinish()
}

func (j *JobRunner) Connect() io.ReadCloser {
	return j.Client.Get(j.Target.Url)
}

func (j *JobRunner) Write(r io.ReadCloser) {
	io.Copy(os.Stdout, r)
}

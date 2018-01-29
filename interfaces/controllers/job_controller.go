package controllers

import (
	"fmt"
	"github.com/tamura2004/httperf2/domain"
	"io"
	"os"
	"time"
)

type Client interface {
	Get(url string) io.ReadCloser
}

type JobRunner struct {
	Client
	*domain.Job
}

func NewJobRunner(c Client, j *domain.Job) *JobRunner {
	return &JobRunner{
		Client: c,
		Job:    j,
	}
}

func (j *JobRunner) Run(ch chan *domain.Result) {
	ch <- domain.StartResult(j)

	r := j.Client.Get(j.Target.Url)
	defer r.Close()

	ch <- domain.ConnectResult(j)

	io.Copy(os.Stdout, r)

	ch <- domain.FinishResult(j)
}

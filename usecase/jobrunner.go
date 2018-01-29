package usecase

import (
	"github.com/tamura2004/httperf2/domain"
)

type JobRunner struct {
	domain.Job
	domain.Client
}

func (j *JobRunner) Run() domain.Result {

}

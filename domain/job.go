package domain

import (
	"time"
)

type Job struct {
	WorkerId int
	JobId    int
	Stime    time.Time
}

func NewJob(w, j int) Job {
	return Job{
		WorkerId: w,
		JobId:    j,
		Stime:    _time.Now(),
	}
}

func (j Job) TimeStart() *Result {
	return NewResult(j, START)
}

func (j Job) TimeConnect() *Result {
	return NewResult(j, CONNECT)
}

func (j Job) TimeFinish() *Result {
	return NewResult(j, FINISH)
}

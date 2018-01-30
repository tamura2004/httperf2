package domain

import (
	"time"
)

type Job struct {
	Id    int
	Stime time.Time
	Worker
}

func NewJob(id int) *Job {
	return &Job{
		Id:    id,
		Stime: time.Now(),
	}
}

func (j *Job) TimeStart() *Result {
	return NewResult(j, START)
}

func (j *Job) TimeConnect() *Result {
	return NewResult(j, CONNECT)
}

func (j *Job) TimeFinish() *Result {
	return NewResult(j, FINISH)
}

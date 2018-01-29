package domain

import (
	"sync"
	"time"
)

type Job struct {
	Id    int
	Stime time.Time
	Worker
	*Target
}

var mu sync.Mutex
var id int

func NewJob(t *Target) *Job {
	return &Job{
		Id:     getId(),
		Stime:  time.Now(),
		Target: t,
	}
}

func getId() int {
	mu.Lock()
	id++
	mu.Unlock()
	return id
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

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

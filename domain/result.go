package domain

import (
	"time"
)

type Result struct {
	*Job
	Check    CheckPoint    // 記録が取られたタイミング
	Time     time.Time     // 現在時刻
	Duration time.Duration // 経過時間
}

func NewResult(j *Job, cp CheckPoint) *Result {
	return &Result{
		Job:      j,
		Check:    cp,
		Time:     time.Now(),
		Duration: time.Since(j.Stime),
	}
}

func StartResult(j *Job) *Result {
	return NewResult(j, START)
}

func ConnectResult(j *Job) *Result {
	return NewResult(j, CONNECT)
}

func FinishResult(j *Job) *Result {
	return NewResult(j, Finish)
}

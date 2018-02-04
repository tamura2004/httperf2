package domain

import (
	"time"
)

type Result struct {
	Job
	Check    CheckPoint    // 記録が取られたタイミング
	Time     time.Time     // 現在時刻
	Duration time.Duration // 経過時間
}

func NewResult(j Job, cp CheckPoint) *Result {
	return &Result{
		Job:      j,
		Check:    cp,
		Time:     _time.Now(),
		Duration: _time.Since(j.Stime),
	}
}

package infra

import (
	"github.com/tamura2004/httperf2/domain"
	"time"
)

type _time struct{}

func InitTime() {
	domain.InitTime(&_time{})
}

func (*_time) Now() time.Time {
	return time.Now()
}

func (*_time) Since(t time.Time) time.Duration {
	return time.Since(t)
}

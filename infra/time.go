package infra

import (
	"github.com/tamura2004/httperf2/domain"
	// "log"
	// "math/rand"
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

type mock struct{}

func InitMockTime() {
	domain.InitTime(&mock{})
}

func (c *mock) Now() time.Time {
	return time.Date(2018, time.February, 1, 8, 0, 0, 0, time.UTC)
}

func (c *mock) Since(t time.Time) time.Duration {
	return time.Duration(10 * time.Second)
}

package mock

import (
	"github.com/tamura2004/httperf2/domain"
	"time"
)

type _time struct{}

func InitTime() {
	domain.InitTime(&_time{})
}

func (c *_time) Now() time.Time {
	return time.Date(2018, time.February, 1, 8, 0, 0, 0, time.UTC)
}

func (c *_time) Since(t time.Time) time.Duration {
	return time.Duration(10 * time.Second)
}

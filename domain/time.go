package domain

import (
	"time"
)

type clock interface {
	Now() time.Time
	Since(time.Time) time.Duration
}

var _time clock

func InitTime(c clock) {
	_time = c
}

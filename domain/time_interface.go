package domain

import (
	"time"
)

type clock interface {
	Now() time.Time
	Since(time.Time) time.Duration
}

var Time clock

func InitTime(c clock) {
	Time = c
}
